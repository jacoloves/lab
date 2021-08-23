package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
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

func color() {
	rand.Seed(time.Now().UnixNano())
	rand_num := rand.Intn(7)

	switch rand_num {
	case 0:
		fmt.Print("\x1b[31m")
	case 1:
		fmt.Print("\x1b[32m")
	case 2:
		fmt.Print("\x1b[33m")
	case 3:
		fmt.Print("\x1b[34m")
	case 4:
		fmt.Print("\x1b[35m")
	case 5:
		fmt.Print("\x1b[36m")
	case 6:
		fmt.Print("\x1b[37m")
	}
}

func main() {
	file, err := packr.NewBox("./text").Open("umamusume.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var buffer string

	clear()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == endline {
			color()
			fmt.Println(buffer)
			time.Sleep(5 * time.Second)

			clear()
			buffer = ""
		} else {
			buffer += fmt.Sprintln(scanner.Text())
		}

	}
}
