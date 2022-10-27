package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundedIn := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, foundedIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, foundedIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, foundedIn[line], line)
		}
	}
}

func fileAlreadyCounted(curFile string, countedFiles []string) bool {
	for _, f := range countedFiles {
		if curFile == f {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, foundedIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !fileAlreadyCounted(f.Name(), foundedIn[line]) {
			foundedIn[line] = append(foundedIn[line], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}