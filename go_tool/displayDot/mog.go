package main

import "fmt"

var wordColors = map[int]string{
	1:  "\x1b[38;2;255;255;255m",
	2:  "\x1b[38;2;33;32;33m",
	3:  "\x1b[38;2;255;154;173m",
	4:  "\x1b[38;2;239;85;156m",
	5:  "\x1b[38;2;189;69;115m",
	6:  "\x1b[38;2;222;223;222m",
	7:  "\x1b[38;2;0;32;16m",
	8:  "\x1b[38;2;156;154;156m",
	9:  "\x1b[38;2;49;48;173m",
	10: "\x1b[38;2;49;85;239m",
	11: "\x1b[38;2;49;85;156m",
	12: "\x1b[38;2;0;101;140m",
	13: "\x1b[38;2;255;239;222m",
	14: "\x1b[38;2;140;138;140m",
	15: "\x1b[38;2;156;101;156m",
}

var backGroundColors = map[int]string{
	1:  "\x1b[48;2;255;255;255m",
	2:  "\x1b[48;2;33;32;33m",
	3:  "\x1b[48;2;255;154;173m",
	4:  "\x1b[48;2;239;85;156m",
	5:  "\x1b[48;2;189;69;115m",
	6:  "\x1b[48;2;222;223;222m",
	7:  "\x1b[48;2;0;32;16m",
	8:  "\x1b[48;2;156;154;156m",
	9:  "\x1b[48;2;49;48;173m",
	10: "\x1b[48;2;49;85;239m",
	11: "\x1b[48;2;49;85;156m",
	12: "\x1b[48;2;0;101;140m",
	13: "\x1b[48;2;255;239;222m",
	14: "\x1b[48;2;140;138;140m",
	15: "\x1b[48;2;156;101;156m",
}

