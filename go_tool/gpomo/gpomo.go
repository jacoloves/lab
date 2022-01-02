package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"time"
)

func progressTime() {
	start := time.Now()
	var elapsed time.Duration
	var m, s int64
	for {
		<-time.After(time.Second)
		elapsed = time.Since(start)

		m = mod(elapsed.Minutes(), 60)
		s = mod(elapsed.Seconds(), 60)

		fmt.Fprint(os.Stdout, fmt.Sprintf("%v:%v\r", m, s))
		if m == 25 {
			break
		}
	}
	fmt.Println()
}

func breakTime() {
	start := time.Now()
	var elapsed time.Duration
	var m, s int64
	for {
		<-time.After(time.Second)
		elapsed = time.Since(start)

		m = mod(elapsed.Minutes(), 60)
		s = mod(elapsed.Seconds(), 60)

		fmt.Fprint(os.Stdout, fmt.Sprintf("%v:%v\r", m, s))
		if m == 5 {
			break
		}
	}
	fmt.Println()
}

func mod(val float64, mod int64) int64 {
	raw := big.NewInt(int64(val))
	return raw.Mod(raw, big.NewInt(mod)).Int64()
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println()
		fmt.Println("finish!")
		os.Exit(0)
	}()

	for {
		progressTime()
		fmt.Print("Please Push Enter")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		breakTime()
		fmt.Print("Please Push Enter")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
