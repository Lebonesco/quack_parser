package codegen

import (
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/environment"
	"bytes"
	"strings"
	"fmt"
)


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

const None = "none"

var TMP_COUNT int // track temp count
var LABEL_COUNT int
var EXIT_COUNT int
var Indent string // tracks indentation
var fieldTable = map[string]string{}

// when handling class, clean when exist
var symbolTable = map[string]string{}

// write c code to buffer
func write(b *bytes.Buffer, code string, args ...interface{}) {
	b.WriteString(fmt.Sprintf(code, args...))
}

func CodeGen(p *ast.Program) (bytes.Buffer, error) {
	symbolTable = map[string]string{} // clean should be cleaned from outside runs
	var b bytes.Buffer 
	write(&b, "#include <stdio.h>\n#include <stdlib.h>\n#include \"Builtins.h\"\n\n") // necessary headers

	parentMethodUnion(p.Classes, p.Env)

	err := genClasses(p.Classes, &b, p.Env)
	if err != nil {
		return b, err
	}

	// generate statements
	err = genMain(p.Statements, &b)
	if err != nil {
		return b, nil
	}

	return b, nil
}

func parentMethodUnion(classes []ast.Class, env *environment.Environment) {
	// help track inheritence and base class
	setBaseMethodTypes(classes, env)
	// set parent inherited and overriden methods
	setParentMethods(classes, env)
}

// right now just handles one level, easy to do parents recursively
func setParentMethods(classes []ast.Class, env *environment.Environment) {
	for _, obj := range (*env.TypeTable) { // where are builtins?
		parent := env.GetClass(env.GetClass(obj.Parent).Type)
		// check parent methods for inheritance and overriding
		for i, method := range parent.MethodTable {
			if subMethod, ok := obj.GetMethod(method.Name); ok { // if overridden
				idx, _ := obj.GetMethodIndex(method.Name)
				obj.MethodTable = append(obj.MethodTable[:idx], obj.MethodTable[idx+1:]...) // remove method
				obj.MethodTable = append(obj.MethodTable, environment.MethodSignature{}) // will get overridden on copy
				copy(obj.MethodTable[i+1:], obj.MethodTable[i:])
				obj.MethodTable[i] = subMethod // add method back at correct placement

			} else { // else is inherited
				obj.MethodTable = append(obj.MethodTable, environment.MethodSignature{}) // will get overridden on copy
				copy(obj.MethodTable[i+1:], obj.MethodTable[i:])
				obj.MethodTable[i] = method // insert method at correct placement
			}
		}
	}
}

func setBaseMethodTypes(classes []ast.Class, env *environment.Environment) {
	for _, obj := range (*env.TypeTable) {
		for i, method := range obj.MethodTable {
			method.Base, method.OverrideType = obj.Type, obj.Type
			obj.MethodTable[i] = method
		}
	}
}

func genClasses(classes []ast.Class, b *bytes.Buffer, env *environment.Environment) error {
	for _, class := range classes {
		genClass(class, b, env)
	}
	return nil
}

func genClass(class ast.Class, b *bytes.Buffer, env *environment.Environment) error {
	name := class.Signature.Name
	b.WriteString(fmt.Sprintf("\nstruct class_%s_struct;\n", name)) // struct class_Name_struct;
	b.WriteString(fmt.Sprintf("typedef struct class_%s_struct* class_%s;\n\n", name, name)) // typedef struct class_Name_struct* class_Name
	
	b.WriteString(fmt.Sprintf("typedef struct obj_%s_struct {\n", name))
	b.WriteString(fmt.Sprintf("\tclass_%s clazz;\n", name))
	// handle class fields
	genClassVariables(class.Body.Statements, b)

	b.WriteString(fmt.Sprintf("} * obj_%s;\n\n", name))

	b.WriteString(fmt.Sprintf("struct class_%s_struct the_class_%s_struct;\n\n", name, name))
	b.WriteString(fmt.Sprintf("struct class_%s_struct {\n", name))
	// handle method table
	genClassMethodTable(class, b, env)

	b.WriteString("};\n\n")

	b.WriteString(fmt.Sprintf("extern class_%s the_class_%s;\n\n", name, name))

	// create constructor method
	genClassConstructor(class, b, env)

	// generator class methods
	genClassMethods(class, b, env)

	// create singleton
	createSingleton(class, b, env)

	b.WriteString(fmt.Sprintf("class_%s the_class_%s = &the_class_%s_struct;\n\n",name, name, name))

	return nil
}

