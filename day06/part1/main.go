package main

import (
	"fmt"
	"os"
	"strings"
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

	split := strings.Split(string(content), "\n")
	floor := []string{}
	guard := point{}
	for y, l := range split {
		for x, c := range l {
			if c == '^' {
				guard.x = x
				guard.y = y
			}
		}
		floor = append(floor, l)
	}

	// fmt.Println("found guard at: ", start)
	directions := []point{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}

	// fmt.Println("direction: ", dir)
	inBound := func(p point) bool {
		return p.x >= 0 && p.y >= 0 && p.y < len(floor) && p.x < len(floor[p.y])
	}
	addDirection := func(g, d point) point {
		return point{x: g.x + d.x, y: g.y + d.y}
	}

	dir := directions[0]
	index := 0
	positions := map[point]struct{}{
		guard: {},
	}
	for inBound(guard) {
		// display(guard, floor)
		// fmt.Println("--------")
		// time.Sleep(500 * time.Millisecond)

		next := addDirection(guard, dir)
		if inBound(next) && floor[next.y][next.x] == '#' {
			index = (index + 1) % len(directions)
			dir = directions[index]

			continue
		}

		guard = next
		positions[guard] = struct{}{}
	}

	fmt.Println("guard is done walking: ", len(positions)-1)
}

func display(guard point, floor []string) {
	for y := range floor {
		for x, c := range floor[y] {
			if y == guard.y && x == guard.x {
				fmt.Print("X")
			} else {
				fmt.Print(string(c))
			}
		}

		fmt.Println()
	}
}
