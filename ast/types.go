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
	Classes []Class
	Statements []Statement
}

// Statements
type LetStatement struct {
	Token token.Token // token.Let token
	Name *Identifier
	Value Expression
}

type WhileStatement struct {
	Token token.Token  // token while
	Cond Expression
	BlockStatement *BlockStatement
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

type Class struct {
	Token token.Token
	Signature ClassSignature
	Body ClassBody
}

type ClassSignature struct {
	Name string
	Args []FormalArgs
	Extend Extends
}

type FormalArgs struct {
	Arg string
	Type string
}

type Extends struct {
	Parent string
}

type ClassBody struct {
	Statements []Statement
	Methods []Method
}

type Method struct {
	Name string
	Args []FormalArgs
	Type ValType
	StmtBlock BlockStatement
}

type ValType struct {
	Type string
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

type IfStatement struct {
	Token token.Token // 'if' token
	Condition Expression
	Consequence *BlockStatement
	Alternative *Statement
}

type FunctionLiteral struct {
	Token token.Token
	Parameters []*Identifier
	Body *BlockStatement
}

type FunctionCall struct {
	Token token.Token
	Name string
	Args []Expression
}

