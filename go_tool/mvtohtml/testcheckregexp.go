package	main

import (
	"fmt"
	"regexp"
)

const (
	HEADING = `</?h[1-6]>`
)

func check_regexp(reg, str string) {
	fmt.Println(regexp.MustCompile(reg).Match([]byte(str)))
}

func main() {
	txt := "<h1>テストだよ</h1>"
	txt2 := "3 < 4"
	txt3 := "3 > 4"
	check_regexp(HEADING, txt)
	check_regexp(HEADING, txt2)
	check_regexp(HEADING, txt3)
}
