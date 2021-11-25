package main

import (
	"fmt"
	"os"
	"regexp"
)

func checkReg(str string) bool {
	checkreg_flg := regexp.MustCompile(`\"+[A-Za-z]+`).Match([]byte(str))

	return checkreg_flg
}

func main() {
	str := os.Args[1]

	fmt_str := ""

	if !checkReg(str) {
		fmt_str = "\"" + str + "\""
	} else {
		fmt_str = str
	}

	fmt.Println(fmt_str)
}
