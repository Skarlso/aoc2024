package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	// fuck the lexer... let's regex that shit
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(string(content), -1)
	sum := 0
	for _, m := range matches {
		left, right := m[1], m[2]
		a, _ := strconv.Atoi(left)
		b, _ := strconv.Atoi(right)

		sum += a * b
	}

	fmt.Println(sum)
}
