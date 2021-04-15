package main

import (
	"regexp"
)

const (
	HEADING = `</?h[1-6]>`
	//PARAGRAPH = `</?p>`
	CODE  = `</?code>`
	BREAK = `<br\s?/?>`
)

func check_regexp(reg string, str string) bool {
	flg := regexp.MustCompile(reg).Match([]byte(str))
	return flg
}

func reg(text string) bool {
	reg_flg := false
	regexp_arr := [...]string{HEADING, CODE, BREAK}
	for _, v := range regexp_arr {
		reg_flg = check_regexp(v, text)
		if reg_flg {
			break
		}
	}

	return reg_flg
}
