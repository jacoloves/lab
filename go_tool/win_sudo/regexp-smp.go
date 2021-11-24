package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	str := os.Args[1]
	fmt.Println(regexp.MustCompile(`\"[A-Za-z]+\"`).Match([]byte(str)))
}
