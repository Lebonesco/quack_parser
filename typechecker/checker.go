package typechecker

import (
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"errors"
	"reflect"
	"strings"
)

type ErrorType string

type CheckError struct {
	Type ErrorType
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
)

func createError(errorType, message string) *CheckError {
	return &CheckError{Type: ErrorType(errorType), Message: errors.New(message)}
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
	case *ast.MethodCall:
		return evalMethodCall(node, env)
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

	for _, statement := range p.Statements {
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
	// extract signatures
	setSignatures(classes, env)
	// check for cycles 
	if env.CycleExist() {
		createError(CLASS_CYCLE, "class cycle")
	}
	// check for nonexistent types
	if !env.TypesExist() {
		createError(CLASS_NOT_EXIST, "class not exist")
	}
	// extract methods
	err := setMethods(classes, env)
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
		obj := (*env.TypeTable)[ObjectType(class.Signature.Name)] // get type object
		if obj.Parent != "" {
			parent, ok := (*env.TypeTable)[obj.Parent]
			if !ok {
				createError(CLASS_NOT_EXIST, "parent class not exist")
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
	ok := compareBlockVars(child.Variables, parent.Variables)
	if !ok {
		createError(CREATE_CLASS_FAIL, "variables in child incompatible with parent")
	}
	// compare methods
	ok = compareMethods(child.MethodTable, parent.MethodTable)
	if !ok {
		createError(CREATE_CLASS_FAIL, "child is missing methods found in parent")
	}
	return nil
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
	obj := (*env.TypeTable)[ObjectType(class.Signature.Name)] // get type object // is this a copy or a pointer?
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
		if strings.HasPrefix(k, "this.") { // if class variable
			obj.Variables[k] = newEnv.Vals[k]
		}
	}
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
		result, err := TypeCheck(method.StmtBlock, newEnv)
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
		sig := MethodSignature{Params: []Variable{}}
		if _, ok := obj.MethodTable[method.Name]; ok {
			return createError(ALREADY_INITIALIZED, "method name already exists in class")
		}

		// add arguments
		for i, arg := range method.Args {
			if !env.TypeExist(ObjectType(arg.Type)) {
				return createError(CLASS_NOT_EXIST, "type in method params not exist")
			}

			sig.Params = append(sig.Params, Variable{arg.Arg, ObjectType(arg.Type)})
		}

		if method.Typ != "" {
			if !env.TypeExist(ObjectType(method.Typ)) {
				return createError(CLASS_NOT_EXIST, "type in method return signature not exist")
			}
			sig.Return = ObjectType(method.Typ)
		} // maybe put nothing type if nothing returned

		obj.AddMethod(method.Name, sig)
	}
	return nil
}

func setSignatures(classes []ast.Class, env *Environment) {
	for _, class := range classes {
		setSignature(class, env)
	}
}

func setSignature(class ast.Class, env *Environment) {
	sig := class.Signature
	obj := NewObject()

	obj.Type = ObjectType(class.Signature.Name)
	// add arguments to constructor
	for _, arg := range sig.Args {
		obj.Constructor = append(obj.Constructor, Variable{arg.Arg, ObjectType(arg.Type)})
	}

	// if subtype, set super type
	if sig.Extend != nil {
		obj.Parent = ObjectType(sig.Extend.Parent)
	}

	(*env.TypeTable)[obj.Type] = obj // store object type
}

// compares 2 environments
func compareStmtBlockVars(env1, env2 *Environment) bool {
	vars1 := env1.Vals 
	vars2 := env2.Vals
	return compareBlockVars(vars1, vars2)
}

// makes sure that two Statement Blocks have same variables/types at end
func compareBlockVars(vars1, vars2 map[string]ObjectType) bool {
	for k := range vars1 {
		if val2, ok := vars2[k]; ok && vars1[k] == val2 {
			continue
		} else {
			return false
		}
	}
	return true
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
		if !env.ValidSubType(result.Type, rightType.Type) {
			return result, createError(INCOMPATABLE_TYPES, "not subtype of the other")
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
	// create environment for each scope
	newEnv1 := env.NewScope()
	//newEnv2 := env.NewScope()

	result1, err := TypeCheck(node.Consequence, newEnv1)
	if err != nil {
		return result1, err
	}

	// if node.Alternative != nil {
	// 	result2, err := TypeCheck(node.Alternative, newEnv2)
	// 	if err != nil {
	// 		return result2, err
	// 	}
	// } 

	// compare environments or to current environment

	return result, nil
}

func evalWhileStatement(node *ast.WhileStatement, env *Environment) (Variable, *CheckError) {
	var result Variable
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

func evalFunctionCall(node *ast.FunctionCall, env *Environment) (Variable, *CheckError) {
	class := (*env.TypeTable)[ObjectType(node.Name)]
	args := class.Constructor

	if len(args) != len(node.Args) {
		return Variable{}, createError(BAD_FUNCTION_CALL, "incorrect amount of arguments")
	}

	for i, arg := range args {
		result, err := TypeCheck(node.Args[i], env)
		if err != nil {
			return result, err
		}

		if arg.Type != result.Type {
			return Variable{}, createError(INCOMPATABLE_TYPES, "incorrect argument type")
		}
	}

	return Variable{Type: class.Type}, nil // return the type of the class 
}

// eval something like this class.method()
func evalMethodCall(node *ast.MethodCall, env *Environment) (Variable, *CheckError) {
	class, ok := (*env.TypeTable)[ObjectType(node.Variable)]
	if !ok {
		return Variable{}, createError(CLASS_NOT_EXIST, "class not exist on method call")
	}

	signature, ok := class.MethodTable[node.Method]
	if !ok {
		return Variable{}, createError(METHOD_NOT_EXIST, "method not exist in class")
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

	return Variable{Type: signature.Return}, nil
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
		return right, createError(INCOMPATABLE_TYPES, "types not work for infix expression")
	}
	return right, nil
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
		return Variable{}, createError(VARIABLE_NOT_INITIALIZED, "ident is not defined")
	}
	return Variable{Name: node.Value, Type: IdentType}, nil
}