package main

import (
	"fmt"
	"os"
	"strings"
)

func arg_check() string {
	// args check
	if len(os.Args) < 2 {
		s := "Usage: go run main.go [markdownfile]"
		fmt.Println(s)
		os.Exit(1)
	}

	return os.Args[1]
}

func filename_get_check(file string) string {
	// date write
	date_split := strings.Split(file, "/")
	filename := ""
	for _, value := range date_split {
		if (strings.Index(value, ".md")) != -1 {
			value_split := strings.Split(value, ".")
			filename = value_split[0]
		}
	}

	return filename

}

func main() {
	// arg check
	mdfile := arg_check()

	// filename get & check
	filename := filename_get_check(mdfile)

	// execute 
	execute(mdfile, filename)

}
