package main

import (
	"errors"
	"unicode"
)

type Lexer struct {
	Text        string
	pos         int64
	currentChar byte
}

func NewLexer(text string) *Lexer {
	return &Lexer{text, 0, text[0]}
}

func (l *Lexer) advance() {
	l.pos += 1
	if l.pos < int64(len(l.Text)) {
		l.currentChar = l.Text[l.pos]
	} else {
		l.currentChar = 0
	}
}

func (l *Lexer) makeTokens() (error, []Token) {
	var tokens []Token

	for l.currentChar != 0 {
		chr := l.currentChar
		switch {
		case chr == ' ':
			l.advance()
		case chr == '\t':
			l.advance()
		case chr == '\n':
			l.advance()
		case unicode.IsDigit(rune(chr)):
			err, token := l.makeNumber()
			if err != nil {
				return err, []Token{}
			}
			tokens = append(tokens, token)
			//l.advance()
		case chr == '+':
			tokens = append(tokens, Token{_TT_PLUS, "+"})
			l.advance()
		case chr == '-':
			tokens = append(tokens, Token{_TT_MINUS, "-"})
			l.advance()
		case chr == '*':
			tokens = append(tokens, Token{_TT_MUL, "*"})
			l.advance()
		case chr == '/':
			tokens = append(tokens, Token{_TT_DIV, "/"})
			l.advance()
		case chr == '(':
			tokens = append(tokens, Token{_TT_LPARENT, "("})
			l.advance()
		case chr == ')':
			tokens = append(tokens, Token{_TT_RPARENT, ")"})
			l.advance()
		default:
			return errors.New("IllegalChar"), []Token{}
		}
	}
	return nil, tokens
}

func (l *Lexer) makeNumber() (error, Token) {
	numStr := ""
	dotCount := 0
	for l.currentChar != 0 && (unicode.IsDigit(rune(l.currentChar)) || l.currentChar == '.') {
		if l.currentChar == '.' {
			if dotCount >= 1 {
				return errors.New("More than one dot in the number"), Token{_TT_ERROR, ""}
			}
			dotCount += 1
			numStr += "."
		} else {
			numStr += string(l.currentChar)
		}
		l.advance()
	}
	if dotCount == 0 {
		return nil, Token{_TT_INT, numStr}
	} else {
		return nil, Token{_TT_FLOAT, numStr}
	}

}
