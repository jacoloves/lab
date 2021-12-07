package main

import (
	"fmt"
	"os"
	"strings"
)

const ABSOLUTE_MODE = 1
const DQUOTAION_MODE = 2

func main() {
	mode := 0
	commond_cnt := 0
	program_cmd := ""
	loop_cnt := 0
	if strings.Contains(os.Args[1], "C:\\") || strings.Contains(os.Args[1], "D:\\") {
		mode = ABSOLUTE_MODE
	}

	if mode == ABSOLUTE_MODE {
		for _, v := range os.Args[1:] {
			if (strings.Contains(v, "C:\\") || strings.Contains(os.Args[1], "D:\\")) && commond_cnt > 0 {
				break
			}
			commond_cnt++
		}
	}

	if commond_cnt == 1 {
		mode = DQUOTAION_MODE
	}

	if mode == DQUOTAION_MODE {
		program_cmd = "\"" + os.Args[1] + "\""
	} else {
		for _, v := range os.Args[1:commond_cnt] {
			if loop_cnt == 0 {
				program_cmd += "\"" + v + " "
			} else {
				program_cmd += v + " "
			}

			loop_cnt++
		}
	}
	program_cmd = program_cmd[:len(program_cmd)-1]
	program_cmd += "\""

	fmt.Println(program_cmd)
	fmt.Println(os.Args[commond_cnt])
}
