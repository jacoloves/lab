package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var marks = []string{"|", "/", "-", "\\"}

func clear() {
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[H")
	fmt.Print("\x1b[32m")
}

func time_cout(cnt *int) {
	str_h := ""
	str_m := ""
	str_s := ""
	*cnt++
	time.Sleep(time.Second * 1)
	// hour
	h := *cnt / 3600
	s := *cnt % 3600
	// minute
	m := s / 60
	s = s % 60

	// string mold
	// hour
	if h < 10 {
		str_h = "0" + strconv.Itoa(h)
	} else {
		str_h = strconv.Itoa(h)
	}

	// minute
	if m < 10 {
		str_m = "0" + strconv.Itoa(m)
	} else {
		str_m = strconv.Itoa(m)
	}

	// second
	if s < 10 {
		str_s = "0" + strconv.Itoa(s)
	} else {
		str_s = strconv.Itoa(s)
	}

	// drowing time
	str_time := str_h + ":" + str_m + ":" + str_s

	time_s := fmt.Sprintf("\r%s %s %s", "⏱Progressing", marks[*cnt%4], str_time)
	io.WriteString(os.Stdout, time_s)
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	i := 0

	clear()

	for {
		time_cout(&i)
	}

	fmt.Print("\x1b[37m")
}
