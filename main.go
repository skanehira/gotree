package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// color
const (
	colorCyan   = "\x1b[36m%s\x1b[0m"
	colorYellow = "\x1b[33m%s\x1b[0m"
)

// flag
var (
	limit          = flag.Int("L", 99, "depth level")
	isColorMode    = flag.Bool("C", false, "color mode")
	exculudeTarget = flag.String("EX", "", "exclude specific file or dir")
)

// count
var (
	dirCount  = 0
	fileCount = 0
)

func walkDir(dir string, hasNexts []bool, limit int) {
	// get entries
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("%s: %s", os.Args[0], err)
		os.Exit(1)
	}

	// limit depth level
	if len(hasNexts) >= limit {
		return
	}

	// walk dir
	for i, entry := range entries {
		// exclude file or dir
		if strings.HasPrefix(entry.Name(), ".") || entry.Name() == *exculudeTarget {
			continue
		}

		for _, hasNext := range hasNexts {
			if hasNext {
				fmt.Print("│   ")
			} else {
				fmt.Print("    ")
			}
		}

		// if file is symlink, print relative path
		var name string
		if entry.Mode()&os.ModeSymlink == os.ModeSymlink {
			realPath, err := os.Readlink(filepath.Join(dir, entry.Name()))
			if err != nil {
				fmt.Printf("%s: %s", os.Args[0], err)
				os.Exit(1)
			}

			name = fmt.Sprintf("%s %s %s", entry.Name(), "->", realPath)
		} else {
			name = entry.Name()
		}

		// if color mode, add color
		if *isColorMode {
			if entry.IsDir() {
				name = fmt.Sprintf(colorCyan, name)
			} else {
				name = fmt.Sprintf(colorYellow, name)
			}
		}

		lastIndex := len(entries) - 1

		// print tree
		if i == lastIndex {
			fmt.Println("└──", name)
		} else {
			fmt.Println("├──", name)
		}

		// if entry is dir, recursive search
		if entry.IsDir() {
			dirCount++
			if i == lastIndex {
				hasNexts = append(hasNexts, false)
			} else {
				hasNexts = append(hasNexts, true)
			}

			walkDir(filepath.Join(dir, entry.Name()), hasNexts, limit)
			hasNexts = hasNexts[:len(hasNexts)-1]
		} else {
			fileCount++
		}
	}
}

func parseArgs() (string, int) {
	// get depth level
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

	return dir, *limit
}

func main() {
	// parse args
	dir, limit := parseArgs()

	// print current dir
	fmt.Println(dir)

	// walk dir
	walkDir(dir, []bool{}, limit)

	fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
}
