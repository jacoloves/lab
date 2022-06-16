package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
)

const (
	DirBrunch   string = "├──"
	LastBrunch  string = "└──"
	BrunchSpace string = "│   "
	DepthSpace  string = "    "
	ResetColor  string = "\x1b[0m"
)

type Counter struct {
	dirs  int
	files int
}

func (counter *Counter) index(path string) {
	stat, _ := os.Stat(path)
	if stat.IsDir() {
		counter.dirs += 1
	} else {
		counter.files += 1
	}
}

func (counter *Counter) output() string {
	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
}

func dirnamesFrom(base string) []string {
	file, err := os.Open(base)
	if err != nil {
		fmt.Println(err)
	}

	names, _ := file.Readdirnames(0)
	file.Close()

	sort.Strings(names)
	return names
}

func color(path string) string {
	stat, _ := os.Stat(path)
	arr := [...]string{"Makefile", "go.mod", "go.sum"}
	if stat.IsDir() {
		return "\x1b[38;2;176;196;222m"
	} else {
		extensionName := filepath.Ext(path)
		switch extensionName {
		case ".go":
			return "\x1b[38;2;0;250;154m"
		case ".c":
			return "\x1b[38;2;255;215;0m"
		case ".md", ".txt":
			return "\x1b[38;2;255;140;0m\x1b[4m"
		default:
			_, file := filepath.Split(path)
			for _, v := range arr {
				if file == v {
					return "\x1b[38;2;255;140;0m\x1b[4m"
				}
			}
			return ""
		}
	}
}

func tree(counter *Counter, base string, prefix string) {
	names := dirnamesFrom(base)

	for index, name := range names {
		if name[0] == '.' {
			continue
		}
		subpath := path.Join(base, name)
		counter.index(subpath)

		if index == len(names)-1 {
			colorStr := color(subpath)
			fmt.Printf("%s %s%s%s\n", (prefix + LastBrunch), colorStr, name, ResetColor)
			tree(counter, subpath, prefix+DepthSpace)
		} else {
			colorStr := color(subpath)
			fmt.Printf("%s %s%s%s\n", (prefix + DirBrunch), colorStr, name, ResetColor)
			tree(counter, subpath, prefix+BrunchSpace)
		}
	}
}

func main() {
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		directory = "."
	}

	counter := new(Counter)
	fmt.Println(directory)

	tree(counter, directory, "")
	fmt.Println(counter.output())
}
