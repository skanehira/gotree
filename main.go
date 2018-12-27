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

		if i == len(infos)-1 {
			fmt.Println("└──", info.Name())
		} else {
			fmt.Println("├──", info.Name())
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
	fmt.Println(dir)

	// walk dir
	walkDir(dir, []bool{})
}
