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
	if strings.Contains(os.Args[1], "C:\\") {
		mode = ABSOLUTE_MODE
	}

	if mode == ABSOLUTE_MODE {
		for _, v := range os.Args[1:] {
			if strings.Contains(v, "C:\\") && commond_cnt > 0 {
				break
			}
			commond_cnt++
		}
	}

	for _, v := range os.Args[1 : commond_cnt+1] {
		fmt.Println(v)
	}
}
