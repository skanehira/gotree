package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type depth struct {
	hasNext bool
}

func search(dir string, depths []depth) {
	// Exclusion dotfiles
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for i, info := range infos {
		if strings.HasPrefix(info.Name(), ".") {
			continue
		}

		for k, depth := range depths {
			if k == 0 {
				fmt.Print("|   ")
				continue
			}
			if depth.hasNext {
				fmt.Print("|   ")
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
			var hasNext bool
			if i == len(infos)-1 {
				hasNext = false
			} else {
				hasNext = true
			}
			depths = append(depths, depth{hasNext})
			search(filepath.Join(dir, info.Name()), depths)
			depths = depths[:len(depths)-1]
		}
	}
}
func main() {
	dir := os.Args[1]
	fmt.Println(dir)
	search(dir, make([]depth, 0))
}
