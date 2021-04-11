package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pattern_check(line string, codeline int) int {
	br_flg := 0
	pattern := "P"

	html_line := ""
	h_string := ""

	// first string search
	if line == "" && codeline == 0 {
		html_line += "<br>"
		fmt.Println(html_line)
		return codeline
	} else if line == "" && codeline == 1 {
		fmt.Println(html_line)
		return codeline
	} else {
		slice := strings.Split(line, "")
		length := len(slice)

		if slice[0] == "#" && slice[1] == " " {
			pattern = "H1"
			for i := 2; i < length; i++ {
				h_string += slice[i]
			}
		} else if slice[0] == "#" && slice[1] == "#" && slice[2] == " " {
			pattern = "H2"
			for i := 3; i < length; i++ {
				h_string += slice[i]
			}
		} else if slice[0] == "#" && slice[1] == "#" && slice[2] == "#" && slice[3] == " " {
			pattern = "H3"
			for i := 4; i < length; i++ {
				h_string += slice[i]
			}
		}

		// code check
		if line == "```" {
			pattern = "CODE"
		}

		// br check
		if slice[length-1] == " " && slice[length-2] == " " && slice[length-3] == " " {
			br_flg = 1
		}

		// in code tag check
		if pattern != "CODE" && codeline == 1 {
			pattern = "INCODE"
		}

		switch pattern {
		case "H1":
			html_line += "<h1>"
			html_line += h_string
			html_line += "</h1>"
		case "H2":
			html_line += "<h2>"
			html_line += h_string
			html_line += "</h2>"
		case "H3":
			html_line += "<h3>"
			html_line += h_string
			html_line += "</h3>"
		case "P":
			html_line += "<P>"
			if br_flg == 1 {
				for i := 0; i < length-3; i++ {
					html_line += slice[i]
				}
			} else {
				html_line += line
			}
			html_line += "</P>"
		case "CODE":
			if codeline == 0 {
				codeline = 1
				html_line += "<code>"
			} else {
				codeline = 0
				html_line += "</code>"
			}
		case "INCODE":
			html_line += line
		}

		if br_flg == 1 {
			html_line += "<br>"
		}

		fmt.Println(html_line)
		return codeline
	}
}

func converte_html(file string) {
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	code_line := 0
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		code_line = pattern_check(line, code_line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	mdfile := os.Args[1]

	converte_html(mdfile)
}
