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

const (
	SELECT_MODE = 100
	OTHER_MODE  = iota
)

func select_word_fornatted(word string, query *string) {
		if word == "SELECT" {
				*query += word + "\n" 
		} else if word == "," {
				*query += "\t" + word + " "
		} else if  {
				*query += word
		}
}

func query_formatted(queryes []string) {
	first_indent_clauses := [...]string{"SELECT", "FROM", "WHERE"}
	var query_depth_level int
	var no_caluses string
	var current_mode int
	var genetare_query string
	operator_flg := false
	queryes[len(queryes)-1] = strings.Replace(queryes[len(queryes)-1], ";", "", 1)
	for _, s := range queryes {
		query_depth_level = NO_CALUSE
		for _, v := range first_indent_clauses {
			switch strings.ToUpper(v) {
			case "SELECT":
				query_depth_level = DEPTH_LEVEL_ONE
				current_mode = SELECT_MODE
				break
			case "FROM":
			case "WHERE":
				current_mode = OTHER_MODE
				query_depth_level = DEPTH_LEVEL_ONE
				break
			}
			if query_depth_level == DEPTH_LEVEL_ONE {
				break
			}
		}
		if current_mode == SELECT_MODE {
				select_word_fornatted(s, &genetare_query)	
		} else if query_depth_level == NO_CALUSE {
			switch s {
			case GREATER_THAN:
			case GREATER_THAN_EQUAL:
			case LESS_THAN:
			case LESS_THAN_EQUAL:
			case EQUAL:
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
