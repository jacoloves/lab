package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	DirBrunch   string = "├──"
	LastBrunch  string = "└──"
	BrunchSpace string = "│   "
	DepthSpace  string = "    "
)

func stringDirFormat(nodeFlg, lastDirFlg bool) string {
	var format string
	if nodeFlg && lastDirFlg {
		format = LastBrunch
	} else if nodeFlg {
		format = DirBrunch
	} else if lastDirFlg {
		format = LastBrunch
	} else {
		format = DirBrunch
	}

	return format

}

func stringFormat(nodeFlg bool, lastFlg bool) string {
	var format string
	if nodeFlg && lastFlg {
		format = LastBrunch
	} else if nodeFlg {
		format = DirBrunch
	} else if lastFlg {
		format = LastBrunch
	} else {
		format = DirBrunch
	}

	return format
}

func tree(count int, nodeFlg bool, lastFlg bool, format string, depthLevel int, lastDirFlg bool) {
	p, _ := os.Getwd()

	infos, err := ioutil.ReadDir(p)
	if err != nil {
		fmt.Println(err)
	}

	var brunchName string
	brunchName += format
	for cnt, info := range infos {

		if info.Name() == "test3-2" {
			fmt.Println("---test start---")
			fmt.Println(format)
			fmt.Println(nodeFlg)
			fmt.Println(lastFlg)
			fmt.Println(depthLevel)
			fmt.Println("---test end---")
		}

		if info.IsDir() {
			if cnt == count {
				lastDirFlg = true
			} else {
				lastDirFlg = false
			}
			fileString := stringDirFormat(nodeFlg, lastDirFlg)
			fileName := format + fileString + " " + info.Name()
			fmt.Println(fileName)
			nodeFlg = true
			/*
				if cnt == count {
					lastFlg = true
				}
			*/
			os.Chdir(info.Name())
			if lastDirFlg {
				brunchName += DepthSpace
			} else {
				brunchName += BrunchSpace
			}
			execute(nodeFlg, lastFlg, brunchName, depthLevel+1, lastDirFlg)
			os.Chdir(p)
		} else {
			/*
				if info.Name() == "gtree_test.go" {
					fmt.Println("---test start---")
					fmt.Println(format)
					fmt.Println(nodeFlg)
					fmt.Println(lastFlg)
					fmt.Println(depthLevel)
					fmt.Println("---test end---")
				}
			*/
			format = ""
			if nodeFlg && lastDirFlg {
				if depthLevel > 1 {
					format += BrunchSpace
					//format += DepthSpace
					depthLevel--
				}
				for i := 0; i < depthLevel; i++ {
					format += DepthSpace
				}
			} else if nodeFlg {
				for i := 0; i < depthLevel; i++ {
					format += BrunchSpace
				}
			}
			/*
				if info.Name() == "gtree_test.go" {
					fmt.Println("---test start---")
					fmt.Println(format)
					fmt.Println(nodeFlg)
					fmt.Println(lastFlg)
					fmt.Println(depthLevel)
					fmt.Println("---test end---")
				}
			*/
			lastFlg = false
			if cnt == count {
				lastFlg = true
			}
			if depthLevel == 1 && nodeFlg && lastDirFlg {
				format = DepthSpace
			}
			fileString := stringFormat(nodeFlg, lastFlg)
			fileName := format + fileString + " " + info.Name()
			fmt.Println(fileName)
		}
	}

}

func counter() int {
	cnt := 0
	p, _ := os.Getwd()

	infos, err := ioutil.ReadDir(p)
	if err != nil {
		fmt.Println(err)
	}

	for a, _ := range infos {
		cnt = a
	}

	return cnt
}

func execute(nodeFlg bool, lastFlg bool, format string, depthLevel int, lastDirFlg bool) {
	count := counter()
	tree(count, nodeFlg, lastFlg, format, depthLevel, lastDirFlg)
}

func main() {
	var directory string
	prevDir, _ := filepath.Abs(".")
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		directory = "."
	}
	fmt.Println(directory)
	targetDir, _ := filepath.Abs(directory)
	os.Chdir(targetDir)
	execute(false, false, "", 0, false)
	defer os.Chdir(prevDir)
}
