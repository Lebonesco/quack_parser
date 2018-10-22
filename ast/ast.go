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

func (w *WhileStatement) statementNode() {}
func(w *WhileStatement) TokenLiteral() string { return "WhileStatement" }

func (is *IfStatement) statementNode() {}
func (is *IfStatement) TokenLiteral() string { return string(is.Token.Lit) }

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string { return "BlockStatement" }

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

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string { return string(fl.Token.Lit) }

func (fc *FunctionCall) expressionNode() {}
func (fc *FunctionCall) TokenLiteral() string { return string(fc.Token.Lit) }


// AST builders
func NewProgram(classes, stmts Attrib) (*Program, error) {
	return &Program{Classes: classes.([]Class), Statements: stmts.([]Statement)}, nil
}

func NewStatementList() ([]Statement, error) {
	return []Statement{}, nil
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
		fmt.Println(string(value.(*token.Token).Lit))
		return nil, fmt.Errorf("invalid type definition of Expression. got=%T", value)
	}

	return &LetStatement{Name: n, Value: v}, nil
}

func NewExpressionStatement(expr Attrib) (*ExpressionStatement, error) {
		return &ExpressionStatement{Expression: expr.(Expression)}, nil
}

func NewClass() ([]Class, error) {
	return []Class{}, nil
}

func AppendClass(classList, class Attrib) ([]Class, error) {
	return append(classList.([]Class), class.(Class)), nil
}

func NewClassSignature(name, args, extend Attrib) (*ClassSignature, error) {
	n, ok := name.(string)
	if !ok {
		return nil, fmt.Errorf("invalid type of name. got=%T", name)
	}

	a, ok := args.([]FormalArgs)
	if !ok {
		return nil, fmt.Errorf("invalid type of args. got=%T", args)
	}

	e, ok := extend.(Extends)
	if !ok {
		return nil, fmt.Errorf("invalid type of extend. got=%T", extend)
	}

	return &ClassSignature{Name: n, Args: a, Extend: e}, nil
}

func NewMethod() ([]Method, error) {
	return []Method{}, nil
}

func AppendMethod(methods, method Attrib) ([]Method, error) {
	return append(methods.([]Method), method.(Method)), nil
}


func NewExtends(parent Attrib) (Extends, error) {
	return Extends{Parent: parent.(string)}, nil
}

func NewStatementBlock(stmts Attrib) (*BlockStatement, error) {
	s, ok := stmts.([]Statement)
	if !ok {
		return nil, fmt.Errorf("invalid type of stmts. got=%T", stmts)
	}

	return &BlockStatement{Statements: s}, nil
}

func NewWhileStatement(cond, stmts Attrib) (*WhileStatement, error) {
	c, ok := cond.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid expression for WhileStatement. got=%T", cond)
	}

	b, ok := stmts.(*BlockStatement)
	if !ok {
		return nil, fmt.Errorf("invalid BlockStatement for WhileStatement. got=%T", stmts)
	}

	return &WhileStatement{Cond: c, BlockStatement: b}, nil
}

func NewIfStatement(cond, cons, alt Attrib) (*IfStatement, error) {
	c, ok := cond.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type of cond. got=%T", cond)
	}

	cs, ok := cons.(*BlockStatement)
	if !ok {
		return nil, fmt.Errorf("invalid type of cons. got=%T", cons)
	}

	a, ok := alt.(Statement)
	if !ok {
		return nil, fmt.Errorf("invalid type of alt. got=%T", alt)
	}

	return &IfStatement{Condition: c, Consequence: cs, Alternative: &a}, nil
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

func NewFunctionCall(name, args Attrib) (Expression, error) {
	n, ok := name.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid type of name. got=%T", name)
	}

	a, ok := args.([]Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type of args. got=%T", args)
	}

	return &FunctionCall{Name: string(n.Lit), Args: a}, nil
}

func NewArg() ([]Expression, error) {
	return []Expression{}, nil
}

func AppendArgs(args, arg Attrib) ([]Expression, error) {
	as, ok := args.([]Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type of args. got=%T", args)
	}

	a, ok := arg.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type of arg. got=%T", arg)
	}

	return append(as, a), nil
}

func NewReturnExpression(exp Attrib) (Statement, error) {
	return &ReturnStatement{ReturnValue: exp.(Expression)}, nil
}
