package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gobuffalo/packr"
)

var (
	out io.Writer = os.Stdout

	endline = "*------------------------------------------------*"
)

func clear() {
	fmt.Fprintln(out, "\x1b[2J")
	fmt.Fprintln(out, "\x1b[1;1H")
}

func main() {
	file, err := packr.NewBox("./text").Open("gologo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// var i int
	var buffer string

	clear()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == endline {
			fmt.Println(buffer)
			time.Sleep(5 * time.Second)

			clear()
			buffer = ""
		} else {
			buffer += fmt.Sprintln(scanner.Text())
		}

		// i++
	}

	// for i := 75; i > 0; i-- {
	// 	fmt.Fprintf(out, "\x1b[10;%dH\x1bK", i)
	// 	fmt.Print(`ʕ◔ϖ◔ʔ `)
	// 	time.Sleep(50 * time.Millisecond)
	// }
}