func createSingleton(class ast.Class, b *bytes.Buffer, env *environment.Environment) {
	name := class.Signature.Name
	b.WriteString(fmt.Sprintf("struct class_%s_struct the_class_%s_struct = {\n", name, name))
	// add fields
	b.WriteString(fmt.Sprintf("new_%s,\n", name))
	for _, method := range env.GetClass(environment.ObjectType(name)).MethodTable {
		b.WriteString(fmt.Sprintf("%s_method_%s,\n", method.Base, method.Name)) // remove last comma
	}
	b.WriteString("};\n\n")
}

func genClassMethods(class ast.Class, b *bytes.Buffer, env *environment.Environment) {
	name := class.Signature.Name
	methods := class.Body.Methods
	//obj := env.GetClass(environment.ObjectType(name))

	// generate class methods
	for _, method := range methods {
		// generate signature
		b.WriteString(fmt.Sprintf("obj_%s %s_method_%s(", name, name, method.Name))
		b.WriteString(fmt.Sprintf("obj_%s this", name))
		for _, arg := range method.Args {
			b.WriteString(fmt.Sprintf(", obj_%s %s", arg.Type, arg.Arg))
			symbolTable[arg.Arg] = arg.Type // so that can be referenced below? or handled by environment
		}
		b.WriteString(") {\n")
		// generate body
		codeGen(method.StmtBlock, b, env)
		b.WriteString("}\n\n") // end of method
	}
}

// handle sorting and parent insertion
func genClassMethodTable(class ast.Class, b *bytes.Buffer, env *environment.Environment) {
	name := class.Signature.Name
	// create constructor pointer
	b.WriteString(fmt.Sprintf("\tobj_%s (*constructor) (", name))
	for i, arg := range class.Signature.Args {
		b.WriteString(fmt.Sprintf("obj_%s", arg.Type))
		if i != len(class.Signature.Args) -1 {
			b.WriteString(",")
		}
	}
	b.WriteString(");\n") // end of constructor

	obj := env.GetClass(environment.ObjectType(name))

	for _, method := range obj.MethodTable {
		b.WriteString(fmt.Sprintf("\tobj_%s (*%s) (", method.Return, method.Name))
		b.WriteString(fmt.Sprintf("obj_%s", method.OverrideType))
		for _, arg := range method.Params {
			b.WriteString(fmt.Sprintf(",obj_%s", arg.Type))
		}
		b.WriteString(");\n")
	}

	// do in correct order
}

// generates class constructor function
func genClassConstructor(class ast.Class, b *bytes.Buffer, env *environment.Environment) {
	name := class.Signature.Name
	b.WriteString(fmt.Sprintf("obj_%s new_%s(", name, name))
	for i, arg := range class.Signature.Args {
		b.WriteString(fmt.Sprintf("obj_%s %s", arg.Type, arg.Arg))
		if i != len(class.Signature.Args) -1 {
			b.WriteString(",")
		}
	}

	b.WriteString(") {\n")
	b.WriteString(fmt.Sprintf("\tobj_%s new_thing = (obj_%s) malloc(sizeof(struct obj_%s_struct));\n", name, name, name))
	b.WriteString(fmt.Sprintf("\tnew_thing->clazz = the_class_%s;\n", name))
	// how to handle constructor work?

	obj := env.GetClass(environment.ObjectType(class.Signature.Name))
	// for _, arg := range class.Signature.Args {
	// 	b.WriteString(fmt.Sprintf("\tnew_thing->%s = %s;\n", arg.Arg, arg.Arg))
	// }
	for k, _ := range obj.Variables {
		fieldTable[k] = fmt.Sprintf("new_thing->%s",strings.Replace(k, "this.", "", -1))
	}

	for _, stmt := range class.Body.Statements {
		codeGen(stmt, b, env)
	}

	fieldTable = map[string]string{} // clean

	b.WriteString("\treturn new_thing;\n")
	b.WriteString("}\n\n")
}

