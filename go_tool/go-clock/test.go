package main

import (
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("finish!")
		os.Exit(0)
	}()

	start := time.Now()
	var elapsed time.Duration
	var h, m, s, ms int64
	for {
		<-time.After(time.Millisecond)
		elapsed = time.Since(start)

		h = mod(elapsed.Hours(), 24)
		m = mod(elapsed.Minutes(), 60)
		s = mod(elapsed.Seconds(), 60)
		ms = mod(float64(elapsed.Nanoseconds())/1000, 100)

		fmt.Fprint(os.Stdout, fmt.Sprintf(" %v:%v:%v:%v\r", h, m, s, ms))
	}
}

func mod(val float64, mod int64) int64 {
	raw := big.NewInt(int64(val))
	return raw.Mod(raw, big.NewInt(mod)).Int64()
}
