package gojp

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		str    string
		output interface{}
	}{
		{
			str:    "{}",
			output: map[interface{}]interface{}{},
		},

		{
			str: "{\"info\": [null, false, true, 123, {\"name\": \"rasm\", \"age\": 23 } ]}",
			output: map[interface{}]interface{}{
				"info": []interface{}{nil, false, true, 123, map[interface{}]interface{}{"name": "rasm", "age": 23}}},
		},

		{
			str:    "{\"hi\" : [1, 2,3]}",
			output: map[interface{}]interface{}{"hi": []interface{}{1, 2, 3}},
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

		p := Parser{
			Tokens: ts,
			Cur:    0,
		}

		ps, err := p.Parse()
		if err != nil {
			t.Error(err)
		}

		if fmt.Sprint(test.output) != fmt.Sprint(ps) {
			t.Error("test failed")
		}
	}
}
