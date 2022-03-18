package main

import (
	"encoding/json"
	"fmt"
)

/*
type Holiday struct {
	Year string `json:"year"`
	Data Data   `json: "data"`
}

type Data struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}
*/

type Holiday struct {
	Year string          `json:"year"`
	Data json.RawMessage `json:"data"`
}

func main() {
	var holi Holiday
	smpJson := holiday_2022()
	err := json.Unmarshal([]byte(smpJson), &holi)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", holi.Data[1])
}
