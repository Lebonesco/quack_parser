package ast

import (
	_ "errors"
	"fmt"
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
func (rs *ReturnStatement) TokenLiteral() string { return string(rs.Token.Lit) }

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

func NewIntLiteral(integer Attrib) (Expression, error) {
	return &Identifier{Value: string(integer.(*token.Token).Lit)}, nil
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
