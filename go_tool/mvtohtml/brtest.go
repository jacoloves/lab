package main

import (
	"fmt"
	"strings"
)

func main() {

	br_flg := 0
	line := "テストbr   "

	slice := strings.Split(line, "")
	length := len(slice)

	if slice[length-1] == " " && slice[length-2] == " " && slice[length-3] == " " {
		br_flg = 1
	}

	fmt.Println(br_flg)
}
