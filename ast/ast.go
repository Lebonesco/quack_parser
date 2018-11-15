package ast

import (
	"fmt"
	_ "github.com/Lebonesco/quack_parser/errors"
	"github.com/Lebonesco/quack_parser/token"
	"strconv"
)

type node map[string][]node

func debug(fun, expected, v string, got interface{}) error {
	return fmt.Errorf("AST construction error: In function: %s, expected %s for %s. got=%T", fun, expected, v, got)
}

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

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return "ExpressionStatement" }

func (w *WhileStatement) statementNode()       {}
func (w *WhileStatement) TokenLiteral() string { return "WhileStatement" }

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return "IfStatement" }

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return "BlockStatement" }

func (tc *TypecaseStatement) statementNode()       {}
func (tc *TypecaseStatement) TokenLiteral() string { return "TypecaseStatement" }

// Expressions
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return string(i.Token.Lit) }
func (i *Identifier) Json() node {
	return node{i.TokenLiteral(): nil}
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return string(sl.Token.Lit) }

func (se *StringEscapeError) expressionNode()      {}
func (se *StringEscapeError) TokenLiteral() string { return string(se.Token.Lit) }

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return string(b.Token.Lit) }

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return string(il.Token.Lit) }

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return string(oe.Token.Lit) }

func (pf *PrefixExpression) expressionNode() {}
func(pf *PrefixExpression) TokenLiteral() string { return string(pf.Token.Lit) }

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return string(fl.Token.Lit) }

func (fc *FunctionCall) expressionNode()      {}
func (fc *FunctionCall) TokenLiteral() string { return string(fc.Token.Lit) }

func (mc *MethodCall) expressionNode() {}
func (mc *MethodCall) TokenLiteral() string { return string(mc.Token.Lit) }

// AST builders
func NewProgram(classes, stmts Attrib) (*Program, error) {
	c, ok := classes.([]Class)
	if !ok {
		return nil, debug("NewProgram", "[]Class", "classes", classes)
	}

	s, ok := stmts.([]Statement)
	if !ok {
		return nil, debug("NewProgram", "[]Statement", "stmts", stmts)
	}

	return &Program{Classes: c, Statements: s}, nil
}

func NewStatementList() ([]Statement, error) {
	return []Statement{}, nil
}

func AppendStatement(stmtList, stmt Attrib) ([]Statement, error) {
	s, ok := stmt.(Statement)
	if !ok {
		return nil, debug("AppendStatement", "Statement", "stmt", stmt)
	}
	return append(stmtList.([]Statement), s), nil
}

func NewLetStatement(name, kind, value interface{}) (*LetStatement, error) {
	n, ok := name.(*Identifier)
	if !ok {
		return nil, fmt.Errorf("invalid type definition of Identifier. got=%T", name)
	}

	k := &token.Token{}
	if kind != nil {
		var ok bool
		k, ok = kind.(*token.Token)
		if !ok {
			return nil, fmt.Errorf("invalid type definition of Identifier. got=%T", kind)
		}
	}

	v, ok := value.(Expression)
	if !ok {
		return nil, debug("NewLetStatement", "Expression", "value", value)
	}

	return &LetStatement{Name: n, Value: v, Kind: string(k.Lit)}, nil
}

func NewAssignmentStatement(name, value interface{}) (*LetStatement, error) {
	n, ok := name.(*Identifier)
	if !ok {
		return nil, fmt.Errorf("invalid type definition of Identifier. got=%T", name)
	}

	v, ok := value.(Expression)
	if !ok {
		return nil, debug("NewLetStatement", "Expression", "value", value)
	}

	return &LetStatement{Name: n, Value: v}, nil
}

func NewExpressionStatement(expr Attrib) (*ExpressionStatement, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, debug("NewExpressionStatement", "Expression", "expr", expr)
	}
	return &ExpressionStatement{Expression: e}, nil
}

func NewClass() ([]Class, error) {
	return []Class{}, nil
}

func AppendClass(classList, classSignature, classBody Attrib) ([]Class, error) {
	cs, ok := classSignature.(*ClassSignature)
	if !ok {
		return nil, debug("AppendClass", "*ClassSignature", "classSignature", classSignature)
	}

	cb, ok := classBody.(*ClassBody)
	if !ok {
		return nil, debug("AppendClass", "*ClassBody", "classBody", classBody)
	}

	return append(classList.([]Class), Class{Signature: cs, Body: cb}), nil
}

