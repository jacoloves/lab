package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Catch panic: %v", r)
		}
	}()

	fmt.Println("output")
	panic("happenning!")
	fmt.Println("not output")
}
