package main

import (
	"bytes"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

/*
\|/
- -
/|\
*/
var directions = []point{
	{x: 1, y: 1},
	{x: 0, y: 1},
	{x: 1, y: 0},
	{x: -1, y: 1},
	{x: -1, y: -1},
	{x: 1, y: -1},
	{x: -1, y: 0},
	{x: 0, y: -1},
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
			if words[y][x] == 'X' {
				found += searchFrom(point{x: x, y: y}, words)
			}
		}
	}

	fmt.Println(found)
}

func searchFrom(p point, words [][]byte) int {
	found := 0

	inBounds := func(p point) bool {
		return p.x >= 0 && p.y >= 0 && p.y < len(words) && p.x < len(words[p.y])
	}

	for _, d := range directions {
		// just add it if the right letter is found
		next := point{x: p.x + d.x, y: p.y + d.y}
		if !inBounds(next) {
			continue
		}

		if words[next.y][next.x] == 'M' {
			next.x += d.x
			next.y += d.y
			if inBounds(next) && words[next.y][next.x] == 'A' {
				next.x += d.x
				next.y += d.y
				if inBounds(next) && words[next.y][next.x] == 'S' {
					found++
				}
			}
		}
	}

	return found
}
