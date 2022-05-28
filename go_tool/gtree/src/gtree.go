package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

const (
	DirBrunch  string = "|--"
	LastBrunch string = "`--"
	DepthSpace string = "|   "
)

func dirWalker() []string {
	var paths []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if d.IsDir() {
			// through
		} else {
			paths = append(paths, filepath.Join(".", path))
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", ".", err)
	}

	return paths
}

func displayTree(paths []string) {
	fmt.Println(".")
	for _, name := range paths {
		fmt.Printf("%s%s\n", DirBrunch, name)
	}
}

func main() {
	currentDirFiles := dirWalker()
	displayTree(currentDirFiles)
}
