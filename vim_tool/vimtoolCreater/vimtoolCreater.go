package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func createFile(fileName string) {
	fmt.Printf("Progress touch %s", fileName)

	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(fileName, currentTime, currentTime)
		if err != nil {
			fmt.Println(err)
		}
	}

	time.Sleep(time.Second * 1)
	fmt.Print("......[OK]\n")
}

func createDir(dirName string) {
	fmt.Printf("Progress mkdir %s", dirName)

	if err := os.MkdirAll(dirName, 0755); err != nil {
		time.Sleep(time.Second * 1)
		fmt.Print("......[NG]\n")
		fmt.Println(err)
	}

	time.Sleep(time.Second * 1)
	fmt.Print("......[OK]\n")
}

func createPluginDir(dirName string, vimFilePath string) {
	var absDirName string
	var absAutoLoadDirName string
	var absPluginDirName string
	var absDocDirName string
	var absAutoLoadFileName string
	var absPluginFileName string
	absDirName = vimFilePath + dirName
	absAutoLoadDirName = absDirName + "/autoload"
	absPluginDirName = absDirName + "/plugin"
	absDocDirName = absDirName + "/doc"
	absAutoLoadFileName = absAutoLoadDirName + "/" + dirName
	absPluginFileName = absPluginDirName + "/" + dirName
	//vim plugin directory
	createDir(absDirName)
	//autoload directory
	createDir(absAutoLoadDirName)
	//plugin directory
	createDir(absPluginDirName)
	//doc directory
	createDir(absDocDirName)
	//autoload file
	createFile(absAutoLoadFileName)
	//plugin file
	createFile(absPluginFileName)
}

func inputKey() string {
	var inputKeyStr string
	for {
		fmt.Print("[(e)xecute/(s)uspension/(r)ename]: ")
		fmt.Scan(&inputKeyStr)
		if inputKeyStr == "s" || inputKeyStr == "S" {
			fmt.Println("process close")
			os.Exit(0)
		} else if inputKeyStr == "r" || inputKeyStr == "R" || inputKeyStr == "e" || inputKeyStr == "E" {
			return inputKeyStr
		} else {
			fmt.Println("Please enter the correct key")
			continue
		}
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
	var dirName string
	var inputStr string

	for {
		fmt.Print("Please Enter the name of vim plugin: ")
		fmt.Scan(&dirName)
		inputStr = inputKey()
		if inputStr == "r" || inputStr == "R" {
			continue
		}
		break
	}

	dirName = dirName + ".vim"

	createPluginDir(dirName, vimFilePath)

}
