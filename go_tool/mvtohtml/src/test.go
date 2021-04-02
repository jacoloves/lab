package main

import "fmt"

func string_format(text string) string {
	return fmt.Sprintf("<title>%s</title>\n", text)
}

func main() {
	title_name := "yyyy_mm_dd.md"
	data := "<!DOCTYPE html>\n"
	data += "<html>\n"
	data += string_format(title_name)
	data += "<meta http-equiv=\"Content-type\" content=\"text/html;charset=UTF-8\">\n"

	fmt.Println(data)
}
