package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mattn/go-tty"
)

var (
	marks           = []string{"|", "/", "-", "\\"}
	out   io.Writer = os.Stdout
)

func clear() {
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[H")
	fmt.Print("\x1b[32m")
}

func time_cout(cnt **int, cancel chan struct{}) {
	str_h := ""
	str_m := ""
	str_s := ""
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-cancel:
			return
		case <-ticker.C:
			time.Sleep(time.Second * 1)
			// hour
			h := **cnt / 3600
			s := **cnt % 3600
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

			time_s := fmt.Sprintf("\r%s %s %s", "⏱Progressing", marks[**cnt%4], str_time)
			io.WriteString(out, time_s)
			**cnt++
			fmt.Println()
			fmt.Printf("\r%s", "Press Enter Stop")
			fmt.Print("\x1b[H")
		}
	}
}

func tty_cencer(cnt *int) {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	cancel := make(chan struct{})

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		go time_cout(&cnt, cancel)
		if r == 13 {
			close(cancel)
			return
		}
	}
}

func main() {
	i := 0

	for {
		clear()
		tty_cencer(&i)
		fmt.Println()
		message := fmt.Sprintf("\r%s", "🖐stop")
		io.WriteString(out, message)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	fmt.Print("\x1b[37m")
}
