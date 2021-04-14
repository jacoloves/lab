package	main

import (
	"fmt"
	"regexp"
)

const (
	HEADING = `</?h[1-6]>`
	//PARAGRAPH = `</?p>`
	CODE = `</?code>`
	BREAK = `<br\s?/?>`
)

func check_regexp(reg, str string) {
	fmt.Println(regexp.MustCompile(reg).Match([]byte(str)))
}

func reg() {
	regexp_arr := [...] string{HEADING, CODE, BREAK}	
	for _, v := range regexp_arr {
		fmt.Println(v)
	}
}
