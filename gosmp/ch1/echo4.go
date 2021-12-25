package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("[%d] %s\n", 0, os.Args[0])
	fmt.Printf("[%d] %s\n", 1, strings.Join(os.Args[1:], " "))
}
