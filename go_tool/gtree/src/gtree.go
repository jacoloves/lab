package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

const (
	DirBrunch   string = "├──"
	LastBrunch  string = "└──"
	BrunchSpace string = "│   "
	DepthSpace  string = "   "
)

type PathCondition struct {
	filename   string
	depth      int
	isDir      bool
	isLastFile bool
}

func getPathCondition(path string, dirFlag bool) PathCondition {
	var pathcon PathCondition
	pathcon.depth = strings.Count(path, "/")
	pathcon.isLastFile = false
	_, pathcon.filename = filepath.Split(path)
	if dirFlag {
		pathcon.isDir = true
	} else {
		pathcon.isDir = false
	}
	return pathcon
}

func dirWalker() ([]string, []PathCondition) {
	var paths []string
	var pcs []PathCondition
	cnt := 0 // WalkDir do count
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		notCurrentDirFlg := false
		dirFlg := false
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if d.IsDir() {
			if cnt == 0 {
				// through
			} else {
				paths = append(paths, filepath.Join(".", path))
				notCurrentDirFlg = true
				dirFlg = true
			}
		} else {
			paths = append(paths, filepath.Join(".", path))
			notCurrentDirFlg = true
		}

		if notCurrentDirFlg {
			pcs = append(pcs, getPathCondition(path, dirFlg))
		}
		cnt++
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", ".", err)
	}

	return paths, pcs
}

func displayTree(pcs []PathCondition) {
	fmt.Println(".")
	for _, pathcon := range pcs {
		if pathcon.depth == 0 {
			fmt.Printf("%s%s\n", DirBrunch, pathcon.filename)
		} else {
			var depthString string
			for i := 0; i < pathcon.depth; i++ {
				depthString += BrunchSpace
			}
			fmt.Printf("%s%s%s\n", depthString, DirBrunch, pathcon.filename)
		}
	}
}

func main() {
	_, pcs := dirWalker()
	displayTree(pcs)
}
