package typechecker

import (
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"errors"
	"reflect"
	"fmt"
	"strings"
)

type CheckError struct {
	Type string
	Message error
}

// error types
const (
	INVALID_SUBCLASS = "INVALID_SUBCLASS"
	CREATE_CLASS_FAIL = "CREATE_CLASS_FAIL"
	CLASS_NOT_EXIST = "CLASS_NOT_EXIST"
	METHOD_NOT_EXIST = "METHOD_NOT_EXIST"
	DUPLICATE_CLASS = "DUPLICATE_CLASS"
	CLASS_CYCLE = "CLASS_CYCLE"
	INVALID_OPERATION_TYPE = "INVALID_OPERATION_TYPE"
	VARIABLE_NOT_INITIALIZED = "VARIABLE_NOT_INITIALIZED"
	ALREADY_INITIALIZED = "ALREADY_INITIALIZED"
	INVALID_METHOD_TYPE = "INVALID_METHOD_TYPE"
	INVALID_CONSTRUCTOR_TYPE = "INVALID_CONSTRUCTOR_TYPE"
	INCOMPATABLE_TYPES = "INCOMPATABLE_TYPES"
	BAD_FUNCTION_CALL = "BAD_FUNCTION_CALL"
	CONDITION_NOT_BOOL = "CONDITION_NOT_BOOL"
)

func createError(errorType, message string, args ...interface{}) *CheckError {
	return &CheckError{Type: errorType, Message: errors.New(fmt.Sprintf(message, args...))}
}

// start of checker
func TypeCheck(node ast.Node, env *Environment) (Variable, *CheckError) {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.ReturnStatement:
		return evalReturnStatement(node, env)
	case *ast.IfStatement:
		return evalIfStatement(node, env)
	case *ast.WhileStatement:
		return evalWhileStatement(node, env)
	case *ast.ExpressionStatement:
		return evalExpressionStatement(node, env)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)
	case *ast.InfixExpression:
		return evalInfixExpression(node, env)
	case *ast.IntegerLiteral:
		return evalInteger(node, env)
	case *ast.StringLiteral:
		return evalString(node, env)
	case *ast.Boolean:
		return evalBoolean(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.LetStatement:
		return evalLetStatement(node, env)
	case *ast.FunctionCall: // actually a class call, ei PT(1, 2);
		return evalFunctionCall(node, env)
	case *ast.MethodCall: // handle class.method()
		return evalMethodCall(node, env)
	case *ast.ClassVariableCall:
		return evalClassVariableCall(node, env)
	}
	return Variable{}, nil
}

func evalProgram(p *ast.Program, env *Environment) (Variable, *CheckError) {
	var result Variable

	if err := evalBuiltIns(env); err != nil {
		return result, createError(INVALID_SUBCLASS, "builtin error")
	}

	if err := evalClasses(p.Classes, env); err != nil {
		return result, err
	}

	for _, statement := range p.Statements { // places this in cycle until env stops changing 
		result, err := TypeCheck(statement, env)
		if err != nil {
			return result, err
		}
	}

	return result, nil
} 

// populate built in classes
func evalBuiltIns(env *Environment) error {
	code := BUIILT_IN_CLASSES
	l := lexer.NewLexer([]byte(code))
	p := parser.NewParser()
	program, err := p.Parse(l)
	if err != nil {
		return err
	}

	classes := program.(*ast.Program).Classes
	errr := evalClasses(classes, env)
	if errr != nil {
		return errr.Message
	}
	return nil
}

func evalClasses(classes []ast.Class, env *Environment) (*CheckError) {
	if len(classes) == 0 {
		return nil
	}
	// extract signatures
	err := setSignatures(classes, env)
	if err != nil {
		return err
	}
	// check for cycles 
	if env.CycleExist() {
		return createError(CLASS_CYCLE, "class cycle")
	}
	// check for nonexistent types
	if !env.TypesExist() {
		return createError(CLASS_NOT_EXIST, "class not exist")
	}
	// extract methods
	err = setMethods(classes, env)
	if err != nil {
		return err
	}
	
	err = checkClasses(classes, env)
	if err != nil {
		return err
	}
	// compare to parent if exist
	err = compareParents(classes, env)
	if err != nil {
		return err
	}

	// check methods
	err = checkMethods(classes, env)
	if err != nil {
		return err
	}
	return nil
}

