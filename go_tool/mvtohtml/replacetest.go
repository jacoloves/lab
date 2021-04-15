package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "<fff>"
	fmt.Println(strings.Replace(text, "<", "d", -1))
}
