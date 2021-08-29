package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"net/http"
)

var strSendFlag string
var boolOpt bool

func init() {
	flag.StringVar(&strSendFlag, "s", "", "send message")
	flag.BoolVar(&boolOpt, "w", false, "send message")
}

func main() {
	flag.Parse()

	name := "たなしょー"
	channel := "times_tanasho"

	if !boolOpt {
		jsonStr := `{
			"channel":"` + channel + `",
			"username":"` + name + `",
			"text":"` + strSendFlag + `",
		}`

		req, err := http.NewRequest(
			"POST",
			"xxx"
			bytes.NewBuffer([]byte(jsonStr)),
		)

		if err != nil {
			fmt.Println(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(resp)
		defer resp.Body.Close()
	} else {
		url = "https://slack.com/api/converstios.history"	
		token = "xxx"	

		header={
			"Authorization": "Bearer {}".format(token)
		}
	} 


}
