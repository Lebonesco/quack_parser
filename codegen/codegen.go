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

var TMP_COUNT int // track temp count

// when handling class, clean when exist
var symbolTable = map[string]bool{}

func CodeGen(p *ast.Program) (string, error) {
	var b bytes.Buffer 

	err := genClasses(p.Classes, &b)
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

func genClasses(classes []ast.Class, b *bytes.Buffer) error {
	for _, class := range classes {
		genClass(class, b)
	}
	return nil
}

func genClass(class ast.Class, b *bytes.Buffer) error {
	name := class.Signature.Name
	b.WriteString(fmt.Sprintf("\nstruct class_%s_struct;\n", name)) // struct class_Name_struct;
	b.WriteString(fmt.Sprintf("typedef struct class_%s_struct* class_%s\n\n", name, name)) // typedef struct class_Name_struct* class_Name
	
	b.WriteString(fmt.Sprintf("typedef struct obj_%s_struct {\n", name))
	b.WriteString(fmt.Sprintf("\tclass_%s clazz;\n", name))
	// handle class fields
	genClassVariables(class.Body.Statements, b)

	b.WriteString(fmt.Sprintf("} * obj_%s;\n\n", name))

	b.WriteString(fmt.Sprintf("struct class_%s_struct the_class_%s_struct;\n\n", name, name))
	b.WriteString(fmt.Sprintf("struct class_%s_struct {\n", name))
	// handle method table
	genClassMethodTable(class, b)

	b.WriteString("};\n\n")

	// create constructor method
	genClassConstructor(class, b)

	b.WriteString(fmt.Sprintf("extern class_%s the_class_%s\n\n", name, name))
	return nil
}

func genClassMethodTable(class ast.Class, b *bytes.Buffer) {
	name := class.Signature.Name
	// create constructor
	b.WriteString(fmt.Sprintf("\tobj_%s (*constructor) (", name))
	for i, arg := range class.Signature.Args {
		b.WriteString(fmt.Sprintf("obj_%s", arg.Type))
		if i != len(class.Signature.Args) -1 {
			b.WriteString(",")
		}
	}
	b.WriteString(");\n")
	// check if overriding parent method

	// do in correct order
}

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
	b.WriteString(fmt.Sprintf("\tnew_thing->clazz = the_class_%s\n", name))
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
		err := codeGen(stmt, b, stmt.GetEnvironment())
		if err != nil {
			return err
		}
	}

	b.WriteString("\treturn 0;\n")
	b.WriteString("}\n")
	return nil
}

func codeGen(node ast.Node, b *bytes.Buffer, env *environment.Environment) error {
	switch node := node.(type) {
	// Statements
	// case *ast.BlockStatement:
	// 	return genBlockStatement(node, b)
	// case *ast.ReturnStatement:
	// 	return genReturnStatement(node, b)
	// case *ast.IfStatement:
	// 	return genIfStatement(node, b)
	// case *ast.WhileStatement:
	// 	return genWhileStatement(node, b)
	case *ast.ExpressionStatement:
		return genExpressionStatement(node, b, env)
	// case *ast.TypecaseStatement:
	// 	return genTypeCaseStatement(node, b)
	// case *ast.PrefixExpression:
	// 	return genPrefixExpression(node, b)
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
	// case *ast.FunctionCall: // actually a class call, ei PT(1, 2);
	// 	return genFunctionCall(node, b)
	// case *ast.MethodCall: // handle class.method()
	// 	return genMethodCall(node, b)
	// case *ast.ClassVariableCall:
	// 	return genClassVariableCall(node, b)
	}
	return nil
}

func freshTemp() string {
	TMP_COUNT += 1
	return fmt.Sprintf("tmp_%d", TMP_COUNT)
}

func genExpressionStatement(node *ast.ExpressionStatement, b *bytes.Buffer, env *environment.Environment) error {
	err := codeGen(node.Expression, b, env)
	if err != nil {
		return err
	}

	b.WriteRune(';')
	return nil
}

func genInteger(node *ast.IntegerLiteral, b *bytes.Buffer) error {
	b.WriteString("int_literal(" + string(node.Token.Lit) + ")")
	return nil
}

func genString(node *ast.StringLiteral, b *bytes.Buffer) error {
	b.WriteString("str_literal(" + string(node.Token.Lit) + ")")
	return nil
}

func genBoolean(node *ast.Boolean, b *bytes.Buffer) error {
	if node.Value {
		b.WriteString("lit_true")
	} else {
		b.WriteString("lit_false")
	}
	return nil
}

func genIdentifier(node *ast.Identifier, b *bytes.Buffer, env *environment.Environment) error {
	err := InitVar(node.Value, env, b)
	if err != nil {
		return err
	}
	return nil
}

func genLetStatement(node *ast.LetStatement, b *bytes.Buffer, env *environment.Environment) error {
	err := codeGen(node.Name, b, env)
	if err != nil {
		return err
	}

	b.WriteString(" = ")

	err = codeGen(node.Value, b, env)
	if err != nil {
		return err
	}

	b.WriteString(";")
	return nil
}

func genInfixExpression(node *ast.InfixExpression, b *bytes.Buffer, env *environment.Environment) error {
	err := codeGen(node.Left, b, env)
	if err != nil {
		return err
	}

	methods := map[string]string{"+": PLUS, "-": MINUS, "==": EQUALS, "<": LESS, ">": MORE, ">=": ATLEAST,
		"<=": ATMOST, "*": TIMES, "/": DIVIDE, "or": OR, "and": AND}

	b.WriteString(fmt.Sprintf("->clazz->%s(", methods[node.Operator]))

	err = codeGen(node.Right, b, env)
	if err != nil {
		return err
	}

	b.WriteString(")")

	return nil
}

// code generation helpers
// assume everything is init at final type? -- this is bad?
func InitVar(name string, env *environment.Environment, b *bytes.Buffer) error {
	// if not used yet
	// get type
	objType, ok := env.Get(name)
	if !ok {
		return nil
	}
	if _, ok := symbolTable[name]; !ok { // if not already exist init
		b.WriteString(fmt.Sprintf("obj_%s* %s;\n", objType, name)) // obj_Type* name;
	}
	symbolTable[name] = true

	b.WriteString(fmt.Sprintf("%s", name))
	return nil
}