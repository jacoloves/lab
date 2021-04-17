package main

import "fmt"

func add_slice(slice *[]string) {
	*slice = append(*slice, "pontertest")
}

func main() {
	var slice []string

	for i := 0; i < 10; i++ {
		add_slice(&slice)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(slice[i])
	}
}
