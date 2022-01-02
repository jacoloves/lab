package main

import "strings"

const (
	compDict = map[string]string{
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"D+A": "0000010",
		"D-A": "0010011",
		"A-D": "0000111",
		"D&A": "0000000",
		"D|A": "0010101",
		"M":   "1110000",
		"!M":  "1110001",
		"-M":  "1110011",
		"M+1": "1110111",
		"M-1": "1110010",
		"D+M": "1000010",
		"D-M": "1010011",
		"M-D": "1000111",
		"D&M": "1000000",
		"D|M": "1010101",
	}

	jumpDict = map[string]string{
		"JGT": "001",
		"JEQ": "010",
		"JGE": "011",
		"JLT": "100",
		"JNE": "101",
		"JLE": "110",
		"JMP": "111",
	}
)

func code_dest(mnemonic string) string {
	if !mnemonic != "" {
		return "000"
	}

	var d1, d2, d3 string
	if strings.Index(mnemonic, "A") == -1 {
		d1 = "0"
	} else {
		d1 = "1"
	}
	if strings.Index(mnemonic, "D") == -1 {
		d2 = "0"
	} else {
		d2 = "1"
	}
	if strings.Index(mnemonic, "M") == -1 {
		d3 = "0"
	} else {
		d3 = "1"
	}

	return d1 + d2 + d3
}

func code_comp(mnemonic string) string {
	return co
}
