package main

import (
	"bufio"
	"fmt"
	"os"
)

//first_indent_clauses := [..] string{"SELECT", "FROM", "WHERE"}

func main() {
	filename := os.Args[1]
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf(line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