func compareParents(classes []ast.Class, env *Environment) (*CheckError) {
	for _, class := range classes {
		obj := env.GetClass(ObjectType(class.Signature.Name)) // get type object
		if obj.Parent != "" {
			parent := env.GetClass(obj.Parent)
			if parent == nil {
				return createError(CLASS_NOT_EXIST, "parent class not exist")
			}
			err := compareParent(obj, parent)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func compareParent(child, parent *Object) (*CheckError) {
	// how to handle parent with input?
	// compare variables
	fmt.Println(child.Variables, parent.Variables)
	ok := compareBlockVars(child.Variables, parent.Variables)
	if !ok {
		return createError(CREATE_CLASS_FAIL, "variables in %s incompatible with %s", child.Type, parent.Type)
	}
	// compare methods
	ok = compareMethods(child.MethodTable, parent.MethodTable)
	if !ok {
		return createError(CREATE_CLASS_FAIL, "child is missing methods found in parent")
	}
	return nil
}

// makes sure that two Statement Blocks have same variables/types at end
func compareBlockVars(child, parent map[string]ObjectType) bool {
	for k := range parent {
		val, ok := child[k]; 
		if !ok {
			return false
		} 

		if parent[k] != val {
			return false
		}
	}
	return true
}

func compareMethods(child, parent map[string]MethodSignature) (bool) {
	return true
}

func checkClasses(classes []ast.Class, env *Environment) (*CheckError) {
	for _, class := range classes {
		err := checkClass(class, env)
		if err != nil {
			return err
		}
	}
	return nil
}

// extract class vars + check constructor/types and methods
func checkClass(class ast.Class, env *Environment) (*CheckError) {
	// check constructor types
	// do it
	newEnv := env.NewScope()

	obj := (*env.TypeTable)[ObjectType(class.Signature.Name)] // get type object
	// populate with constructor variables
	for _, v := range obj.Constructor {
		newEnv.Set(v.Name, v.Type)
	}

	// check constructor
	for _, statement := range class.Body.Statements {
		// extract all identifiers
		_, err := TypeCheck(statement, newEnv)
		if err != nil {
			return err
		}
	}
	// extract class variables
	for k := range newEnv.Vals {
		if strings.HasPrefix(k, "this.") { // if class variable add to class
			obj.Variables[k] = newEnv.Vals[k]
		}
	}
	//(*env.TypeTable)[ObjectType(class.Signature.Name)].Variables = obj.Variables
	return nil
}

func checkMethods(classes []ast.Class, env *Environment) (*CheckError) {
	for _, class := range classes {
		err := checkMethod(class, env)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkMethod(class ast.Class, env *Environment) (*CheckError) {
	methods := class.Body.Methods 
	obj := (*env.TypeTable)[ObjectType(class.Signature.Name)] // get type object

	for _, method := range methods {
		newEnv := env.NewScope()
		newEnv.SetClass(obj.Type) // this will allow methods to access class instance of 'this' and methods
		// populate with params
		for _, arg := range obj.MethodTable[method.Name].Params {
			newEnv.Set(arg.Name, arg.Type)
		}
		// popular with class variables
		for k := range obj.Variables {
			newEnv.Set(string(k), obj.Variables[string(k)])
		}

		result, err := TypeCheck(method.StmtBlock, newEnv)
		if err != nil {
			return err
		}
		// check if received return statement to compare against method signature
		if reflect.TypeOf(result) == reflect.TypeOf(ast.ReturnStatement{}) {
			if result.Type != obj.MethodTable[method.Name].Return {
				return createError(INVALID_SUBCLASS, "incorrect return type in method")
			}
		}
	}
	return nil
}

func setMethods(classes []ast.Class, env *Environment) (*CheckError) {
	for _, class := range classes {
		err := setMethod(class, env)
		if err != nil {
			return err
		}
	}
	return nil
}

// error for: duplicates, nonexistent types
func setMethod(class ast.Class, env *Environment) (*CheckError) {
	methods := class.Body.Methods
	obj := (*env.TypeTable)[ObjectType(class.Signature.Name)] // get type object

	for _, method := range methods {
		sig := MethodSignature{Params: []Variable{}, Return: NOTHING_CLASS}
		if _, ok := obj.MethodTable[method.Name]; ok {
			return createError(ALREADY_INITIALIZED, "method name: %s already exists in class", method.Name)
		}

		if method.Name == string(obj.Type) {
			return createError(ALREADY_INITIALIZED, "method %s can't be the same as the class", method.Name)
		}

		// add arguments
		for _, arg := range method.Args {
			if !env.TypeExist(ObjectType(arg.Type)) {
				return createError(CLASS_NOT_EXIST, "type in method params not exist")
			}

			sig.Params = append(sig.Params, Variable{arg.Arg, ObjectType(arg.Type)})
		}

		if method.Typ != "" {
			if !env.TypeExist(ObjectType(method.Typ)) {
				return createError(CLASS_NOT_EXIST, "type '%s' in method return signature '%s' not exist", method.Typ, method.Name)
			}
			sig.Return = ObjectType(method.Typ)
		} // maybe put nothing type if nothing returned

		obj.AddMethod(method.Name, sig)
	}
	return nil
}

func setSignatures(classes []ast.Class, env *Environment) (*CheckError) {
	for _, class := range classes {
		err := setSignature(class, env)
		if err != nil {
			return err
		}
	}
	return nil
}

func setSignature(class ast.Class, env *Environment) (*CheckError) {
	sig := class.Signature
	obj := NewObject()

	obj.Type = ObjectType(class.Signature.Name)
	if env.TypeExist(obj.Type) {
		return createError(DUPLICATE_CLASS, "class %s already exists", obj.Type)
	}

	// add arguments to constructor
	for _, arg := range sig.Args {
		obj.Constructor = append(obj.Constructor, Variable{arg.Arg, ObjectType(arg.Type)})
	}

	// if subtype, set super type
	if sig.Extend != nil {
		obj.Parent = ObjectType(sig.Extend.Parent)
	}

	(*env.TypeTable)[obj.Type] = obj // store object type
	return nil
}

func evalLetStatement(node *ast.LetStatement, env *Environment) (Variable, *CheckError) {
	result := Variable{Name: node.Name.Value}
	right := node.Value

	rightType, err := TypeCheck(right, env)
	if err != nil {
		return result, err
	} 

	if node.Kind != "" { // if type explicitly set
		result.Type = ObjectType(node.Kind) 
		if !env.TypeExist(result.Type) {
			return Variable{}, createError(CLASS_NOT_EXIST, "class '%s' doesn't exist", result.Type)
		}

		if !env.ValidSubType(result.Type, rightType.Type) {
			return result, createError(INCOMPATABLE_TYPES, "%s not subtype of %s", result.Type, rightType.Type)
		}
	} else {
		result.Type = rightType.Type
	}

	env.Set(result.Name, result.Type) // set variable in environment
	return result, nil
}

func evalBlockStatement(block *ast.BlockStatement, env *Environment) (Variable, *CheckError) {
	var result Variable

	for _, statement := range block.Statements {
		result, err := TypeCheck(statement, env)
		if err != nil {
			return result, err
		}

		if reflect.TypeOf(statement) == reflect.TypeOf(ast.ReturnStatement{}) {
			return result, nil
		}
	}
	return result, nil
}

func evalIfStatement(node *ast.IfStatement, env *Environment) (Variable, *CheckError) {
	var result Variable

	// check that condition evals to bool type
	result, err := TypeCheck(node.Condition, env)
	if err != nil {
		return result, err
	}

	if result.Type != BOOL_CLASS {
		return result, createError(CONDITION_NOT_BOOL, "condition not evaluate to bool value")
	}

	// create environment for each scope
	newEnv1 := env.NewScope()
	newEnv2 := env.NewScope()

	_, err = TypeCheck(node.Consequence, newEnv1)
	if err != nil {
		return result, err
	}

	if node.Alternative == nil { // if no other statement don't bubble up variable
		return result, nil
	}


	_, err = TypeCheck(*node.Alternative, newEnv2)
	if err != nil {
		return result, err
	}

	// compare environments or to current environment
	union := GetUnion(newEnv1, newEnv2)
	for k := range union { 
		env.Set(k, union[k])
	}

	return result, nil
}

func evalWhileStatement(node *ast.WhileStatement, env *Environment) (Variable, *CheckError) {
	var result Variable

	// check that condition evals to bool type
	result, err := TypeCheck(node.Cond, env)
	if err != nil {
		return result, err
	}

	if result.Type != BOOL_CLASS {
		return result, createError(CONDITION_NOT_BOOL, "condition not evaluate to bool value")
	}

	newEnv := env.NewScope()
	for _, statement := range node.BlockStatement.Statements {
		result, err := TypeCheck(statement, newEnv)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func evalReturnStatement(node *ast.ReturnStatement, env *Environment) (Variable, *CheckError) {
	return TypeCheck(node.ReturnValue, env)
}

// init class Object or if in class method call, ei: PT(4, 5);
func evalFunctionCall(node *ast.FunctionCall, env *Environment) (Variable, *CheckError) {
	if env.TypeExist(ObjectType(node.Name)) { // if a class
		class := env.GetClass(ObjectType(node.Name))
		if class == nil {
			return Variable{}, createError(CLASS_NOT_EXIST, "class '%s' doesn't exist", node.Name)
		}
		args := class.Constructor

		if len(args) != len(node.Args) {
			return Variable{}, createError(BAD_FUNCTION_CALL, "incorrect amount of arguments for %s wants %d provided %d", node.Name, len(args), len(node.Args))
		}

		// check argument types
		for i, arg := range args {
			result, err := TypeCheck(node.Args[i], env)
			if err != nil {
				return result, err
			}

			if arg.Type != result.Type {
				return Variable{}, createError(INCOMPATABLE_TYPES, "incorrect argument type for Class %s", node.Name)
			}
		}

		return Variable{Type: class.Type}, nil // return the type of the class 
	}

	if signature, ok := env.GetClassObject().GetMethod(node.Name); ok { // if method
		for i, param := range signature.Params {
			result, err := TypeCheck(node.Args[i], env)
			if err != nil {
				return result, err
			}

			if param.Type != result.Type {
				return Variable{}, createError(INCOMPATABLE_TYPES, "incorrect argument type")
			}
		}
		return Variable{Type: signature.Return, Name: node.Name+"()"}, nil
	}

	return Variable{}, createError(CLASS_NOT_EXIST, "class '%s' doesn't exist", node.Name)
}

// eval something like this class.method()
func evalMethodCall(node *ast.MethodCall, env *Environment) (Variable, *CheckError) {
	class, err := TypeCheck(node.Variable, env) // left can be any expression that returns a type
	if err != nil {
		return Variable{}, err
	}

	signature, ok := env.GetClassMethod(class.Type, node.Method)
	if !ok {
		return Variable{}, createError(METHOD_NOT_EXIST, "method %s not exist in class %s", node.Method, class.Type)
	}

	// recursively check correct arguments provided
	for i, param := range signature.Params {
		result, err := TypeCheck(node.Args[i], env)
		if err != nil {
			return result, err
		}

		if param.Type != result.Type {
			return Variable{}, createError(INCOMPATABLE_TYPES, "incorrect argument type")
		}
	}

	return Variable{Type: signature.Return, Name: node.Method+"()"}, nil
}

func evalPrefixExpression(expr *ast.PrefixExpression, env *Environment) (Variable, *CheckError) {
	return Variable{}, nil
}

func evalInfixExpression(node *ast.InfixExpression, env *Environment) (Variable, *CheckError) {
	left, err := TypeCheck(node.Left, env)
	if err != nil {
		return left, err
	}

	right, err := TypeCheck(node.Right, env)
	if err != nil {
		return right, err
	}

	if right.Type != left.Type { // maybe should compare if subtypes
		return right, createError(INCOMPATABLE_TYPES, "types %s-%s and %s-%s not work for expression '%s' on line %d", left.Type, left.Name, right.Type, right.Name, node.Operator, node.Token.Pos.Line)
	}

	switch node.Operator {
	case "<", ">", "<=", ">=", "==", "!=", "and", "or":
		return Variable{Type: BOOL_CLASS}, nil
	}

	return left, nil
}

func evalOpeExpression(left, right Object) (Variable, *CheckError) {
	return Variable{}, nil
}

func evalInteger(node *ast.IntegerLiteral, env *Environment) (Variable, *CheckError) {
	return Variable{Name: string(node.Token.Lit), Type: INTEGER_CLASS}, nil
}

func evalString(node *ast.StringLiteral, env *Environment) (Variable, *CheckError) {
	return Variable{Name: node.Value, Type: STRING_CLASS}, nil
}

func evalBoolean(node *ast.Boolean, env *Environment) (Variable, *CheckError) {
	return Variable{Name: string(node.Token.Lit), Type: BOOL_CLASS}, nil
}

func evalIdentifier(node *ast.Identifier, env *Environment) (Variable, *CheckError) {
	IdentType, ok := env.Get(node.Value) // check if it has been defined
	if !ok {
		return Variable{}, createError(VARIABLE_NOT_INITIALIZED, "ident %s is not defined on line: %d", node.Value, node.Token.Pos.Line)
	}
	return Variable{Name: node.Value, Type: IdentType}, nil
}

func evalClassVariableCall(node *ast.ClassVariableCall, env *Environment) (Variable, *CheckError) {
	left, err := TypeCheck(node.Expression, env)
	if err != nil {
		return Variable{}, err
	}

	kind, ok :=  env.GetClassVariable(left.Type, node.Ident) // type, method
	if !ok {
		return Variable{}, createError(VARIABLE_NOT_INITIALIZED, "type %s doesn't have variable %s", left.Type, node.Ident)
	}
	return Variable{Type: kind, Name: string(node.Token.Lit)}, nil
}

// ei: "string";
func evalExpressionStatement(node *ast.ExpressionStatement, env *Environment) (Variable, *CheckError) {
	return TypeCheck(node.Expression, env)
}