package main

import (
	"bufio"
	"fmt"
	"os"
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

	sysmbolTable := scabSymbol(NewScanner(scanner))

	fp.Seek(0, 0)
	scanner = bufio.NewScanner(fp)
	ramAddr := 0x0010
	p := NewParser(scanner)
	for p.HasMoreCommnads() {
		p.Advance()
		var output string
		switch p.CommandType() {

		}
	}
}
