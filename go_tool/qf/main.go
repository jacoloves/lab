package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NO_CALUSE         = iota
	DEPTH_LEVEL_ONE   = iota
	DEPTH_LEVEL_TWO   = iota
	DEPTH_LEVEL_THREE = iota
)

const (
	GREATER_THAN       = ">"
	GREATER_THAN_EQUAL = ">="
	LESS_THAN          = "<"
	LESS_THAN_EQUAL    = "<="
	EQUAL              = "="
)

func query_formatted(queryes []string) {
	first_indent_clauses := [...]string{"SELECT", "FROM", "WHERE"}
	var query_depth_level int
	var no_caluses string
	operator_flg := false
	queryes[len(queryes)-1] = strings.Replace(queryes[len(queryes)-1], ";", "", 1)
	for _, s := range queryes {
		query_depth_level = NO_CALUSE
		for _, v := range first_indent_clauses {
			if v == strings.ToUpper(s) {
				query_depth_level = DEPTH_LEVEL_ONE
				break
			}
		}
		if query_depth_level == NO_CALUSE {
			if s == GREATER_THAN || s == GREATER_THAN_EQUAL || s == LESS_THAN || s == LESS_THAN_EQUAL || s == EQUAL {
				operator_flg = true
			}
			if operator_flg {
				no_caluses += " " + s
			} else {
				no_caluses += s
			}
		} else if query_depth_level == DEPTH_LEVEL_ONE {
			if no_caluses != "" {
				fmt.Printf("\t%s\n", no_caluses)
				no_caluses = ""
				operator_flg = false
			}
			fmt.Printf("%s\n", s)
		}
	}
	fmt.Printf("\t%s\n", no_caluses)
	operator_flg = false
	fmt.Printf(";\n")
}

func line_space_split(line string) []string {
	query := strings.Split(line, " ")

	return query
}

func main() {
	filename := os.Args[1]
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		query := line_space_split(line)
		query_formatted(query)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
