package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/mattn/go-tty"
)

var (
	marks                      = []string{"|", "/", "-", "\\"}
	out              io.Writer = os.Stdout
	restart_time     time.Time
	restart_duration time.Duration
	end_time         time.Time
)

func clear() {
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[H")
	fmt.Print("\x1b[32m")
}

func time_drow(elapsed time.Duration) {
	var h, m, s, ms int64

	h = mod(elapsed.Hours(), 24)
	m = mod(elapsed.Minutes(), 60)
	s = mod(elapsed.Seconds(), 60)
	ms = mod(float64(elapsed.Nanoseconds())/1000, 100)

	fmt.Fprint(os.Stdout, fmt.Sprintf("%v:%v:%v:%v\r", h, m, s, ms))
}

func mod(val float64, mod int64) int64 {
	raw := big.NewInt(int64(val))
	return raw.Mod(raw, big.NewInt(mod)).Int64()
}

func message_press_ent() {
	fmt.Println()
	fmt.Printf("\r%s", "Press Enter Stop")
	fmt.Print("\x1b[H")
}

func time_cout(start time.Time, cancel chan struct{}) {
	var elapsed time.Duration
	for {
		select {
		case <-cancel:
			restart_duration = elapsed
			return
		case <-time.After(time.Millisecond):
			//elapsed = time.Since(start)
			end_time += end_time.Add(time.Millisecond)
			elapsed = end_time.Sub(start)
			time_drow(elapsed)
		}
	}
}

func tty_cencer(start time.Time) {
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
		go time_cout(start, cancel)
		if r == 13 {
			close(cancel)
			return
		}
		message_press_ent()
	}
}

func stop_process() {
	message := fmt.Sprintf("\r\x1b[K%s", "🖐stop")
	io.WriteString(out, message)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	//i := 0
	start := time.Now()
	end_time = start

	for {
		clear()
		tty_cencer(start)
		stop_process()
	}

	fmt.Print("\x1b[37m")
}
