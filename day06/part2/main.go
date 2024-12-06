package main

import (
	"bytes"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

type origin struct {
	p point
	d point
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := bytes.Split(content, []byte("\n"))
	floor := make([][]byte, 0)
	start := point{}
	for y, l := range split {
		for x, c := range l {
			if c == '^' {
				start.x = x
				start.y = y
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

	guard := start
	dir := directions[0]
	index := 0
	positions := map[point]struct{}{
		guard: {},
	}
	// collect the route, then go through those coordinates.
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

	loops := 0
	for position := range positions {
		if !inBound(position) {
			continue
		}
		if floor[position.y][position.x] == '#' || floor[position.y][position.x] == '^' {
			// skip if there is already an obstacle there
			continue
		}

		visited := map[origin]int{}
		// twice means the same location has been seen already coming from the same direction.
		twice := func(p, d point) bool {
			return visited[origin{p: p, d: d}] > 1
		}
		// set obstacle
		floor[position.y][position.x] = '#'
		current := start
		dir := directions[0]
		index := 0

		for {
			if twice(current, dir) {
				loops++
				break
			}

			next := addDirection(current, dir)
			if !inBound(next) {
				break
			}

			if floor[next.y][next.x] == '#' {
				index = (index + 1) % len(directions)
				dir = directions[index]

				continue
			}

			current = next
			visited[origin{p: current, d: dir}]++
		}

		floor[position.y][position.x] = '.' // set it back, try next item.
	}

	fmt.Println(loops)
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
