package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func createPluginDir(dirName string) {
	fmt.Println(dirName)
}

func inputKey() string {
	var inputKeyStr string
	fmt.Print("[(e)xecute/(s)uspension/(r)ename]: ")
	fmt.Scan(&inputKeyStr)
	if inputKeyStr == "s" || inputKeyStr == "S" {
		fmt.Println("process close")
		exit(1)
	} else if inputKeyStr == "r" || inputKeyStr == "R" {
		return
		continue
	} else if inputKeyStr == "e" || inputKeyStr == "E" {
		pluginPath := vimFilePath + dirName + "/"
		createPluginDir(pluginPath)
		break
	}
}

func main() {
	f, err := os.Open("vimDirectory.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	b, _ := ioutil.ReadAll(f)

	vimFilePath := string(b)
	if !strings.HasSuffix(vimFilePath, "/") {
		vimFilePath = strings.TrimRight(vimFilePath, "\n")
		if strings.HasSuffix(vimFilePath, "\r") {
			vimFilePath = strings.TrimRight(vimFilePath, "\r")
		}
		vimFilePath = vimFilePath + "/"
	}

	for {
		var dirName string
		var inputStr string
		fmt.Print("Please Enter the name of vim plugin: ")
		fmt.Scan(&dirName)
		inputStr == inputKey()
		fmt.Print("[(e)xecute/(s)uspension/(r)ename]: ")
		fmt.Scan(&inputStr)
		if inputStr == "s" || inputStr == "S" {
			fmt.Println("process close")
			exit(1)
		} else if inputStr == "r" || inputStr == "R" {
			continue
		} else if inputStr == "e" || inputStr == "E" {
			pluginPath := vimFilePath + dirName + "/"
			createPluginDir(pluginPath)
			break
		}
	}

}
