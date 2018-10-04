// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/Lebonesco/quack_scanner/token"
)

const (
	NoState    = -1
	NumStates  = 107
	NumSymbols = 134
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: 'c'
1: 'l'
2: 'a'
3: 's'
4: 's'
5: 'd'
6: 'e'
7: 'f'
8: 'l'
9: 'e'
10: 't'
11: 'e'
12: 'x'
13: 't'
14: 'e'
15: 'n'
16: 'd'
17: 's'
18: 'i'
19: 'f'
20: 'e'
21: 'l'
22: 'i'
23: 'f'
24: 'e'
25: 'l'
26: 's'
27: 'e'
28: 'w'
29: 'h'
30: 'i'
31: 'l'
32: 'e'
33: 'r'
34: 'e'
35: 't'
36: 'u'
37: 'r'
38: 'n'
39: 't'
40: 'y'
41: 'p'
42: 'e'
43: 'c'
44: 'a'
45: 's'
46: 'e'
47: 't'
48: 'r'
49: 'u'
50: 'e'
51: 'f'
52: 'a'
53: 'l'
54: 's'
55: 'e'
56: 'S'
57: 't'
58: 'r'
59: 'i'
60: 'n'
61: 'g'
62: 'I'
63: 'n'
64: 't'
65: 'O'
66: 'b'
67: 'j'
68: 'B'
69: 'o'
70: 'o'
71: 'l'
72: 'e'
73: 'a'
74: 'n'
75: 'a'
76: 'n'
77: 'd'
78: 'o'
79: 'r'
80: 'n'
81: 'o'
82: 't'
83: 'N'
84: 'o'
85: 't'
86: 'h'
87: 'i'
88: 'n'
89: 'g'
90: 'n'
91: 'o'
92: 'n'
93: 'e'
94: '+'
95: '-'
96: '*'
97: '/'
98: '='
99: '='
100: '<'
101: '='
102: '<'
103: '>'
104: '='
105: '>'
106: '{'
107: '}'
108: '='
109: '('
110: ')'
111: ','
112: ';'
113: '.'
114: ':'
115: '_'
116: '/'
117: '/'
118: '\n'
119: '/'
120: '*'
121: '*'
122: '*'
123: '/'
124: '\t'
125: '\n'
126: '\r'
127: ' '
128: '1'-'9'
129: 'A'-'Z'
130: 'a'-'z'
131: '0'-'9'
132: '0'-'9'
133: .
*/
