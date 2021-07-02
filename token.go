package main

type tokenType uint8

const (
	_TT_INT tokenType = 1
	_TT_FLOAT
	_TT_PLUS
	_TT_MINUS
	_TT_MUL
	_TT_DIV
	_TT_LPARENT
	_TT_RPARENT
)

type Token struct {
	_type tokenType
	value string
}
