package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func walkDir(dir string, hasNexts []bool, limit int) {
	// get entry
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("%s: %s", os.Args[0], err)
		os.Exit(1)
	}

	// limit depth level
	if len(hasNexts) >= limit {
		return
	}

	// walk dir
	for i, info := range infos {
		// Exclusion dotfiles
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
				fmt.Printf("%s: %s", os.Args[0], err)
				os.Exit(1)
			}

			name = fmt.Sprintf("%s %s %s", info.Name(), "->", realPath)
		} else {
			name = info.Name()
		}

		// print tree
		if i == len(infos)-1 {
			fmt.Println("└──", name)
		} else {
			fmt.Println("├──", name)
		}

		// if entry is dir, recursive search
		if info.IsDir() {
			if i == len(infos)-1 {
				hasNexts = append(hasNexts, false)
			} else {
				hasNexts = append(hasNexts, true)
			}

			walkDir(filepath.Join(dir, info.Name()), hasNexts, limit)
			hasNexts = hasNexts[:len(hasNexts)-1]
		}
	}
}

func parseArgs() (string, int) {
	// get depth level
	limit := flag.Int("L", 99, "depth level")
	flag.Parse()

	if *limit < 1 {
		fmt.Printf("%s: Invalid level, must be greater than 0\n", os.Args[0])
		os.Exit(1)
	}

	// get specified dir
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}

	// print current dir
	var current string
	separator := string(os.PathSeparator)

	if strings.HasSuffix(dir, separator) {
		current = dir
	} else {
		current = dir + separator
	}

	return current, *limit
}

func main() {
	// parse args
	dir, limit := parseArgs()

	fmt.Println(dir)

	// walk dir
	walkDir(dir, []bool{}, limit)
}