func genClassVariables(stmts []ast.Statement, b *bytes.Buffer) {
	if len(stmts) <= 0 {
		return
	}

	env := stmts[0].GetEnvironment()
	for k, tp := range env.Vals {
		if strings.HasPrefix(k, "this.") {
			b.WriteString(fmt.Sprintf("\tobj_%s %s;\n", tp, strings.Replace(k, "this.", "", -1))) 
		}
	}
	// get 'this' variables
}

func genMain(stmts []ast.Statement, b *bytes.Buffer) error {
	b.WriteString("\nint main() {\n")

	for _, stmt := range stmts {
		_, err := codeGen(stmt, b, stmt.GetEnvironment())
		if err != nil {
			return err
		}
	}

	b.WriteString("fprintf(stdout, \"\\n--- Terminated SuccessFully (woot!) ---\");\n")
	b.WriteString("\treturn 0;\n")
	b.WriteString("}\n")

	return nil
}

func codeGen(node ast.Node, b *bytes.Buffer, env *environment.Environment) (string, error) {
	switch node := node.(type) {
	// Statements
	case *ast.BlockStatement:
		return genBlockStatement(node, b, env)
	case *ast.ReturnStatement:
		return genReturnStatement(node, b, env)
	case *ast.IfStatement:
		return genIfStatement(node, b, env)
	case *ast.WhileStatement:
		return genWhileStatement(node, b, env)
	case *ast.ExpressionStatement:
		return genExpressionStatement(node, b, env)
	// case *ast.TypecaseStatement:
	// 	return genTypeCaseStatement(node, b, env)
	// case *ast.PrefixExpression:
	// 	return genPrefixExpression(node, b, env)
	case *ast.InfixExpression:
		return genInfixExpression(node, b, env)
	case *ast.IntegerLiteral:
		return genInteger(node, b)
	case *ast.StringLiteral:
		return genString(node, b)
	case *ast.Boolean:
		return genBoolean(node, b)
	case *ast.Identifier:
		return genIdentifier(node, b, env)
	case *ast.LetStatement:
		return genLetStatement(node, b, env)
	case *ast.FunctionCall: // actually a class call, ei PT(1, 2);
		return genFunctionCall(node, b, env)
	case *ast.MethodCall: // handle class.method()
		return genMethodCall(node, b, env)
	case *ast.ClassVariableCall:
		return genClassVariableCall(node, b, env)
	}
	return None, nil
}

func freshTemp() string {
	TMP_COUNT += 1
	return fmt.Sprintf("tmp_%d", TMP_COUNT)
}

func freshLabel() string {
	LABEL_COUNT += 1
	return fmt.Sprintf("label_%d", LABEL_COUNT)
}

func freshExit() string {
	EXIT_COUNT += 1
	return fmt.Sprintf("exit_%d", EXIT_COUNT)
}

func genExpressionStatement(node *ast.ExpressionStatement, b *bytes.Buffer, env *environment.Environment) (string, error) {
	expr, err := codeGen(node.Expression, b, env)
	if err != nil {
		return "", err
	}

	if expr != None {
		write(b, Indent + "%s;\n", expr)
	}
	return "", nil
}

func genInteger(node *ast.IntegerLiteral, b *bytes.Buffer) (string, error) {
	tmp := freshTemp()
	write(b, "obj_Int %s = int_literal(%s);\n", tmp, string(node.Token.Lit))
	return tmp, nil
}

func genString(node *ast.StringLiteral, b *bytes.Buffer) (string, error) {
	tmp := freshTemp()
	write(b, "obj_String %s = str_literal(%s);\n", tmp, string(node.Token.Lit))
	return tmp, nil
}

func genBoolean(node *ast.Boolean, b *bytes.Buffer) (string, error) {
	if node.Value {
		return "lit_true", nil
	} else {
		return "lit_false", nil
	}
	return None, nil
}

