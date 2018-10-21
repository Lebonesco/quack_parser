package ast

import (
	"fmt"
	"strconv"
	"github.com/Lebonesco/quack_parser/token"
)

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return "LetStatement" }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return "ReturnStatement" }

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return "ExpressionStatement" }

// Expressions
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return string(i.Token.Lit) }

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return string(sl.Token.Lit) }

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string { return string(b.Token.Lit) }

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return string(il.Token.Lit) }

func (oe *InfixExpression) expressionNode() {}
func (oe *InfixExpression) TokenLiteral() string { return string(oe.Token.Lit) }

func (oe *IfExpression) expressionNode() {}
func (oe *IfExpression) TokenLiteral() string { return string(oe.Token.Lit) }

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string { return string(fl.Token.Lit) }




func NewProgram(stmts Attrib) (*Program, error) {
	// fmt.Println(stmts)
	// stmts, ok := stmts.([]Statement)
	// if !ok {
	// 	return &Program{}, nil
	// }
	if stmts == nil {
		return &Program{}, nil
	}

	return &Program{Statements: stmts.([]Statement)}, nil
}

func NewStatementList(stmt Attrib) ([]Statement, error) {
	return []Statement{stmt.(Statement)}, nil
}

func AppendStatement(stmtList, stmt Attrib) ([]Statement, error) {
	return append(stmtList.([]Statement), stmt.(Statement)), nil
}

func NewLetStatement(name, value interface{}) (*LetStatement, error) {
	n, ok := name.(*Identifier)
	if !ok {
		return nil, fmt.Errorf("invalid type definition of Identifier. got=%T", name)
	}

	v, ok := value.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type definition of Identifier. got=%T", value)
	}

	return &LetStatement{Name: n, Value: v}, nil
}

func NewExpressionStatement(expr Attrib) (*ExpressionStatement, error) {
		return &ExpressionStatement{Expression: expr.(Expression)}, nil
}

func NewInfixExpression(left, oper, right Attrib) (Expression, error) {
	l, ok := left.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid left expression. got=%T", left)
	}

	o, ok := oper.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("operator invalid token. got=%T", oper)
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid rigth expression. got=%T", right)
	}

	return &InfixExpression{Left: l, Operator: string(o.Lit), Right: r}, nil
}

func NewIntLiteral(integer Attrib) (Expression, error) {
	intLit, _ := integer.(*token.Token)
	value, err := strconv.ParseInt(string(intLit.Lit), 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as integer", string(intLit.Lit))
	}
	return &IntegerLiteral{Token: *intLit, Value: value}, nil
}

func NewStringLiteral(str Attrib) (Expression, error) {
	return &Identifier{Value: string(str.(*token.Token).Lit)}, nil
}

func NewIdentifier(ident Attrib) (*Identifier, error) {
	return &Identifier{Value: string(ident.(*token.Token).Lit)}, nil
}

func NewBool(val Attrib) (Expression, error) {
	return &Boolean{Value: val.(bool)}, nil
}

func NewBoolExpr(left, right Attrib, oper string) (Expression, error) {
	l, ok := left.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid left expression. got=%T", left)
	}

	r, ok := right.(bool)
	if !ok {
		return nil, fmt.Errorf("invalid right expression. got=%T", right)
	}

	return &InfixExpression{Left: l, Operator: oper, Right: &Boolean{Value: r}}, nil
}