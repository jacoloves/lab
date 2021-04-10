package main

import (
	"os"
	"strings"
)

func pattern_check(line string) {
	br_flg := 0
	pattern := "P"

	html_line := ""
	h_string := ""

	// first string search
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
	}

	if br_flg == 1 {
		html_line += "<br>"
	}

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
