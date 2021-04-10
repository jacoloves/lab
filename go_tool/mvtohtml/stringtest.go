package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "### テストだよ"
	str2 := ""

	slice := strings.Split(str, "")

	if slice[0] == "#" && slice[1] == " " {
		fmt.Println("test")
	} else if slice[0] == "#" && slice[1] == "#" && slice[2] == " " {
		fmt.Println("test2")
	} else if slice[0] == "#" && slice[1] == "#" && slice[2] == "#" && slice[3] == " " {
		for i := 4; i < len(slice); i++ {
			str2 += slice[i]
		}
	}

	fmt.Println(str2)
}
