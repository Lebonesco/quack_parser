package codegen

import (
	"github.com/Lebonesco/quack_parser/ast"
	"bytes"
	"fmt"
)

var TMP_COUNT int // track temp count

func CodeGen(p *ast.Program) (string, error) {
	var b bytes.Buffer 

	// generate statements
	err := genMain(p.Statements, &b)
	if err != nil {
		return b.String(), nil
	}

	return b.String(), nil
}

func genClasses(classes []ast.Classes, b *bytes.Buffer) error {
	for _, class := range classes {
		genClass(class, b)
	}
	return nil
}

func genClass(class *ast.Class, b *bytes.Buffer) error {
	name := class.Name
	b.WriteString(fmt.Sprintf("struct class_%s_struct;\n", name)) // struct class_Name_struct;
	b.WriteString(fmt.Sprintf("typedef struct class_%s_struct* class_%s\n\n", name, name)) // typedef struct class_Name_struct* class_Name
	b.WriteString(fmt.Sprintf("typedef struct obj_%s_struct {\n", name))
	b.WriteString(fmt.Sprintf("\tclass_%s clazz;\n", name))

	


	b.WriteString(fmt.Sprintf("} * obj_%s;", name))
}

func genMain(stmts []ast.Statement, b *bytes.Buffer) error {
	b.WriteString("int main() {\n")

	for _, stmt := range stmts {
		err := codeGen(stmt, b)
		if err != nil {
			return err
		}
	}

	b.WriteString("\treturn 0;\n")
	b.WriteString("}\n")
	return nil
}

func codeGen(node ast.Node, b *bytes.Buffer) error {
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
		return genExpressionStatement(node, b)
	// case *ast.TypecaseStatement:
	// 	return genTypeCaseStatement(node, b)
	// case *ast.PrefixExpression:
	// 	return genPrefixExpression(node, b)
	// case *ast.InfixExpression:
	// 	return genInfixExpression(node, b)
	case *ast.IntegerLiteral:
		return genInteger(node, b)
	case *ast.StringLiteral:
		return genString(node, b)
	case *ast.Boolean:
		return genBoolean(node, b)
	// case *ast.Identifier:
	// 	return genIdentifier(node, b)
	// case *ast.LetStatement:
	// 	return genLetStatement(node, b)
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

func genExpressionStatement(node *ast.ExpressionStatement, b *bytes.Buffer) error {
	err := codeGen(node.Expression, b)
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

func genBoolean(node *ast.Boolean, b *byts.Buffer) error {
	if node.Value {
		b.WriteString("lit_true")
	} else {
		b.WriteString("lit_false")
	}
	return nil
}

// code generation helpers

func InitVar(name string, node ast.Node, b *bytes.Buffer) {
	// if not used yet
	// get type
	// obj_Type* name;
}