func genIdentifier(node *ast.Identifier, b *bytes.Buffer, env *environment.Environment) (string, error) {
	name, err := InitVar(node.Value, env, b)
	if err != nil {
		return None, err
	}

	return name, nil
}

// code generation helpers
// assume everything is init at final type? -- this is bad?
func InitVar(name string, env *environment.Environment, b *bytes.Buffer) (string, error) {
	if res, ok := fieldTable[name]; ok { // handle constructor
		return res, nil
	}

	// if not used yet
	// get type
	if _, ok := symbolTable[name]; ok { // if ident already established
		return name, nil
	}

	objType, ok := env.Get(name)
	if !ok {
		res := strings.Split(name, ".")
		if len(res) == 2 { // if class field
			return fmt.Sprintf("%s->%s", res[0], res[1]), nil
		}
		return name, nil
	}

	if _, ok := symbolTable[name]; !ok { // if not already exist init
		res := strings.Split(name, ".")
		if len(res) == 2 { // if class field
			name = res[0] + "->" + res[1]
			return name, nil
		}
		b.WriteString(fmt.Sprintf("obj_%s %s;\n", objType, name)) // obj_Type* name;
	}
	symbolTable[name] = string(objType)

	return name, nil
}

func genLetStatement(node *ast.LetStatement, b *bytes.Buffer, env *environment.Environment) (string, error) {
	left, err := codeGen(node.Name, b, node.Env)
	if err != nil {
		return None, err
	}

	right, err := codeGen(node.Value, b, env)
	if err != nil {
		return None, err
	}

	kind := node.LeftType
	if k, ok := symbolTable[node.Name.Value]; ok {
		kind = k
	}


	write(b,"%s =%s%s;\n", left, convertType(kind, node.RightType), right)
	return None, nil
}

func convertType(lType, rType string) string {
	if lType != rType {
		return fmt.Sprintf(" (obj_%s) ", lType) // will be true because of previous type checking
	}

	return " "
}

func genBlockStatement(node *ast.BlockStatement, b *bytes.Buffer, env *environment.Environment) (string, error) {
	Indent += "\t" // add indent
	for _, stmt := range node.Statements {
		_, err := codeGen(stmt, b, env)
		if err != nil {
			return None, err
		}
	}
	Indent = Indent[:len(Indent)-1]
	return None, nil
}

func genIfStatement(node *ast.IfStatement, b *bytes.Buffer, env *environment.Environment) (string, error) {
	//condLabel := freshLabel()
	exitLabel := freshExit()
	codeLabel := freshLabel()

	// handle idents defined in both parts
	for _, arg := range node.SharedArgs {
		env.Set(arg.Arg, environment.ObjectType(arg.Type)) // set in environment
		InitVar(arg.Arg, env, b)
	}

	cond, _ := codeGen(node.Condition, b, env)

	write(b, "if (1 == %s->value) {\n", cond)
	write(b, "\tgoto %s;\n", codeLabel)
	write(b, "} else {\n")

	// check if alternative statement
	if *node.Alternative == nil {
		write(b, "goto %s;\n}\n", exitLabel) // if no more conditions
	} 

	if *node.Alternative != nil { // more conditional blocks
		// do stuff
		codeGen(*node.Alternative, b, env)
		write(b, "goto %s;\n}\n", exitLabel)
	}

	write(b, "%s: ;\n", codeLabel)
	_, _ = codeGen(node.Consequence, b, env) // generate code
	write(b, "goto %s;\n\n", exitLabel)

	write(b, "%s: ;\n\n", exitLabel) // end of statement
	return None, nil
}

func genWhileStatement(node *ast.WhileStatement, b *bytes.Buffer, env *environment.Environment) (string, error) {
	test := freshLabel()
	again := freshLabel()
	write(b, "goto %s;\n", test)
	write(b, "%s: ;\n", again)
	// generate code
	codeGen(node.BlockStatement, b, env)
	cond, _ := codeGen(node.Cond, b, env)
	write(b, "%s: ;\n", test)
	write(b, "if (1 == %s->value) {\n", cond)
	write(b, "\tgoto %s;\n}\n", again)
	return None, nil
}

