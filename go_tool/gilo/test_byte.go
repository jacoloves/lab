package main

import "fmt"

func main() {
	str := "hello"
	vec := []byte(str)
	fmt.Println(vec[0])
	fmt.Println(vec[:])
	fmt.Println(vec[0:2])
	fmt.Println(vec[2:])
}
