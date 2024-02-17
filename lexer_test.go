package gojp

import (
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		str    string
		tokens []Token
	}{
		{
			str: "{}",
			tokens: []Token{
				{
					Kind:  LeftBraceSymbol,
					Value: []rune("{"),
				},
				{
					Kind:  RightBraceSymbol,
					Value: []rune("}"),
				},
			},
		},

		{
			str: "{\"info\": [null, false, true, 123, {\"name\": \"rasm\", \"age\": 23 } ]}",
			tokens: []Token{
				{
					Kind:  LeftBraceSymbol,
					Value: []rune("{"),
				},
				{
					Kind:  StringKind,
					Value: []rune("info"),
				},
				{
					Kind:  ColonSymbol,
					Value: []rune(":"),
				},
				{
					Kind:  LeftBracketSymbol,
					Value: []rune("["),
				},
				{
					Kind:  NullKind,
					Value: []rune("null"),
				},
				{
					Kind:  CommaSymbol,
					Value: []rune(","),
				},
				{
					Kind:  BoolKind,
					Value: []rune("false"),
				},
				{
					Kind:  CommaSymbol,
					Value: []rune(","),
				},
				{
					Kind:  BoolKind,
					Value: []rune("true"),
				},
				{
					Kind:  CommaSymbol,
					Value: []rune(","),
				},
				{
					Kind:  NumericKind,
					Value: []rune("123"),
				},
				{
					Kind:  CommaSymbol,
					Value: []rune(","),
				},
				{
					Kind:  LeftBraceSymbol,
					Value: []rune("{"),
				},

				{
					Kind:  StringKind,
					Value: []rune("name"),
				},
				{
					Kind:  ColonSymbol,
					Value: []rune(":"),
				},
				{
					Kind:  StringKind,
					Value: []rune("rasm"),
				},
				{
					Kind:  CommaSymbol,
					Value: []rune(","),
				},
				{
					Kind:  StringKind,
					Value: []rune("age"),
				},
				{
					Kind:  ColonSymbol,
					Value: []rune(":"),
				},
				{
					Kind:  NumericKind,
					Value: []rune("23"),
				},

				{
					Kind:  RightBraceSymbol,
					Value: []rune("}"),
				},

				{
					Kind:  RightBracketSymbol,
					Value: []rune("]"),
				},
				{
					Kind:  RightBraceSymbol,
					Value: []rune("}"),
				},
			},
		},

		{
			str: "{\"hi\" : [1, 2,3]}",
			tokens: []Token{
				{
					Kind:  LeftBraceSymbol,
					Value: []rune("{"),
				},

				{
					Kind:  StringKind,
					Value: []rune("hi"),
				},

				{
					Kind:  ColonSymbol,
					Value: []rune(":"),
				},

				{
					Kind:  LeftBracketSymbol,
					Value: []rune("["),
				},

				{
					Kind:  NumericKind,
					Value: []rune("1"),
				},

				{
					Kind:  CommaSymbol,
					Value: []rune(","),
				},
				{
					Kind:  NumericKind,
					Value: []rune("2"),
				},

				{
					Kind:  CommaSymbol,
					Value: []rune(","),
				},

				{
					Kind:  NumericKind,
					Value: []rune("3"),
				},

				{
					Kind:  RightBracketSymbol,
					Value: []rune("]"),
				},

				{
					Kind:  RightBraceSymbol,
					Value: []rune("}"),
				},
			},
		},
	}

	for _, test := range tests {
		rs := []rune(test.str)
		l := Lexer{
			Str:  rs,
			Cur:  0,
			Size: len(rs),
		}
		ts, err := l.Lex()
		if err != nil {
			t.Error(err)
		}
		if len(ts) != len(test.tokens) {
			t.Error(len(ts), len(test.tokens))
		}
		for i := 0; i < len(ts); i++ {
			if ts[i].Kind != test.tokens[i].Kind {
				t.Error(ts[i], test.tokens[i])
			}
			if string(ts[i].Value) != string(test.tokens[i].Value) {
				t.Error(ts[i], test.tokens[i])
			}
		}
	}
}
