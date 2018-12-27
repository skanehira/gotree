package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := os.Args[1]
	fmt.Println(dir)
	search(dir, 0)
}

func search(dir string, depth int) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for i, info := range infos {
		if strings.HasPrefix(info.Name(), ".") {
			continue
		}

		for j := 0; j < depth; j++ {
			fmt.Print("|   ")
		}

		if i == len(infos)-1 {
			fmt.Println("└──", info.Name())
		} else {
			fmt.Println("├──", info.Name())
		}

		if info.IsDir() {
			depth++
			search(filepath.Join(dir, info.Name()), depth)
			depth--
		}
	}
}
