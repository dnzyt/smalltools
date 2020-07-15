package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stdin, "Invalid number of arguments")
		return
	}
	for i := 1; i < len(os.Args); i++ {
		filename := os.Args[i]
		numOfLines, err := countLine(os.Args[i])
		if err != nil {
			fmt.Fprintf(os.Stdin, "Error while reading file %s\n", filename)
			continue
		}
		fmt.Fprintf(os.Stdin, "There are %d lines in %s\n", numOfLines, filename)
	}

}

func countLine(filename string) (int, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return -1, err
	}
	reader := bufio.NewReader(file)
	numOfLines := 0
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return -1, err
		}
		// isPrefix will be set if the line is too long (> 4096 bytes)
		if !isPrefix {
			numOfLines++
		}
	}
	return numOfLines, nil
}
