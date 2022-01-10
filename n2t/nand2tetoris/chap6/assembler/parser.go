package main

import (
	"bufio"
	"strings"
)

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

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, "", "", "", ""}
}

func (p *Parser) HasMoreCommnads() bool {
	return p.scanner.Scan()
}

func (p *Parser) Advance() {
	p.text = p.scanner.Text()
	if len(p.text) <= 0 {
		return
	}

	tokens := strings.SplitN(p.text, "//", 2)
	if len(tokens) > 0 {
		p.text = tokens[0]
	}

	p.text = strings.TrimSpace(p.text)
}

func (p *Parser) CommandType() Command {
	p.dest = ""
	p.comp = ""
	p.jump = ""
	if len(p.text) == 0 {
		return N_COMMAND
	}
	if p.text[0] == '@' {
		return A_COMMAND
	}
	if p.text[0] == '(' {
		return L_COMMAND
	}
	p.analyze()
	return C_COMMAND
}

func (p *Parser) Symbol() string {
	if strings.HasPrefix(p.text, "@") {
		return strings.TrimPrefix(p.text, "@")
	} else {
		return strings.Trim(p.text, "()")
	}
}

func (p *Parser) analyze() {
	tokens := strings.SplitN(p.text, ";", 2)
	destcomp := tokens[0]
	if len(tokens) == 2 {
		p.jump = tokens[1]
	} else {
		p.jump = ""
	}

	tokens = strings.SplitN(destcomp, "=", 2)
	if len(tokens) == 2 {
		p.dest = tokens[0]
		p.comp = tokens[1]
	} else {
		p.dest = ""
		p.comp = destcomp
	}
}

func (p *Parser) Dest() string {
	return p.dest
}

func (p *Parser) Comp() string {
	return p.comp
}

func (p *Parser) Jump() string {
	return p.Jump
}
