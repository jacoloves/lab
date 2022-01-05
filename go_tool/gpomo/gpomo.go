package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func clear() {
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[H")
}

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

func outputMusuic() {
	f, err := os.Open("./clock.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
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
		clear()
		progressTime()
		fmt.Print("Please Push Enter")
		outputMusuic()
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		clear()
		breakTime()
		fmt.Print("Please Push Enter")
		outputMusuic()
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
