package gojp

import (
	"errors"
	"fmt"
	"strconv"
)

type Parser struct {
	Tokens []Token
	Cur    int
}

func (p *Parser) ParseArray() (interface{}, error) {
	out := []interface{}{}
	t := p.Tokens[p.Cur]
	if TokenKind(t.Value) == RightBracketSymbol {
		p.Cur++
		return out, nil
	}
	for {
		j, err := p.Parse()
		if err != nil {
			return nil, err
		}
		out = append(out, j)
		t = p.Tokens[p.Cur]
		if TokenKind(t.Value) == RightBracketSymbol {
			p.Cur++
			return out, nil
		} else if TokenKind(t.Value) == CommaSymbol {
			p.Cur++
		} else {
			return nil, errors.New(fmt.Sprintf("Expected comma found %v", string(t.Value)))
		}
	}
}

func (p *Parser) ParseObject() (interface{}, error) {
	out := map[interface{}]interface{}{}
	t := p.Tokens[p.Cur]
	if t.Kind == RightBraceSymbol {
		p.Cur++
		return out, nil
	}
	for {
		t = p.Tokens[p.Cur]
		// key (must be string)
		k, err := p.Parse()
		if err != nil {
			return nil, err
		}

		if t.Kind != StringKind {
			return nil, errors.New("Key must be string")
		}

		t = p.Tokens[p.Cur]
		if t.Kind == ColonSymbol {
			p.Cur++
		} else {
			return nil, errors.New(fmt.Sprintf("Expected colon found %v", string(t.Value)))
		}

		// value
		v, err := p.Parse()
		if err != nil {
			return nil, err
		}

		out[k] = v

		t = p.Tokens[p.Cur]
		if t.Kind == RightBraceSymbol {
			p.Cur++
			return out, nil
		} else if t.Kind == CommaSymbol {
			p.Cur++
		} else {
			return nil, errors.New(fmt.Sprintf("Expected comma found %v", string(t.Value)))
		}
	}
}

func (p *Parser) Parse() (interface{}, error) {
	t := p.Tokens[p.Cur]
	p.Cur++
	switch t.Kind {
	case LeftBracketSymbol:
		return p.ParseArray()
	case LeftBraceSymbol:
		return p.ParseObject()
	case NullKind:
		return nil, nil
	case NumericKind:
		v, _ := strconv.ParseFloat(string(t.Value), 64)
		return v, nil
	case StringKind:
		return string(t.Value), nil
	case BoolKind:
		if string(t.Value) == "true" {
			return true, nil
		} else {
			return false, nil
		}
	}
	return nil, errors.New("invalid token")
}
