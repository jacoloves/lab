package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: Assembler foo.asm")
		os.Exit(1)
	}
	filename := os.Args[1]

	if !strings.HasSuffix(filename, ".asm") {
		fmt.Println("Usage: Assembler foo.asm")
		os.Exit(2)
	}

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("Err! The file cannot open")
		os.Exit(3)
	}

	scanner := bufio.NewScanner(fp)
	hackfile := strings.TrimSuffix(filename, ".asm") + ".hack"
	writerfp, err := os.OpenFile(hackfile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Err! .hack file couldn't create.")
		os.Exit(4)
	}
	writer := bufio.NewWriter(writerfp)

	sysmbolTable := scanSymbol(NewScanner(scanner))

	fp.Seek(0, 0)
	scanner = bufio.NewScanner(fp)
	ramAddr := 0x0010
	p := NewParser(scanner)
	for p.HasMoreCommnads() {
		p.Advance()
		var output string
		switch p.CommandType() {
		case A_COMMAND:
			output = "0"
			symbol := p.Symbol
			addr, err := strconv.Atoi(symbol)
			if err == nil {
				output += int2bin(addr)
			} else {
				if sysmbolTable.Contains(sysmbol) {
					addr = sysmbolTable.GetAddress(symbol)
					output += int2bin(addr)
				} else {
					sysmbolTable.AddEntry(symbol, ramAddr)
					output += int2bin(ramAddr)
					ramAddr++
				}
			}
			fmt.Fprintln(writer, output)
		case C_COMMAND:
			output = "111"
			comp, err := CodeComp(p.Comp())
			if err != nil {
				os.Exit(5)
			}
			output += comp
			dest, err := CodeDest(p.Dest())
			if err != nil {
				os.Exit(5)
			}
			output += dest
			jump, err := CodeJump(p.Jump())
			if err != nil {
				os.Exit(5)
			}
			output += jump
			fmt.Fprintln(writer, output)

		case L_COMMAND:
			// do nothing
		}
		writer.Flush()
	}
}

func int2bin(num int) string {
	var bin string
	for i := 1 << 14; i > 0; i = i >> 1 {
		if i&num != 0 {
			bin += "1"
		} else {
			bin += "0"
		}
	}
	return bin
}

func scanSymbol(p *Parser) SymbolTable {
	st := NewSymbolTable()
	romAddr := 0

	for p.HasMoreCommnads() {
		p.Advance()
		switch p.CommandType() {
		case A_COMMAND, C_COMMAND:
			romAddr++
		case L_COMMAND:
			st.AddEntry(p.Symbol(), romAddr)
		}
	}

	return st
}
