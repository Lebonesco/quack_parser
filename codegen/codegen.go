package codegen

import (
	"github.com/Lebonesco/quack_parser/ast"
	"github.com/Lebonesco/quack_parser/environment"
	"bytes"
	"fmt"
)

var TMP_COUNT int // track temp count

// when handling class, clean when exist
var symbolTable = map[string]bool{}

func CodeGen(p *ast.Program) (string, error) {
	var b bytes.Buffer 

	// generate statements
	err := genMain(p.Statements, &b)
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
	b.WriteString(fmt.Sprintf("struct class_%s_struct;\n", name)) // struct class_Name_struct;
	b.WriteString(fmt.Sprintf("typedef struct class_%s_struct* class_%s\n\n", name, name)) // typedef struct class_Name_struct* class_Name
	b.WriteString(fmt.Sprintf("typedef struct obj_%s_struct {\n", name))
	b.WriteString(fmt.Sprintf("\tclass_%s clazz;\n", name))
	// handle internals
	


	b.WriteString(fmt.Sprintf("} * obj_%s;", name))
	return nil
}

func genMain(stmts []ast.Statement, b *bytes.Buffer) error {
	b.WriteString("int main() {\n")

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

	b.WriteString(fmt.Sprintf(" %s ", node.Operator))

	err = codeGen(node.Right, b, env)
	if err != nil {
		return err
	}

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