package main

import "bufio"

type Command int

const (
	N_COMMAND Command = 0
	A_COMMAND Command = 1
	C_COMMAND Command = 2
	L_COMMAND Command = 3
)

type Parser struct {
	scanner *bufio.Scanner
	text    string
	dest    string
	comp    string
	jump    string
}

type ParserInterface interface {
	HasMoreCommnads() bool
	Advance()
	CommandType() Command
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}
