package typechecker

import (
	"errors"
	"fmt"
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/lexer"
	"github.com/Lebonesco/quack_parser/parser"
	"github.com/Lebonesco/quack_parser/environment"
	"reflect"
	"strings"
)

type CheckError struct {
	Type    string
	Message error
}

const (
	PLUS    = "PLUS"
	EQUALS  = "EQUALS"
	ATMOST  = "ATMOST"
	ATLEAST = "ATLEAST"
	LESS    = "LESS"
	MORE    = "MORE"
	MINUS   = "MINUS"
	DIVIDE  = "DIVIDE"
	TIMES   = "TIMES"
	AND     = "AND"
	OR      = "OR"
)

// error types
const (
	INVALID_SUBCLASS         = "INVALID_SUBCLASS"
	CREATE_CLASS_FAIL        = "CREATE_CLASS_FAIL"
	CLASS_NOT_EXIST          = "CLASS_NOT_EXIST"
	METHOD_NOT_EXIST         = "METHOD_NOT_EXIST"
	DUPLICATE_CLASS          = "DUPLICATE_CLASS"
	CLASS_CYCLE              = "CLASS_CYCLE"
	INVALID_OPERATION_TYPE   = "INVALID_OPERATION_TYPE"
	VARIABLE_NOT_INITIALIZED = "VARIABLE_NOT_INITIALIZED"
	ALREADY_INITIALIZED      = "ALREADY_INITIALIZED"
	INVALID_METHOD_TYPE      = "INVALID_METHOD_TYPE"
	INVALID_CONSTRUCTOR_TYPE = "INVALID_CONSTRUCTOR_TYPE"
	INCOMPATABLE_TYPES       = "INCOMPATABLE_TYPES"
	BAD_FUNCTION_CALL        = "BAD_FUNCTION_CALL"
	CONDITION_NOT_BOOL       = "CONDITION_NOT_BOOL"
	INVALID_RETURN_TYPE      = "INVALID_RETURN_TYPE"
	INCORRECT_ARGUMENT_COUNT = "INCORRECT_ARGUMENT_COUNT"
)

func createError(errorType, message string, args ...interface{}) *CheckError {
	return &CheckError{Type: errorType, Message: errors.New(fmt.Sprintf(message, args...))}
}

// start of checker
func TypeCheck(node ast.Node, env *environment.Environment) (environment.Variable, *CheckError) {
	// inject current environment into each statement node for use in code gen
	switch node := node.(type) {
	case *ast.BlockStatement:
		node.Env = env
	case *ast.IfStatement:
		node.Env = env
	case *ast.WhileStatement:
		node.Env = env
	case *ast.ReturnStatement:
		node.Env = env
	case *ast.ExpressionStatement:
		node.Env = env
	case *ast.TypecaseStatement:
		node.Env = env
	case *ast.LetStatement:
		node.Env = env
	}

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
	case *ast.TypecaseStatement:
		return evalTypeCaseStatement(node, env)
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
	return environment.Variable{}, nil
}

