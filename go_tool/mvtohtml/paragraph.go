package main

import (
	"bufio"
	"os"
	"strings"
)

func paragraph_in(line string, slice *[]string) {

	slice := strings.Split(line, "")

	if slice[0] != "#" || slice[0] != "`" || line != "" {
		if *slice == "" {
			*slice = append(*slice, "<p>")
		}
		*slice = append(*slice, line)
	}

}

func main() {

	mdfile := os.Args[1]

	var slice []string

	fp, err := os.Open(mdfile)
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		paragraph_in(line, &slice)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
