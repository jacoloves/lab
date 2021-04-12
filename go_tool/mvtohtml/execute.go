package main

import (
	"bufio"
	"os"
	"fmt"
)

func string_format(text string) string {
	return fmt.Sprintf("<title>%s</title>\n", text)
}

func execute(mdfile string, filename string) {
	html_filename := filename + ".html"
	file, err := os.Create(html_filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("<!DOCTYPE html>")
	writer.WriteString("<html>")
	writer.WriteString(string_format(filename))
	writer.WriteString("<meta http-equiv=\"Content-type\" content=\"text/html;charset=UTF-8\">")
	writer.WriteString(css())
	writer.WriteString(generate(mdfile))
	writer.WriteString("</body>")
	writer.WriteString("</html>")
	
	writer.Flush()
}
