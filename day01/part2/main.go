package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	left, right := []int{}, []int{}
	rightCount := map[int]int{}
	for _, l := range split {
		lines := strings.Split(l, "   ")
		n, _ := strconv.Atoi(lines[0])
		m, _ := strconv.Atoi(lines[1])
		left = append(left, n)
		right = append(right, m)
		rightCount[m]++
	}

	sum := 0
	for _, n := range left {
		sum += n * rightCount[n]
	}

	fmt.Println(sum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
