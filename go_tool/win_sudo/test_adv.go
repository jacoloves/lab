package main

import (
	"os"
	"strings"
)

func main() {
	tmp_str := ""
	cnt := 0
	program_str := ""
	for _, v := range os.Args[1:] {
		//if strings.Index(v, "C:") == 1 and
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