func genReturnStatement(node *ast.ReturnStatement, b *bytes.Buffer, env *environment.Environment) (string, error) {
	// not sure what to do yet
	res, _ := codeGen(node.ReturnValue, b, env)
	write(b, "return %s;\n", res)
	return "", nil
}

func genInfixExpression(node *ast.InfixExpression, b *bytes.Buffer, env *environment.Environment) (string, error) {
	left, err := codeGen(node.Left, b, env)
	if err != nil {
		return None, err
	}

	methods := map[string]string{"+": PLUS, "-": MINUS, "==": EQUALS, "<": LESS, ">": MORE, ">=": ATLEAST,
		"<=": ATMOST, "*": TIMES, "/": DIVIDE, "or": OR, "and": AND}

	right, err := codeGen(node.Right, b, env)
	if err != nil {
		return None, err
	}

	tmp := freshTemp()
	// get type method returns
	method, ok := env.GetClassMethod(environment.ObjectType(node.Type), methods[node.Operator])
	if !ok {
		fmt.Println(node.Type, methods[node.Operator])
		return None, nil
	}
	write(b, "obj_%s %s = %s->clazz->%s(%s, %s);\n", method.Return, tmp, left, methods[node.Operator], left, right)

	return tmp, nil
}

func genFunctionCall(node *ast.FunctionCall, b *bytes.Buffer, env *environment.Environment) (string, error) {
	tmp := make([]string, len(node.Args)) // contain Class parameters

	for i, arg := range node.Args {
		res, err := codeGen(arg, b, env)
		if err != nil {
			return None, err
		}
		tmp[i] = res
	}

	name := node.Name // function or method name
	v := freshTemp()

	// check if name is a reference to a class or functional call inside class
	if env.TypeExist(environment.ObjectType(node.Name)) { 
		b.WriteString(fmt.Sprintf("obj_%s %s = the_class_%s->constructor(", name, v, name))
	} else { // a method
		// get current class
		class := node.Class
		// get return type 
		obj := env.GetClass(environment.ObjectType(class))
		signature, _ := obj.GetMethod(name)
		ret := signature.Return

		write(b, "obj_%s %s = the_class_%s->%s(the_class_%s", ret, v, class, name, class)
	}

	for i, arg := range tmp {
		write(b, arg)
		if i != len(tmp) - 1{
			b.WriteString(",")
		}
	}

	b.WriteString(");\n") // end of Class init

	return v, nil
}

func genMethodCall(node *ast.MethodCall, b *bytes.Buffer, env *environment.Environment) (string, error) {
	method := node.Method
	lexpr, err := codeGen(node.Variable , b, env)
	if err != nil {
		return None, err
	}

	// handle params
	tmp := make([]string, len(node.Args)) // contain Class parameters

	for i, arg := range node.Args {
		res, err := codeGen(arg, b, env)
		if err != nil {
			return None, err
		}
		tmp[i] = res
	}

	register := freshTemp()
	// get type method returns
	meth, ok := env.GetClassMethod(environment.ObjectType(node.LeftType), method)
	if !ok {
		return None, nil
	}

	// check if inherits
		write(b, "obj_%s %s = %s->clazz->%s(%s", meth.Return, register, lexpr, method, lexpr)
	// method params 
	for _, arg := range tmp {
		write(b, ",%s", arg)
	}

	b.WriteString(");\n")

	return register, nil
}

func genClassVariableCall(node *ast.ClassVariableCall, b *bytes.Buffer, env *environment.Environment) (string, error) {
	lexpr, err := codeGen(node.Expression, b, env)
	if err != nil {
		return None, err
	}

	tmp := freshTemp()
	if strings.Contains(node.Ident, "this.") {
		node.Ident = strings.Replace(node.Ident, "this.", "", -1)
	}

	obj := env.GetClass(environment.ObjectType(node.LeftType))
	kind, ok := obj.GetVariableType(node.Ident)
	if !ok {
		return None, fmt.Errorf("class %s not have field %s", node.LeftType, node.Ident)
	}

	write(b, "obj_%s %s = %s->%s;\n", kind, tmp, lexpr, node.Ident)
	return tmp, nil

}