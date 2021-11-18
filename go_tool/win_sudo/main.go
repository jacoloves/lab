package main

import (
	"fmt"
	"os"
)

func arg_check() string {
	if len(os.Args) < 2 {
		s := "Usage: sudo.exe [commad1] [commad2]"
		fmt.Println(s)
		os.Exit(1)
	}

	if len(os.Args) = 2 {
		
	}
		
}

func main() {
	// arg check
	commad1,commad2 := arg_check()
}
