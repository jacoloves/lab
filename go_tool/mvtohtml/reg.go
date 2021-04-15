package	main

import (
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 2e4f9e2923aad2152e7bbfb41bb844406fa2c5f7
	"regexp"
)

const (
	HEADING = `</?h[1-6]>`
	//PARAGRAPH = `</?p>`
	CODE = `</?code>`
	BREAK = `<br\s?/?>`
)

<<<<<<< HEAD
func check_regexp(reg string, str string) bool {
	flg := regexp.MustCompile(reg).Match([]byte(str))
	return flg
}

func reg(text string) bool {
	reg_flg := false
	regexp_arr := [...] string{HEADING, CODE, BREAK}	
	for _, v := range regexp_arr {
		reg_flg = check_regexp(v, text)
		if reg_flg {
			break
		}
	}

	return reg_flg
=======
func check_regexp(reg, str string) {
	fmt.Println(regexp.MustCompile(reg).Match([]byte(str)))
}

func reg() {
	regexp_arr := [...] string{HEADING, CODE, BREAK}	
	for _, v := range regexp_arr {
		fmt.Println(v)
	}
>>>>>>> 2e4f9e2923aad2152e7bbfb41bb844406fa2c5f7
}
