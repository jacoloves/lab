package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if stdin.Text() == "q" {
			break
		}
	}
}
