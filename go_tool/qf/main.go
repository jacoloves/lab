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
	FROM_MODE   = 100 + iota
	WHERE_MODE  = 100 + iota
	JOIN_MODE   = 100 + iota
)

func select_word_formatted(word string, query *string, flg *bool) {
	if strings.ToUpper(word) == "SELECT" {
		*query += word + "\n"
	} else if word == "," {
		*query += "\t" + word + " "
		*flg = true
	} else if *flg {
		*query += word + "\n"
		*flg = false
	} else {
		*query += "\t" + word + "\n"
	}
}

func from_word_formatted(word string, query *string, depth_level int, inquery_flg bool) {
	if depth_level == DEPTH_LEVEL_ONE {
		*query += "\t"
	}

	if strings.ToUpper(word) == "FROM" {
		if inquery_flg {
			*query += word + " "
		} else {
			*query += word + "\n"
		}
	} else if word == "(" {
		*query += word + "\n"
	} else if word == ")" {
		*query += word + "\n"
	} else {
		*query += "\t" + word + "\n"
	}
}

func join_word_formatted(word string, query *string) {
	if strings.ToUpper(word) == "INNER" || strings.ToUpper == "LEFT" {
		*query += word + " "
	} else if strings.ToUpper(word) == "JOIN" {
		*query += word + "\n"
	}
}

func query_formatted(queryes []string) {
	var query_depth_level int
	var no_caluses string
	var current_mode int
	var genetare_query string
	inquery_flg := false
	select_comma_flg := false
	operator_flg := false
	query_cnt := 0
	queryes[len(queryes)-1] = strings.Replace(queryes[len(queryes)-1], ";", "", 1)
	for _, s := range queryes {
		query_depth_level = NO_CALUSE
		switch strings.ToUpper(s) {
		case "SELECT":
			current_mode = SELECT_MODE
			break
		case "FROM":
			fmt.Printf("%s", genetare_query)
			current_mode = FROM_MODE
			genetare_query = ""
			if queryes[query_cnt+1] == "(" {
				inquery_flg = true
			}
			break
		case "INNER", "LEFT":
			fmt.Printf("%s", genetare_query)
			current_mode = JOIN_MODE
			genetare_query = ""
			break
		case "WHERE":
			current_mode = WHERE_MODE
			query_depth_level = DEPTH_LEVEL_ONE
			fmt.Printf("%s", genetare_query)
			genetare_query = ""
			break
		}
		if current_mode == SELECT_MODE {
			select_word_formatted(s, &genetare_query, &select_comma_flg)
		} else if current_mode == FROM_MODE {
			from_word_formatted(s, &genetare_query, query_depth_level, inquery_flg)
		} else if current_mode == JOIN_MODE {

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
				genetare_query += "\t" + no_caluses + "\n"
				no_caluses = ""
				operator_flg = false
			}
			genetare_query += s + "\n"
		}
		query_cnt++
	}
	genetare_query += "\t" + no_caluses + "\n" + ";\n"
	operator_flg = false
	fmt.Printf("%s", genetare_query)
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
