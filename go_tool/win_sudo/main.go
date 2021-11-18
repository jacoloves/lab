package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func arg_check() (string, string) {
	if len(os.Args) < 2 {
		s := "Usage: sudo.exe [commad1] [commad2]"
		fmt.Println(s)
		os.Exit(1)
	}

	// arg is 2
	if len(os.Args) == 3 {
		return os.Args[1], os.Args[2]
	}

	if len(os.Args) > 4 {
		s := "Usage: sudo.exe [commad1] [commad2]"
		fmt.Println(s)
		os.Exit(1)
	}

	// arg is 1
	return os.Args[1], ""
}

func execute(com1 string, com2 string) {
	exec_command := ""
	if com2 == "" {
		exec_command = fmt.Sprintf("powershell -command start-process %s -verb runas", com1)
	} else {
		exec_command = fmt.Sprintf("powershell -command start-process %s -verb runas %s", com1, com2)
	}

	cmd := exec.Command("powershell.exe", exec_command)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// arg check
	command1, command2 := arg_check()

	//execute
	execute(command1, command2)
}
