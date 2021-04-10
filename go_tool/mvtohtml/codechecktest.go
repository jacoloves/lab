package main

import (
	"bufio"
	"fmt"
	"os"
)

func code_check(line string, codeline int) int {

	if line == "```" {
		if codeline == 0 {
			codeline = 1
			fmt.Println("<code>")
		} else {
			codeline = 0
			fmt.Println("</code>")
		}

		return codeline
	}

	fmt.Println(line)
	return codeline
}

func main() {

	file := os.Args[1]

	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	code_line := 0
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		code_line = code_check(line, code_line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