func NewClassSignature(name, args, extend Attrib) (*ClassSignature, error) {
	n, ok := name.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid type of name. got=%T", name)
	}

	a := []FormalArgs{}
	if args != nil {
		var ok bool
		a, ok = args.([]FormalArgs)
		if !ok {
			return nil, debug("NewClassSignature", "[]FormalArgs", "args", args)
		}
	}

	e := &Extends{}
	if extend != nil {
		var ok bool
		e, ok = extend.(*Extends)
		if !ok {
			return nil, debug("NewClassSignature", "Extends", "extend", extend)
		}
	} else {
		e = nil
	}

	return &ClassSignature{Name: string(n.Lit), Args: a, Extend: e}, nil
}

func NewClassBody(stmts, methods Attrib) (*ClassBody, error) {
	s, ok := stmts.([]Statement)
	if !ok {
		return nil, debug("NewClassBody", "[]Statement", "stmts", stmts)
	}

	m, ok := methods.([]Method)
	if !ok {
		return nil, debug("NewClassSignature", "[]Method", "methods", methods)
	}

	return &ClassBody{Statements: s, Methods: m}, nil
}

func NewMethod() ([]Method, error) {
	return []Method{}, nil
}

func AppendMethod(methods, name, args, kind, stmts Attrib) ([]Method, error) {
	n, ok := name.(*token.Token)
	if !ok {
		return nil, debug("AppendMethod", "*token.Token", "name", name)
	}

	a := []FormalArgs{}
	if args != nil {
		var ok bool
		a, ok = args.([]FormalArgs)
		if !ok {
			return nil, debug("AppendMethod", "[]FormalArgs", "args", args)
		}
	}

	k := &token.Token{}
	if kind != nil {
		var ok bool
		k, ok = kind.(*token.Token) // this is optional, add nil case
		if !ok {
			return nil, debug("AppendMethod", "*token.Token", "kind", kind)
		}
	}

	s, ok := stmts.(*BlockStatement)
	if !ok {
		return nil, debug("AppendMethod", "*BlockStatement", "stmts", stmts)
	}

	method := Method{Name: string(n.Lit), Args: a, Typ: string(k.Lit), StmtBlock: s}

	return append(methods.([]Method), method), nil
}

func NewExtends(parent Attrib) (*Extends, error) {
	p, ok := parent.(*token.Token)
	if !ok {
		return nil, debug("NewExtends", "*token.Token", "parent", parent)
	}
	return &Extends{Parent: string(p.Lit)}, nil
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

	var a Statement // could be BlockStatement or IfStatement
	if alt != nil {
		switch alt.(type) {
		case *BlockStatement:
			a, ok = alt.(*BlockStatement)
		case *IfStatement:
			a, ok = alt.(*IfStatement)
		}

		if !ok {
			return nil, fmt.Errorf("invalid type of alt. got=%T", alt)
		}
	}

	return &IfStatement{Condition: c, Consequence: cs, Alternative: &a}, nil
}

func NewInfixExpression(left, oper, right Attrib) (Expression, error) {
	l, ok := left.(Expression)
	if !ok {
		return nil, debug("NewInfixExpression", "Expression", "left", left)
	}

	o, ok := oper.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("operator invalid token. got=%T", oper)
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, debug("NewInfixExpression", "Expression", "right", right)
	}

	return &InfixExpression{Left: l, Operator: string(o.Lit), Right: r, Token: *o}, nil
}

func NewPrefixExpression(oper, value Attrib) (Expression, error) {
	o, ok := oper.(*token.Token)
	if !ok {
		return nil, debug("NewPrefixExpression", "*token.Token", "oper", oper)
	}

	v, ok := value.(Expression)
	if !ok {
		return nil, debug("NewPrefixExpression", "Expression", "value", value)
	}

	return &PrefixExpression{Value: v, Operator: string(o.Lit)}, nil
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
	return &StringLiteral{Value: string(str.(*token.Token).Lit)}, nil
}

func NewIdentifier(ident Attrib) (*Identifier, error) {
	return &Identifier{Value: string(ident.(*token.Token).Lit), Token: *ident.(*token.Token)}, nil
}

func NewBool(val Attrib) (Expression, error) {
	return &Boolean{Value: val.(bool)}, nil
}

func NewBoolExpr(left, right Attrib, oper string) (Expression, error) {
	l, ok := left.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid left expression. got=%T", left)
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, debug("NewBoolExpr", "*bool", "right", right)
	}

	return &InfixExpression{Left: l, Operator: oper, Right: r}, nil
}