func clear() {
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[H")
}

func main() {
	clear()
	dotTestArr := [3][44][32]int{
		{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 3, 3, 3, 3, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 3, 3, 3, 3, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 1, 1, 1, 1, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 1, 1, 1, 1, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 3, 3, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 3, 3, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 3, 3, 5, 5, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 3, 3, 5, 5, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 5, 5, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 3, 3, 1, 0, 0, 0},
			{0, 0, 0, 1, 5, 5, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 3, 3, 1, 0, 0, 0},
			{0, 1, 1, 3, 5, 5, 7, 7, 5, 5, 2, 2, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 3, 3, 2, 2, 3, 1, 1, 0},
			{0, 1, 1, 3, 5, 5, 7, 7, 5, 5, 2, 2, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 3, 3, 2, 2, 3, 1, 1, 0},
			{0, 1, 1, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 6, 6, 3, 3, 2, 2, 2, 2, 3, 1, 1, 0},
			{0, 0, 0, 1, 4, 4, 5, 5, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3, 2, 2, 3, 1, 1, 0},
			{0, 0, 0, 1, 4, 4, 5, 5, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3, 2, 2, 3, 1, 1, 0},
			{0, 1, 1, 6, 1, 1, 1, 1, 6, 6, 4, 4, 5, 5, 1, 1, 6, 6, 5, 5, 4, 4, 1, 1, 3, 3, 1, 1, 3, 1, 1, 0},
			{0, 1, 1, 6, 1, 1, 1, 1, 6, 6, 4, 4, 5, 5, 1, 1, 6, 6, 5, 5, 4, 4, 1, 1, 3, 3, 1, 1, 3, 1, 1, 0},
			{0, 1, 1, 5, 1, 1, 6, 6, 4, 4, 5, 5, 1, 1, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 1, 1, 3, 1, 1, 0},
			{0, 1, 1, 5, 1, 1, 6, 6, 4, 4, 5, 5, 1, 1, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 1, 1, 3, 1, 1, 0},
			{0, 0, 0, 1, 1, 1, 4, 4, 5, 5, 5, 5, 1, 1, 4, 4, 4, 4, 6, 6, 1, 1, 6, 6, 1, 1, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 4, 4, 5, 5, 5, 5, 1, 1, 4, 4, 4, 4, 6, 6, 1, 1, 6, 6, 1, 1, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 6, 6, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 1, 1, 4, 4, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 6, 6, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 1, 1, 4, 4, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 4, 4, 4, 4, 5, 5, 4, 4, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 4, 4, 4, 4, 5, 5, 4, 4, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 6, 6, 1, 1, 1, 1, 6, 6, 4, 4, 6, 6, 6, 6, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 6, 6, 1, 1, 1, 1, 6, 6, 4, 4, 6, 6, 6, 6, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 4, 4, 5, 5, 5, 5, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 4, 4, 5, 5, 5, 5, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 3, 3, 3, 3, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 3, 3, 3, 3, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 1, 1, 1, 1, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 1, 1, 1, 1, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 3, 3, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 3, 3, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 3, 3, 5, 5, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 3, 3, 5, 5, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 3, 3, 1, 1, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 3, 3, 1, 1, 0, 0},
			{0, 0, 1, 1, 3, 3, 5, 5, 7, 7, 5, 5, 2, 2, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 3, 3, 3, 3, 1, 1},
			{0, 0, 1, 1, 3, 3, 5, 5, 7, 7, 5, 5, 2, 2, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 3, 3, 3, 3, 1, 1},
			{0, 0, 1, 1, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 6, 6, 3, 3, 2, 2, 3, 3, 1, 1},
			{0, 0, 1, 1, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 6, 6, 3, 3, 2, 2, 3, 3, 1, 1},
			{0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 4, 4, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 2, 2, 3, 3, 1, 1},
			{0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 4, 4, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 2, 2, 3, 3, 1, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 6, 6, 4, 4, 5, 5, 1, 1, 6, 6, 5, 5, 4, 4, 1, 1, 3, 3, 3, 3, 1, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 6, 6, 4, 4, 5, 5, 1, 1, 6, 6, 5, 5, 4, 4, 1, 1, 3, 3, 3, 3, 1, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 6, 6, 4, 4, 5, 5, 1, 1, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 3, 3, 1, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 6, 6, 4, 4, 5, 5, 1, 1, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 3, 3, 1, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 1, 1, 4, 4, 4, 4, 6, 6, 1, 1, 6, 6, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 1, 1, 4, 4, 4, 4, 6, 6, 1, 1, 6, 6, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, 1, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, 1, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 6, 6, 1, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 6, 6, 1, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 6, 6, 1, 1, 1, 1, 4, 4, 4, 4, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 6, 6, 1, 1, 1, 1, 4, 4, 4, 4, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 3, 3, 3, 3, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 3, 3, 3, 3, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 1, 1, 1, 1, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 4, 4, 1, 1, 1, 1, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 3, 3, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 3, 3, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 3, 3, 5, 5, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 3, 3, 5, 5, 4, 4, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 3, 3, 1, 1, 0},
			{0, 0, 0, 1, 1, 5, 5, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 3, 3, 1, 1, 0},
			{0, 1, 1, 3, 3, 5, 5, 7, 7, 5, 5, 2, 2, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 3, 3, 2, 2, 3, 3, 1},
			{0, 1, 1, 3, 3, 5, 5, 7, 7, 5, 5, 2, 2, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 1, 1, 3, 3, 2, 2, 3, 3, 1},
			{0, 1, 1, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 6, 6, 3, 3, 2, 2, 2, 2, 3, 3, 1},
			{0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 1, 1, 3, 3, 2, 2, 3, 3, 1},
			{0, 0, 0, 1, 1, 4, 4, 5, 5, 5, 5, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 1, 1, 3, 3, 2, 2, 3, 3, 1},
			{0, 0, 0, 1, 1, 1, 1, 1, 1, 6, 6, 4, 4, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 1, 1, 3, 3, 1, 1, 3, 3, 1},
			{0, 0, 0, 1, 1, 1, 1, 1, 1, 6, 6, 4, 4, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 1, 1, 3, 3, 1, 1, 3, 3, 1},
			{0, 1, 1, 4, 4, 1, 1, 6, 6, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 3, 3, 1},
			{0, 1, 1, 4, 4, 1, 1, 6, 6, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 1, 1, 1, 1, 3, 3, 1},
			{0, 1, 1, 5, 5, 1, 1, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 1, 1, 0},
			{0, 1, 1, 5, 5, 1, 1, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 0, 0, 1, 1, 0},
			{0, 0, 0, 1, 1, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 6, 6, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 6, 6, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 6, 6, 1, 1, 1, 1, 4, 4, 4, 4, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 6, 6, 1, 1, 1, 1, 4, 4, 4, 4, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, v := range dotTestArr {
		for _, v2 := range v {
			if v2 == 0 {
				fmt.Print(" ")
				fmt.Print("\x1b[0m")
			} else {
				fmt.Print(backGroundColors[v2])
				fmt.Print(wordColors[v2])
				fmt.Print("■")
			}
		}
		fmt.Println()
	}
	fmt.Print("\x1b[0m")
}
