package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "###テストだよ"

	slice := strings.Split(str, "")

	if slice[0] == "#" && slice[1] == " " {
		fmt.Println("test")
	} else if slice[0] == "#" && slice[1] == "#" && slice[2] == " " {
		fmt.Println("test2")
	} else if slice[0] == "#" && slice[1] == "#" && slice[2] == "#" && slice[3] == " " {
		fmt.Println("test3")
	}
}
