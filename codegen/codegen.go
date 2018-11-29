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
var Indent string // tracks indentation

// when handling class, clean when exist
var symbolTable = map[string]bool{}

// write c code to buffer
func write(b *bytes.Buffer, code string, args ...interface{}) {
	b.WriteString(fmt.Sprintf(code, args...))
}

func CodeGen(p *ast.Program) (string, error) {
	var b bytes.Buffer 

	parentMethodUnion(p.Classes, p.Env)

	err := genClasses(p.Classes, &b, p.Env)
	if err != nil {
		return b.String(), err
	}

	// generate statements
	err = genMain(p.Statements, &b)
	if err != nil {
		return b.String(), nil
	}

	return b.String(), nil
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
	genClassConstructor(class, b)

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
func genClassConstructor(class ast.Class, b *bytes.Buffer) {
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
	for _, arg := range class.Signature.Args {
		b.WriteString(fmt.Sprintf("\tnew_thing->%s = %s;\n", arg.Arg, arg.Arg))
		
	}
	b.WriteString("\treturn new_thing;\n")
	b.WriteString("}\n\n")
}

func genClassVariables(stmts []ast.Statement, b *bytes.Buffer) {
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
	// case *ast.IfStatement:
	// 	return genIfStatement(node, b, env)
	// case *ast.WhileStatement:
	// 	return genWhileStatement(node, b, env)
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
	write(b, "obj_Int* %s = int_literal(%s);\n", tmp, string(node.Token.Lit))
	return tmp, nil
}

func genString(node *ast.StringLiteral, b *bytes.Buffer) (string, error) {
	tmp := freshTemp()
	write(b, "obj_String %s = str_literal(%s);\n", tmp, string(node.Token.Lit))
	return tmp, nil
}

func genBoolean(node *ast.Boolean, b *bytes.Buffer) (string, error) {
	if node.Value {
		b.WriteString("lit_true")
		return "lit_true", nil
	} else {
		b.WriteString("lit_false")
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
	// if not used yet
	// get type
	objType, ok := env.Get(name)
	if !ok {
		
		res := strings.Split(name, ".")
		if len(res) == 2 { // if class field
			return fmt.Sprintf("%s->%s", res[0], res[1]), nil
		}
		return name, nil
	}
	if _, ok := symbolTable[name]; !ok { // if not already exist init
		b.WriteString(fmt.Sprintf("obj_%s* %s;\n", objType, name)) // obj_Type* name;
	}
	symbolTable[name] = true

	return name, nil
}

func genLetStatement(node *ast.LetStatement, b *bytes.Buffer, env *environment.Environment) (string, error) {
	left, err := codeGen(node.Name, b, env)
	if err != nil {
		return None, err
	}

	right, err := codeGen(node.Value, b, env)
	if err != nil {
		return None, err
	}

	lType, _ := env.Get(node.Name.Value) // might need to change this to get correct value if not ident?
	write(b, Indent + "%s =%s%s;\n", left, convertType(string(lType), node.RightType), right)
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
	// fix type
	write(b, "obj_%s %s = %s->clazz->%s(%s, %s);\n", "Int", tmp, left, methods[node.Operator], left, right)

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

	name := node.Name
	v := freshTemp()
	b.WriteString(fmt.Sprintf("obj_%s %s = the_class_%s->clazz->constructor(", name, v, name))

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
	// check if inherits
	write(b, "%s->clazz->%s(%s", lexpr, method, lexpr)
	// method params 

	b.WriteString(");\n")

	return None, nil
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