package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func outputToc(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("parse [%s]\n\n", os.Args[1])
	for _, value := range parseMarkdown(scanner) {
		fmt.Println(value)
	}
}

// parse is markdown parse to index
func parseMarkdown(scanner *bufio.Scanner) []string {
	toc := make([]string, 0, 128)
	inCodeSyntax := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "```") {
			inCodeSyntax = !inCodeSyntax
			continue
		}
		if !inCodeSyntax && strings.HasPrefix(line, "#") {
			toc = append(toc, line)
		}
	}
	return toc
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not found arguments of toc file path.")
		os.Exit(1)
	}

	outputToc(os.Args[1])
}
