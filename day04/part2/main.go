package main

import (
	"bytes"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := bytes.Split(content, []byte("\n"))
	words := make([][]byte, 0)
	for _, l := range split {
		words = append(words, l)
	}

	found := 0
	for y := 0; y < len(words); y++ {
		for x := 0; x < len(words[y]); x++ {
			if words[y][x] == 'A' {
				found += searchFrom(point{x: x, y: y}, words)
			}
		}
	}

	fmt.Println(found)
}

func searchFrom(p point, words [][]byte) int {
	// fmt.Println("point: ", p)
	found := 0

	inBounds := func(p point) bool {
		return p.x >= 0 && p.y >= 0 && p.y < len(words) && p.x < len(words[p.y])
	}

	upperLeft := point{x: p.x - 1, y: p.y - 1}
	upperRight := point{x: p.x + 1, y: p.y - 1}
	lowerLeft := point{x: p.x - 1, y: p.y + 1}
	lowerRight := point{x: p.x + 1, y: p.y + 1}
	if !inBounds(upperLeft) || !inBounds(upperRight) || !inBounds(lowerLeft) || !inBounds(lowerRight) {
		return 0
	}

	// M M
	// S S
	if words[upperLeft.y][upperLeft.x] == 'M' && words[upperRight.y][upperRight.x] == 'M' &&
		words[lowerLeft.y][lowerLeft.x] == 'S' && words[lowerRight.y][lowerRight.x] == 'S' {
		found++
	}

	// S S
	// M M
	if words[upperLeft.y][upperLeft.x] == 'S' && words[upperRight.y][upperRight.x] == 'S' &&
		words[lowerLeft.y][lowerLeft.x] == 'M' && words[lowerRight.y][lowerRight.x] == 'M' {
		found++
	}

	// S M
	// S M

	if words[upperLeft.y][upperLeft.x] == 'S' && words[upperRight.y][upperRight.x] == 'M' &&
		words[lowerLeft.y][lowerLeft.x] == 'S' && words[lowerRight.y][lowerRight.x] == 'M' {
		found++
	}

	// M S
	// M S
	if words[upperLeft.y][upperLeft.x] == 'M' && words[upperRight.y][upperRight.x] == 'S' &&
		words[lowerLeft.y][lowerLeft.x] == 'M' && words[lowerRight.y][lowerRight.x] == 'S' {
		found++
	}

	return found
}