func evalProgram(p *ast.Program, env *environment.Environment) (environment.Variable, *CheckError) {
	var result environment.Variable

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
func evalBuiltIns(env *environment.Environment) error {
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

// driver for class analysis
func evalClasses(classes []ast.Class, env *environment.Environment) *CheckError {
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

func compareParents(classes []ast.Class, env *environment.Environment) *CheckError {
	for _, class := range classes {
		obj := env.GetClass(environment.ObjectType(class.Signature.Name)) // get type environment.Object
		if obj.Parent != "" {
			parent := env.GetClass(obj.Parent)
			if parent == nil {
				return createError(CLASS_NOT_EXIST, "parent class not exist")
			}
			err := compareParent(obj, parent, env)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func compareParent(child, parent *environment.Object, env *environment.Environment) *CheckError {
	// how to handle parent with input?
	// compare variables
	ok := compareBlockVars(child.Variables, parent.Variables)
	if !ok {
		return createError(CREATE_CLASS_FAIL, "variables in %s incompatible with %s", child.Type, parent.Type)
	}
	// compare methods
	err := compareMethods(child, parent, env)
	if err != nil {
		return err
	}
	return nil
}

// makes sure that two Statement Blocks have same variables/types at end
func compareBlockVars(child, parent map[string]environment.ObjectType) bool {
	for k := range parent {
		val, ok := child[k]
		if !ok {
			return false
		}

		if parent[k] != val {
			return false
		}
	}
	return true
}

func compareMethods(child, parent *environment.Object, env *environment.Environment) *CheckError {
	for k, v := range child.MethodTable {
		if res, ok := parent.GetMethod(k); ok { // if method name in parent do check
			// check input types
			if len(v.Params) != len(res.Params) {
				return createError(INCORRECT_ARGUMENT_COUNT, "child overriding method have incorrect param length %d vs %d",
					len(v.Params), len(res.Params))
			}

			for i, param := range v.Params {
				if !env.ValidSubType(res.Params[i].Type, param.Type) {
					return createError(INCOMPATABLE_TYPES, "%s not supertype of %s", param.Type, res.Params[i].Type)
				}

				v.Params[i].Type = res.Params[i].Type // reassign param type to parents
			}

			if res.Return != environment.NOTHING_CLASS && !env.ValidSubType(v.Return, res.Return) {
				return createError(INVALID_RETURN_TYPE, "overriding method '%s' in %s has incompatible return type '%s'. parent %s has type '%s'",
					k, child.Type, v.Return, parent.Type, res.Return)
			}
		}
	}
	return nil
}

func checkClasses(classes []ast.Class, env *environment.Environment) *CheckError {
	for _, class := range classes {
		err := checkClass(class, env)
		if err != nil {
			return err
		}
	}
	return nil
}

// extract class vars + check constructor/types and methods
func checkClass(class ast.Class, env *environment.Environment) *CheckError {
	// check constructor types
	// do it
	newEnv := env.NewScope()

	obj := env.GetClass(environment.ObjectType(class.Signature.Name)) // get type environment.Object
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
		if strings.HasPrefix(k, "this.") { // if class environment.Variable add to class
			obj.Variables[k] = newEnv.Vals[k]
		}
	}

	return nil
}

func checkMethods(classes []ast.Class, env *environment.Environment) *CheckError {
	for _, class := range classes {
		err := checkMethod(class, env)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkMethod(class ast.Class, env *environment.Environment) *CheckError {
	methods := class.Body.Methods
	obj := (*env.TypeTable)[environment.ObjectType(class.Signature.Name)] // get type environment.Object

	for _, method := range methods {
		newEnv := env.NewScope()
		newEnv.SetClass(obj.Type) // this will allow methods to access class instance of 'this' and methods
		// populate with params
		objMeth, _ := obj.GetMethod(method.Name)
		for _, arg := range objMeth.Params {
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
		if result.Type != environment.NOTHING_CLASS && objMeth.Return != environment.NOTHING_CLASS { // not sure how to handle non returns in methods
			if !env.ValidSubType(result.Type, objMeth.Return) { // child, parent
				return createError(INVALID_RETURN_TYPE, "incorrect return subtype %s in method %s, wanted %s", result.Type, method.Name, objMeth.Return)
			}
		}
	}
	return nil
}

func setMethods(classes []ast.Class, env *environment.Environment) *CheckError {
	for _, class := range classes {
		err := setMethod(class, env)
		if err != nil {
			return err
		}
	}
	return nil
}

// error for: duplicates, nonexistent types
func setMethod(class ast.Class, env *environment.Environment) *CheckError {
	methods := class.Body.Methods
	obj := (*env.TypeTable)[environment.ObjectType(class.Signature.Name)] // get type environment.Object

	for _, method := range methods {
		sig := environment.MethodSignature{Params: []environment.Variable{}, Return: environment.NOTHING_CLASS} // default to no return
		if _, ok := obj.MethodTable[method.Name]; ok {
			return createError(ALREADY_INITIALIZED, "method name: %s already exists in class", method.Name)
		}

		if method.Name == string(obj.Type) {
			return createError(ALREADY_INITIALIZED, "method %s can't be the same as the class", method.Name)
		}

		// add arguments
		for _, arg := range method.Args {
			if !env.TypeExist(environment.ObjectType(arg.Type)) {
				return createError(CLASS_NOT_EXIST, "type in method params not exist")
			}

			sig.Params = append(sig.Params, environment.Variable{arg.Arg, environment.ObjectType(arg.Type)})
		}

		if method.Typ != "" {
			if !env.TypeExist(environment.ObjectType(method.Typ)) {
				return createError(CLASS_NOT_EXIST, "type '%s' in method return signature '%s' not exist", method.Typ, method.Name)
			}
			sig.Return = environment.ObjectType(method.Typ)
		} // maybe put nothing type if nothing returned

		obj.AddMethod(method.Name, sig)
	}
	return nil
}

func setSignatures(classes []ast.Class, env *environment.Environment) *CheckError {
	for _, class := range classes {
		err := setSignature(class, env)
		if err != nil {
			return err
		}
	}
	return nil
}

func setSignature(class ast.Class, env *environment.Environment) *CheckError {
	sig := class.Signature
	obj := environment.NewObject()

	obj.Type = environment.ObjectType(class.Signature.Name)
	if env.TypeExist(obj.Type) {
		return createError(DUPLICATE_CLASS, "class %s already exists", obj.Type)
	}

	// add arguments to constructor
	for _, arg := range sig.Args {
		obj.Constructor = append(obj.Constructor, environment.Variable{arg.Arg, environment.ObjectType(arg.Type)})
	}

	// if subtype, set super type
	if sig.Extend != nil {
		obj.Parent = environment.ObjectType(sig.Extend.Parent)
	}

	(*env.TypeTable)[obj.Type] = obj // store environment.Object type
	return nil
}

func evalLetStatement(node *ast.LetStatement, env *environment.Environment) (environment.Variable, *CheckError) {
	result := environment.Variable{Name: node.Name.Value}
	right := node.Value

	rightType, err := TypeCheck(right, env)
	if err != nil {
		return result, err
	}

	if node.Kind != "" { // if type explicitly set
		result.Type = environment.ObjectType(node.Kind)
		if !env.TypeExist(result.Type) {
			return environment.Variable{}, createError(CLASS_NOT_EXIST, "class '%s' doesn't exist", result.Type)
		}

		if !env.ValidSubType(rightType.Type, result.Type) {
			return result, createError(INCOMPATABLE_TYPES, "%s not supertype of %s", result.Type, rightType.Type)
		}
	} else {
		result.Type = rightType.Type
	}

	env.Set(result.Name, result.Type) // set environment.Variable in environment
	return result, nil
}

func evalBlockStatement(block *ast.BlockStatement, env *environment.Environment) (environment.Variable, *CheckError) {
	var result environment.Variable

	for _, statement := range block.Statements {
		result, err := TypeCheck(statement, env)
		if err != nil {
			return result, err
		}


		// need to change this to check type return instead. default to environment.NOTHING_CLASS
		fmt.Println()
		if reflect.TypeOf(statement) == reflect.TypeOf(&ast.ReturnStatement{}) {
			return result, nil
		}
	}
	result.Type = environment.NOTHING_CLASS // if no return default to environment.NOTHING_CLASS
	return result, nil
}

func evalIfStatement(node *ast.IfStatement, env *environment.Environment) (environment.Variable, *CheckError) {
	var result environment.Variable
	// check that condition evals to bool type
	result, err := TypeCheck(node.Condition, env)
	if err != nil {
		return result, err
	}

	if result.Type != environment.BOOL_CLASS {
		return result, createError(CONDITION_NOT_BOOL, "condition not evaluate to bool value")
	}
	// create environment for each scope
	newEnv1 := env.NewScope()
	newEnv2 := env.NewScope()

	result1, err := TypeCheck(node.Consequence, newEnv1)
	if err != nil {
		return result1, err
	}

	if node.Alternative == nil { // if no other statement don't bubble up environment.Variable
		return result, nil
	}

	result2, err := TypeCheck(*node.Alternative, newEnv2)
	if err != nil {
		return result2, err
	}
	// compare environments or to current environment
	union := environment.GetUnion(newEnv1, newEnv2)
	for k := range union {
		env.Set(k, union[k])
	}

	// check for 'return'
	if (result1.Type == environment.NOTHING_CLASS || result1.Type == "") && (result2.Type == environment.NOTHING_CLASS || result2.Type == "") {
		return result, nil
	} else if result1.Type == environment.NOTHING_CLASS {
		return result, createError(INVALID_RETURN_TYPE, "not return in all paths") 
	} else if result2.Type == environment.NOTHING_CLASS {
		return result2, createError(INVALID_RETURN_TYPE, "not return in all paths")
	} else if result1.Type != result2.Type {
		return result, createError(INVALID_RETURN_TYPE, "not same types %s and %s", result1.Type, result2.Type)
	} else {
		return result1, nil
	}
	return result, nil
}

func evalWhileStatement(node *ast.WhileStatement, env *environment.Environment) (environment.Variable, *CheckError) {
	var result environment.Variable
	// check that condition evals to bool type
	result, err := TypeCheck(node.Cond, env)
	if err != nil {
		return result, err
	}

	if result.Type != environment.BOOL_CLASS {
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

func evalReturnStatement(node *ast.ReturnStatement, env *environment.Environment) (environment.Variable, *CheckError) {
	return TypeCheck(node.ReturnValue, env)
}

// init class environment.Object or if in class method call, ei: PT(4, 5);
func evalFunctionCall(node *ast.FunctionCall, env *environment.Environment) (environment.Variable, *CheckError) {
	if env.TypeExist(environment.ObjectType(node.Name)) { // if a class
		class := env.GetClass(environment.ObjectType(node.Name))
		if class == nil {
			return environment.Variable{}, createError(CLASS_NOT_EXIST, "class '%s' doesn't exist", node.Name)
		}
		args := class.Constructor

		if len(args) != len(node.Args) {
			return environment.Variable{}, createError(BAD_FUNCTION_CALL, "incorrect amount of arguments for %s wants %d provided %d", node.Name, len(args), len(node.Args))
		}

		// check argument types // do this only for functions

		for i, arg := range args {
			result, err := TypeCheck(node.Args[i], env)
			if err != nil {
				return result, err
			}

			if arg.Type != result.Type {
				return environment.Variable{}, createError(INCOMPATABLE_TYPES, "incorrect argument type %s for Class %s, expected %s on line %d", arg.Type, node.Name, result.Type, node.Token.Pos.Line)
			}
		}

		return environment.Variable{Type: class.Type}, nil // return the type of the class
	}

	if signature, ok := env.GetClassObject().GetMethod(node.Name); ok { // if method
		for i, param := range signature.Params {
			result, err := TypeCheck(node.Args[i], env)
			if err != nil {
				return result, err
			}

			if param.Type != result.Type {
				return environment.Variable{}, createError(INCOMPATABLE_TYPES, "incorrect argument type")
			}
		}
		return environment.Variable{Type: signature.Return, Name: node.Name + "()"}, nil
	}

	return environment.Variable{}, createError(CLASS_NOT_EXIST, "class '%s' doesn't exist", node.Name)
}

// eval something like this class.method()
func evalMethodCall(node *ast.MethodCall, env *environment.Environment) (environment.Variable, *CheckError) {
	class, err := TypeCheck(node.Variable, env) // left can be any expression that returns a type
	if err != nil {
		return environment.Variable{}, err
	}

	signature, ok := env.GetClassMethod(class.Type, node.Method)
	if !ok {
		return environment.Variable{}, createError(METHOD_NOT_EXIST, "method %s not exist in class %s", node.Method, class.Type)
	}

	// recursively check correct arguments provided
	for i, param := range signature.Params {
		result, err := TypeCheck(node.Args[i], env)
		if err != nil {
			return result, err
		}

		if param.Type != result.Type {
			return environment.Variable{}, createError(INCOMPATABLE_TYPES, "incorrect argument type")
		}
	}

	return environment.Variable{Type: signature.Return, Name: node.Method + "()"}, nil
}

func evalPrefixExpression(expr *ast.PrefixExpression, env *environment.Environment) (environment.Variable, *CheckError) {
	e, err := TypeCheck(expr, env)
	if err != nil {
		return e, err
	}
	return e, nil
}

func evalInfixExpression(node *ast.InfixExpression, env *environment.Environment) (environment.Variable, *CheckError) {
	left, err := TypeCheck(node.Left, env)
	if err != nil {
		return left, err
	}

	right, err := TypeCheck(node.Right, env)
	if err != nil {
		return right, err
	}

	if right.Type != left.Type { // maybe should compare if subtypes
		return right, createError(INCOMPATABLE_TYPES, "types %s-%s and %s-%s not work for expression '%s' on line %d",
			left.Type, left.Name, right.Type, right.Name, node.Operator, node.Token.Pos.Line)
	}

	// get least common type
	obj := env.GetClass(env.GetLowestCommonType(right.Type, left.Type))
	// check if method exists in type
	methods := map[string]string{"+": PLUS, "-": MINUS, "==": EQUALS, "<": LESS, ">": MORE, ">=": ATLEAST,
		"<=": ATMOST, "*": TIMES, "/": DIVIDE, "or": OR, "and": AND}

	if _, ok := env.GetClassMethod(obj.Type, methods[node.Operator]); !ok {
		return environment.Variable{}, createError(METHOD_NOT_EXIST, "method %s not exist in class %s on line %d", methods[node.Operator], obj.Type, node.Token.Pos.Line)
	}

	switch node.Operator { // evaluates to a bool
	case "<", ">", "<=", ">=", "==", "!=", "and", "or":
		return environment.Variable{Type: environment.BOOL_CLASS}, nil
	}

	return left, nil
}

func evalInteger(node *ast.IntegerLiteral, env *environment.Environment) (environment.Variable, *CheckError) {
	return environment.Variable{Name: string(node.Token.Lit), Type: environment.INTEGER_CLASS}, nil
}

func evalString(node *ast.StringLiteral, env *environment.Environment) (environment.Variable, *CheckError) {
	return environment.Variable{Name: node.Value, Type: environment.STRING_CLASS}, nil
}

func evalBoolean(node *ast.Boolean, env *environment.Environment) (environment.Variable, *CheckError) {
	return environment.Variable{Name: string(node.Token.Lit), Type: environment.BOOL_CLASS}, nil
}

func evalIdentifier(node *ast.Identifier, env *environment.Environment) (environment.Variable, *CheckError) {
	IdentType, ok := env.Get(node.Value) // check if it has been defined
	if !ok {
		return environment.Variable{}, createError(VARIABLE_NOT_INITIALIZED, "ident %s is not defined on line: %d", node.Value, node.Token.Pos.Line)
	}
	return environment.Variable{Name: node.Value, Type: IdentType}, nil
}

func evalClassVariableCall(node *ast.ClassVariableCall, env *environment.Environment) (environment.Variable, *CheckError) {
	left, err := TypeCheck(node.Expression, env)
	if err != nil {
		return environment.Variable{}, err
	}

	kind, ok := env.GetClassVariable(left.Type, node.Ident) // type, method
	if !ok {
		return environment.Variable{}, createError(VARIABLE_NOT_INITIALIZED, "type %s doesn't have environment.Variable %s", left.Type, node.Ident)
	}
	return environment.Variable{Type: kind, Name: string(node.Token.Lit)}, nil
}

// ei: "string";
func evalExpressionStatement(node *ast.ExpressionStatement, env *environment.Environment) (environment.Variable, *CheckError) {
	return TypeCheck(node.Expression, env)
}

func evalTypeCaseStatement(node *ast.TypecaseStatement, env *environment.Environment) (environment.Variable, *CheckError) {
	var result environment.Variable
	left, err := TypeCheck(node.Expression, env) // evaluate left side
	if err != nil {
		return left, err
	}

	// type check each alt statement
	for i, stmt := range node.TypeAlt {
		if !env.TypeExist(environment.ObjectType(stmt.Kind)) {
			return result, createError(CLASS_NOT_EXIST, "class %s not exist", stmt.Kind)
		}

		newEnv := env.NewScope()
		newEnv.Set(stmt.Value, environment.ObjectType(stmt.Kind)) // inject typecase var into statement

		res, err := TypeCheck(stmt.StmtBlock, newEnv) // default return nothing type in statement block
		if err != nil {
			return res, err
		}

		if environment.ObjectType(stmt.Kind) == res.Type {
			return res, createError(INVALID_RETURN_TYPE, "typecase block %d wanted type %s, instead got %s", i, stmt.Kind, res.Type)
		}
	}

	return result, nil
}
