package ast

import (
	"github.com/Lebonesco/quack_parser/token"
)

type Attrib interface{}

// base interface
type Node interface {
	TokenLiteral() string
	// Json() string // used to contruct json tree
}

// all statement nodes
type Statement interface {
	Node
	statementNode()
}

// all expression nodes
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	//Classes []Class
	Statements []Statement
}

// Statements
type LetStatement struct {
	Token token.Token // token.Let token
	Name *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token token.Token // 'return' token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token token.Token 
	Expression Expression
}

type BlockStatement struct {
	Token token.Token
	Statements []Statement
}

// Expression
type Identifier struct {
	Token token.Token // Token.Ident token
	Value string
}

type Boolean struct {
	Token token.Token
	Value bool 
}

type Op int64

const (
	NOOP = iota
	OR
	AND
)

type BoolExpr struct {
	A bool
	B bool
	Op Op
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type InfixExpression struct {
	Token token.Token // operator token
	Left Expression
	Operator string
	Right Expression 
}

type IfExpression struct {
	Token token.Token // 'if' token
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type FunctionLiteral struct {
	Token token.Token
	Parameters []*Identifier
	Body *BlockStatement
}








