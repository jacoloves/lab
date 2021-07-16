package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	buf := make([]byte, 'a')

	for {
		n, err := syscall.Read(1, buf)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(int(buf))

		if n == 0 {
			break
		}
	}

}
