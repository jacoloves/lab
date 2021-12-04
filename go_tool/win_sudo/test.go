package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cnt := len(os.Args) - 1
	arr := make([]string, cnt)
	for i, v := range os.Args[1:] {
		arr[i] = v
	}

	fmt.Println(arr[0])
	exec_command := ""
	exec_command = fmt.Sprintf("powershell -command start-process %s", arr[0])

	exec_command += " -verb runas"

	cmd := exec.Command("powershell.exe", exec_command)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
