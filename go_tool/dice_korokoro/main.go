package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(6) + 1
	switch randNum {
	case 6:
		dice := 0x2685
		fmt.Printf("%c %d\n", dice, randNum)
	case 5:
		dice := 0x2684
		fmt.Printf("%c %d\n", dice, randNum)
	case 4:
		dice := 0x2683
		fmt.Printf("%c %d\n", dice, randNum)
	case 3:
		dice := 0x2682
		fmt.Printf("%c %d\n", dice, randNum)
	case 2:
		dice := 0x2681
		fmt.Printf("%c %d\n", dice, randNum)
	case 1:
		dice := 0x2680
		fmt.Printf("%c %d\n", dice, randNum)
	}
}
