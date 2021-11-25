package main

import (
	"os"
	"regexp"
	"strings"
)

func checkReg(str string) bool {
	checkreg_flg := regexp.MustCompile(`\"+[A-Za-z]+`).Match([]byte(str))

	return checkreg_flg
}

func main() {
	tmp_str := ""
	cnt := 0
	program_str := ""
	for _, v := range os.Args[1:] {
		if strings.Index(v, "C:") == 1 && !checkReg(v) {
			program_str = "\"" + v + "\""
		}

		if strings.Index(v, "C:") == 1 {
			tmp_str = v
		}

		if cnt >= 1 && tmp_str != "" {
			program_str = tmp_str + " " + v
			tmp_str = ""
		}

		cnt++
	}

	//fmt.Println(program_str)
}