func NewFunctionCall(name, args Attrib) (Expression, error) {
	n, ok := name.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid type of name. got=%T", name)
	}

	a := []Expression{}
	if args != nil {
		var ok bool
		a, ok = args.([]Expression)
		if !ok {
			return nil, debug("NewFunctionCall", "[]Expression", "args", args)
		}
	}

	return &FunctionCall{Name: string(n.Lit), Args: a}, nil
}

func NewArg() ([]Expression, error) {
	return []Expression{}, nil
}

func AppendArgs(args, arg Attrib) ([]Expression, error) {
	as, ok := args.([]Expression)
	if !ok {
		return nil, debug("AppendArgs", "[]Expression", "args", args)
	}

	a, ok := arg.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type of arg. got=%T", arg)
	}

	return append(as, a), nil
}

func NewReturnExpression(exp Attrib) (Statement, error) {
	if exp == nil {
		return &ReturnStatement{}, nil
	}
	return &ReturnStatement{ReturnValue: exp.(Expression)}, nil
}

func NewFormalArg() ([]FormalArgs, error) {
	return []FormalArgs{}, nil
}

func AppendFormalArgs(arg, kind, args Attrib) ([]FormalArgs, error) {
	as, ok := args.([]FormalArgs)
	if !ok {
		return nil, debug("AppendFormalArgs", "[]FormalArgs", "args", args)
	}

	a, ok := arg.(*token.Token)
	if !ok {
		return nil, debug("AppendFormalArgs", "*token.Token", "arg", arg)
	}

	k, ok := kind.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid type of kind. got=%T", kind)
	}

	return append(as, FormalArgs{string(a.Lit), string(k.Lit)}), nil
}

// need to fix this up? need to  handle class variable calls
func NewClassVariable(exp, ident Attrib) (Expression, error) {
	_, ok := exp.(Expression)
	if !ok {
		return nil, debug("NewClassVariable", "Expresssion", "exp", exp)
	}

	i, ok := ident.(*token.Token)
	if !ok {
		return nil, debug("NewClassVariable", "*token.Token", "ident", ident)
	}

	return &Identifier{Value: "this." + string(i.Lit), Token: *i}, nil
}

func NewTypeAlt() ([]TypeAlt, error) {
	return []TypeAlt{}, nil
}
// ident.thing()
func NewMethodCall(lexpr, method, args Attrib) (Expression, error) {
	expr, ok := lexpr.(Expression)
	if !ok {
		return nil, debug("NewMethodCall", "Expression", "lexpr", lexpr)
	}

	m, ok := method.(*token.Token)
	if !ok {
		return nil, debug("NewMethodCall", "*token.Token", "method", method)
	}

	a := []Expression{}
	if args != nil {
		var ok bool
		a, ok = args.([]Expression)
		if !ok {
			return nil, debug("NewMethodCall", "[]Expression", "args", args)
		}
	}

	return &MethodCall{Variable: expr, Method: string(m.Lit), Args: a, Token: *m}, nil
}


func AppendTypeAlt(alts, value, kind, stmts Attrib) ([]TypeAlt, error) {
	v, ok := value.(*token.Token)
	if !ok {
		return nil, debug("AppendTypeAlt", "*token.Token", "value", value)
	}

	k, ok := kind.(*token.Token)
	if !ok {
		return nil, debug("AppendTypeAlt", "*token.Token", "kind", kind)
	}

	s, ok := stmts.(*BlockStatement)
	if !ok {
		return nil, debug("AppendTypeAlt", "*BlockStatement", "stmts", stmts)
	}

	alt := TypeAlt{Value: string(v.Lit), Kind: string(k.Lit), StmtBlock: s}
	return append(alts.([]TypeAlt), alt), nil
}

func NewTypecase(expr, typeAlt Attrib) (Statement, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, debug("NewTypecase", "Expression", "expr", expr)
	}

	t, ok := typeAlt.([]TypeAlt)
	if !ok {
		return nil, debug("NewTypecase", "[]TypeAlt", "typeAlt", typeAlt)
	}

	return &TypecaseStatement{Expression: e, TypeAlt: t}, nil
}

// handles unknown tokens
func Unknown(unknown Attrib) (Expression, error) {
	u, ok := unknown.(*token.Token)
	if !ok {
		return nil, debug("Unknown", "*token.Token", "unknown", unknown)
	}

	return &StringLiteral{Token: *u, Value: string(u.Lit)}, nil
}

func NewStringEscapeError(s_error Attrib) (Expression, error) {
	se, ok := s_error.(token.Token)
	if !ok {
		return nil, debug("NewStringEscapeError", "token.Token", "s_error", s_error)
	}

	return &StringEscapeError{Token: se, Value: string(se.Lit)}, nil
}
