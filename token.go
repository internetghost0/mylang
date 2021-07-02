package main

type tokenType uint8

const (
	_TT_ERROR tokenType = 0
	_TT_INT
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
