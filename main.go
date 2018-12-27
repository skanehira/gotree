package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func walkDir(dir string, hasNexts []bool) {
	// Exclusion dotfiles
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// walk dir
	for i, info := range infos {
		if strings.HasPrefix(info.Name(), ".") {
			continue
		}

		for _, hasNext := range hasNexts {
			if hasNext {
				fmt.Print("│   ")
			} else {
				fmt.Print("    ")
			}
		}

		// if file is symlink
		var name string
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			realPath, err := os.Readlink(filepath.Join(dir, info.Name()))
			if err != nil {
				panic(err)
			}

			name = fmt.Sprintf("%s %s %s", info.Name(), "->", realPath)
		} else {
			name = info.Name()
		}

		if i == len(infos)-1 {
			fmt.Println("└──", name)
		} else {
			fmt.Println("├──", name)
		}

		if info.IsDir() {
			if i == len(infos)-1 {
				hasNexts = append(hasNexts, false)
			} else {
				hasNexts = append(hasNexts, true)
			}

			walkDir(filepath.Join(dir, info.Name()), hasNexts)
			hasNexts = hasNexts[:len(hasNexts)-1]
		}
	}
}

func main() {
	dir := os.Args[1]

	// print current dir
	var current string
	separator := string(os.PathSeparator)
	if strings.HasSuffix(dir, separator) {
		current = dir
	} else {
		current = dir + separator
	}

	fmt.Println(current)

	// walk dir
	walkDir(dir, []bool{})
}
