// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"github.com/Lebonesco/quack_parser/token"
	"github.com/Lebonesco/quack_parser/util"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Program	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Program : Class Statements	<< X[1], nil >>`,
		Id:         "Program",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Class : ClassSignature ClassBody Class	<<  >>`,
		Id:         "Class",
		NTType:     2,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Class : empty	<<  >>`,
		Id:         "Class",
		NTType:     2,
		Index:      3,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Statements : Statement Statements	<<  >>`,
		Id:         "Statements",
		NTType:     3,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statements : empty	<<  >>`,
		Id:         "Statements",
		NTType:     3,
		Index:      5,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `ClassSignature : class ident lparen FormalArgs rparen Extend	<<  >>`,
		Id:         "ClassSignature",
		NTType:     4,
		Index:      6,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Extend : extends ident	<<  >>`,
		Id:         "Extend",
		NTType:     5,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Extend : empty	<<  >>`,
		Id:         "Extend",
		NTType:     5,
		Index:      8,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `FormalArgs : ident colon ident FormalArgsList	<<  >>`,
		Id:         "FormalArgs",
		NTType:     6,
		Index:      9,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FormalArgs : empty	<<  >>`,
		Id:         "FormalArgs",
		NTType:     6,
		Index:      10,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `FormalArgsList : comma ident colon ident FormalArgsList	<<  >>`,
		Id:         "FormalArgsList",
		NTType:     7,
		Index:      11,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FormalArgsList : empty	<<  >>`,
		Id:         "FormalArgsList",
		NTType:     7,
		Index:      12,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `ClassBody : lbrace Statements Method rbrace	<<  >>`,
		Id:         "ClassBody",
		NTType:     8,
		Index:      13,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Method : def ident lparen FormalArgs rparen Type StatementBlock	<<  >>`,
		Id:         "Method",
		NTType:     9,
		Index:      14,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Method : empty	<<  >>`,
		Id:         "Method",
		NTType:     9,
		Index:      15,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `StatementBlock : lbrace Statements rbrace	<<  >>`,
		Id:         "StatementBlock",
		NTType:     10,
		Index:      16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : if RExpr StatementBlock IfStatement	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      17,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : while RExpr StatementBlock	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : LExpr Type assign RExpr semicolon	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      19,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : RExpr	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : return semicolon	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      21,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Typecase	<<  >>`,
		Id:         "Statement",
		NTType:     11,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : colon ident	<<  >>`,
		Id:         "Type",
		NTType:     12,
		Index:      23,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : empty	<<  >>`,
		Id:         "Type",
		NTType:     12,
		Index:      24,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `IfStatement : elif RExpr StatementBlock IfStatement	<<  >>`,
		Id:         "IfStatement",
		NTType:     13,
		Index:      25,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IfStatement : else StatementBlock	<<  >>`,
		Id:         "IfStatement",
		NTType:     13,
		Index:      26,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IfStatement : empty	<<  >>`,
		Id:         "IfStatement",
		NTType:     13,
		Index:      27,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `LExpr : ident	<<  >>`,
		Id:         "LExpr",
		NTType:     14,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LExpr : RExpr period ident	<<  >>`,
		Id:         "LExpr",
		NTType:     14,
		Index:      29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RExpr : string_literal	<< X[0].(string), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(string), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : RExpr plus Term	<< X[0].(int64) + X[2].(int64), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) + X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : RExpr minus Term	<< X[0].(int64) - X[2].(int64), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) - X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : RExpr atleast Term	<< X[0].(int64) <= X[2].(int64), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) <= X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : RExpr atmost Term	<< X[0].(int64) >= X[2].(int64), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      34,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) >= X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : RExpr lt Term	<< X[0].(int64) < X[2].(int64), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) < X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : RExpr gt Term	<< X[0].(int64) > X[2].(int64), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) > X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : Term	<<  >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RExpr : Bool and Bool	<< X[0].(bool) && X[2].(bool), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      38,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(bool) && X[2].(bool), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : Bool or Bool	<< X[0].(bool) || X[2].(bool), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(bool) || X[2].(bool), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : not Bool	<< !X[1].(bool), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      40,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return !X[1].(bool), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : Bool	<< X[0].(bool), nil >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(bool), nil
		},
	},
	ProdTabEntry{
		String: `RExpr : RExpr period ident lparen ActualArgs rparen	<<  >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      42,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RExpr : ident lparen ActualArgs rparen	<<  >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      43,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RExpr : return RExpr semicolon	<<  >>`,
		Id:         "RExpr",
		NTType:     15,
		Index:      44,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term mul Factor	<< X[0].(int64) * X[2].(int64), nil >>`,
		Id:         "Term",
		NTType:     16,
		Index:      45,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) * X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `Term : Term div Factor	<< X[0].(int64) / X[2].(int64), nil >>`,
		Id:         "Term",
		NTType:     16,
		Index:      46,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) / X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id:         "Term",
		NTType:     16,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : lparen RExpr rparen	<< X[1], nil >>`,
		Id:         "Factor",
		NTType:     17,
		Index:      48,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Factor : int	<< util.IntValue(X[0].(*token.Token).Lit) >>`,
		Id:         "Factor",
		NTType:     17,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return util.IntValue(X[0].(*token.Token).Lit)
		},
	},
	ProdTabEntry{
		String: `Bool : true	<< true, nil >>`,
		Id:         "Bool",
		NTType:     18,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true, nil
		},
	},
	ProdTabEntry{
		String: `Bool : false	<< false, nil >>`,
		Id:         "Bool",
		NTType:     18,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `ActualArgs : RExpr ArgsList	<<  >>`,
		Id:         "ActualArgs",
		NTType:     19,
		Index:      52,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ActualArgs : empty	<<  >>`,
		Id:         "ActualArgs",
		NTType:     19,
		Index:      53,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `ArgsList : comma RExpr ArgsList	<<  >>`,
		Id:         "ArgsList",
		NTType:     20,
		Index:      54,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ArgsList : empty	<<  >>`,
		Id:         "ArgsList",
		NTType:     20,
		Index:      55,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Typecase : typecase RExpr lbrace TypeAlternative rbrace	<<  >>`,
		Id:         "Typecase",
		NTType:     21,
		Index:      56,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TypeAlternative : ident colon ident StatementBlock TypeAlternative	<<  >>`,
		Id:         "TypeAlternative",
		NTType:     22,
		Index:      57,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TypeAlternative : empty	<<  >>`,
		Id:         "TypeAlternative",
		NTType:     22,
		Index:      58,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
}
