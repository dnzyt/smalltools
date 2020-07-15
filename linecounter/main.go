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
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stdin, "Cannot open file: %s\n", filename)
		return
	}
	reader := bufio.NewReader(file)
	numOfLines := 0
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stdin, "Error while reading file: %v\n", err.Error())
			break
		}
		// isPrefix will be set if the line is too long (> 4096 bytes)
		if !isPrefix {
			numOfLines++
		}
	}
	fmt.Fprintf(os.Stdin, "There are %d lines in %s\n", numOfLines, filename)
}
