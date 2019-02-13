package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// parse is markdown parse to index
func parse(scanner *bufio.Scanner) []string {
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

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("parse [%s]\n\n", os.Args[1])
	for _, value := range parse(scanner) {
		fmt.Println(value)
	}
}
