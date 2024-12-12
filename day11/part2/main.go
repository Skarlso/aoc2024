package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// we keep track of individual stone blinks and their value
type stone struct {
	value  string
	blinks int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	numbers := strings.Split(string(content), " ")
	// fmt.Println("numbers: ", numbers)
	cache := map[stone]int{}

	count := 0
	for _, s := range numbers {
		count += processStone(s, 0, cache)
	}

	fmt.Println(count)
}

func removeLeadingZeros(s string) string {
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return "0"
	}

	return s
}

func processStone(value string, blink int, cache map[stone]int) int {
	count := 1

	// we reached the end of the number of blinks for this stone.
	// only one remains.
	if blink == 75 {
		cache[stone{value: value, blinks: blink}] = count
		return 1
	}

	// if we have seen this stone with this state just return it
	if v, ok := cache[stone{value: value, blinks: blink}]; ok {
		return v
	}

	// do the transformations for this stone.
	if value == "0" {
		count = processStone("1", blink+1, cache)
	} else if len(value)%2 == 0 {
		front := value[0 : len(value)/2]
		back := value[len(value)/2:]
		// fmt.Println("front, back", front, back)
		front = removeLeadingZeros(front)
		back = removeLeadingZeros(back)

		count = processStone(front, blink+1, cache) + processStone(back, blink+1, cache)
	} else {
		n, _ := strconv.Atoi(value)
		n *= 2024
		count = processStone(strconv.Itoa(n), blink+1, cache)
	}

	// cache the state and the count for this state.
	cache[stone{value: value, blinks: blink}] = count

	// here the number will be the stones transformation
	// basically, we build a tree from all the stones and all their transformations
	// we individually process them and their positions and counts
	// and add everything back up once done.
	// like the fibonacci calculations, the cache will return already
	// calculated routes.
	return count
}
