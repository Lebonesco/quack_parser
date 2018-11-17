package ast

import (
	"github.com/Lebonesco/quack_parser/token"
)

type Attrib interface{}

// base interface
type Node interface {
	TokenLiteral() string
	//Json() []node // used to contruct json tree
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
	Classes    []Class
	Statements []Statement
}

// Statements
type LetStatement struct {
	Token token.Token // token.Let token
	Name  *Identifier
	Value Expression
	Kind  string
}

type WhileStatement struct {
	Token          token.Token // token while
	Cond           Expression
	BlockStatement *BlockStatement
}

type ReturnStatement struct {
	Token       token.Token // 'return' token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

type TypecaseStatement struct {
	Token      token.Token // 'typecase'
	Expression Expression
	TypeAlt    []TypeAlt
}

type TypeAlt struct {
	Value     string
	Kind      string
	StmtBlock *BlockStatement
}

type Class struct {
	Token     token.Token
	Signature *ClassSignature
	Body      *ClassBody
}

type ClassSignature struct {
	Name   string
	Args   []FormalArgs
	Extend *Extends
}

type FormalArgs struct {
	Arg  string
	Type string
}

type Extends struct {
	Parent string
}

type ClassBody struct {
	Statements []Statement
	Methods    []Method
}

type Method struct {
	Token token.Token
	Name      string
	Args      []FormalArgs
	Typ       string
	StmtBlock *BlockStatement
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
	Token    token.Token // operator token
	Left     Expression
	Operator string
	Right    Expression
}

type PrefixExpression struct {
	Token token.Token
	Value Expression
	Operator string
}

type IfStatement struct {
	Token       token.Token // 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *Statement
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

// initializing object
type FunctionCall struct {
	Token token.Token
	Name  string
	Args  []Expression
}

type MethodCall struct {
	Token token.Token
	Variable Expression // left Side, this will be recursive
	Method string
	Args []Expression
}

type MethodExpression struct {
	Token      token.Token
	Expression Expression
	Ident      string
}

type StringEscapeError struct {
	Token token.Token // string escape error
	Value string
}

// class variable call
type ClassVariableCall struct {
	Token token.Token 
	Expression Expression
	Ident string
}
