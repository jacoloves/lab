package main

import (
	"os"
)

func pattern_check(line string) {
	//pattern := 'P'

	// first string search

}

func converte_html(file string) {
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	buf := make([]byte, 1024)
	for {
		n, err := fp.Read(buf)
		if n == 0 {
			break
		}

		if err != nil {
			panic(err)
		}

		line := string(buf)

		pattern_check(line)
	}
}

func main() {
	mdfile := os.Args[1]

	converte_html(mdfile)
}
