package gojp

import (
	"errors"
	"log"
	"strconv"
)

const (
	LEN_TRUE  = len("true")
	LEN_FALSE = len("false")
	LEN_NULL  = len("null")
)

type TokenKind string

const (
	StringKind         TokenKind = "STRING"
	NumericKind        TokenKind = "NUMBER"
	BoolKind           TokenKind = "BOOL"
	NullKind           TokenKind = "NULL"
	QuoteSymbol        TokenKind = "\""
	CommaSymbol        TokenKind = ","
	ColonSymbol        TokenKind = ":"
	LeftBracketSymbol  TokenKind = "["
	RightBracketSymbol TokenKind = "]"
	LeftBraceSymbol    TokenKind = "{"
	RightBraceSymbol   TokenKind = "}"
)

type Token struct {
	Value []rune
	Kind  TokenKind
}

type Lexer struct {
	Cur  int
	Size int
	Str  []rune
}

/* TODO
   "\b"  Backspace
   "\f"  Form feed
   "\n"  New line
   "\r"  Carriage return
   "\t"  Tab
   "\""  Double quote
   "\\"  Backslash character
*/

func (l *Lexer) lex_string() ([]rune, error) {
	start := l.Cur
	if TokenKind(l.Str[l.Cur]) != QuoteSymbol {
		return []rune(""), errors.New("error lexing string")
	}
	l.Cur++
	out := []rune{}
	for TokenKind(l.Str[l.Cur]) != QuoteSymbol {
		out = append(out, l.Str[l.Cur])
		if l.Cur >= l.Size {
			l.Cur = start
			return []rune(""), errors.New("no closing quote")
		}
		l.Cur++
	}
	l.Cur++
	return out, nil
}

// float64
func (l *Lexer) lex_number() ([]rune, error) {
	start := l.Cur
	out := []rune{}
	for (l.Str[l.Cur] >= '0' && l.Str[l.Cur] <= '9') || l.Str[l.Cur] == '.' || l.Str[l.Cur] == '-' || l.Str[l.Cur] == 'e' || l.Str[l.Cur] == 'E' {
		out = append(out, l.Str[l.Cur])
		l.Cur++
		if l.Cur >= l.Size {
			l.Cur = start
			return []rune(""), errors.New("parsing number")
		}
	}
	_, err := strconv.ParseFloat(string(out), 64)
	if err != nil {
		l.Cur = start
		return []rune(""), err
	}
	return out, nil
}

func (l *Lexer) lex_boolean() ([]rune, error) {
	start := l.Cur
	if l.Cur+LEN_TRUE <= l.Size && string(l.Str[l.Cur:l.Cur+LEN_TRUE]) == "true" {
		l.Cur += LEN_TRUE
		return []rune("true"), nil
	}
	if l.Cur+LEN_FALSE <= l.Size && string(l.Str[l.Cur:l.Cur+LEN_FALSE]) == "false" {
		l.Cur += LEN_FALSE
		return []rune("false"), nil
	}
	l.Cur = start
	return []rune(""), errors.New("error lexing bolean")
}

func (l *Lexer) lex_null() ([]rune, error) {
	start := l.Cur
	if l.Cur+LEN_NULL <= l.Size && string(l.Str[l.Cur:l.Cur+LEN_NULL]) == "null" {
		l.Cur += LEN_NULL
		return []rune("null"), nil
	}
	l.Cur = start
	return []rune(""), errors.New("error lexing null")
}

func (l *Lexer) Lex() ([]Token, error) {
	out := []Token{}
tg:
	for l.Cur < l.Size {
		flag := false
		if l.Str[l.Cur] == ' ' || l.Str[l.Cur] == '\t' || l.Str[l.Cur] == '\b' || l.Str[l.Cur] == '\n' || l.Str[l.Cur] == '\r' {
			l.Cur++
			continue
		}
		switch TokenKind(l.Str[l.Cur]) {
		case QuoteSymbol:
			e, err := l.lex_string()
			if err != nil {
				// log.Printf("Error: %v", err)
				break tg
			}
			out = append(out, Token{
				Value: e,
				Kind:  StringKind,
			})
			continue tg
		case CommaSymbol:
			out = append(out, Token{
				Value: []rune(CommaSymbol),
				Kind:  CommaSymbol,
			})
			l.Cur++
			continue tg
		case ColonSymbol:
			out = append(out, Token{
				Value: []rune(ColonSymbol),
				Kind:  ColonSymbol,
			})
			l.Cur++
			continue tg
		case LeftBracketSymbol:
			out = append(out, Token{
				Value: []rune(LeftBracketSymbol),
				Kind:  LeftBracketSymbol,
			})
			l.Cur++
			continue tg
		case RightBracketSymbol:
			out = append(out, Token{
				Value: []rune(RightBracketSymbol),
				Kind:  RightBracketSymbol,
			})
			l.Cur++
			continue tg
		case LeftBraceSymbol:
			out = append(out, Token{
				Value: []rune(LeftBraceSymbol),
				Kind:  LeftBraceSymbol,
			})
			l.Cur++
			continue tg
		case RightBraceSymbol:
			out = append(out, Token{
				Value: []rune(RightBraceSymbol),
				Kind:  RightBraceSymbol,
			})
			l.Cur++
			continue tg
		}
		if l.Cur < l.Size {
			b, err := l.lex_boolean()
			if err != nil {
				//log.Printf("Error: %v", err)
			} else {
				out = append(out, Token{
					Value: b,
					Kind:  BoolKind,
				})
				flag = true
			}
		}
		if l.Cur < l.Size {
			nl, err := l.lex_null()
			if err != nil {
				//log.Printf("Error: %v", err)
			} else {
				out = append(out, Token{
					Value: nl,
					Kind:  NullKind,
				})
				flag = true
			}
		}
		if l.Cur < l.Size {
			n, err := l.lex_number()
			if err != nil {
				// log.Printf("Error: %v", err)
			} else {
				out = append(out, Token{
					Value: n,
					Kind:  NumericKind,
				})
				flag = true
			}
		}
		if !flag {
			log.Printf("Error: lexing")
			break
		}
	}
	return out, nil
}
