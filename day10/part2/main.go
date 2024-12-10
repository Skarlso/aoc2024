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
	grid := [][]int{}
	for _, l := range split {
		row := []int{}
		for _, c := range l {
			row = append(row, int(c)-'0')
		}
		grid = append(grid, row)
	}

	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				sum += validTrailHeads(grid, point{x: x, y: y})
			}
		}
	}

	fmt.Println(sum)
}

func validTrailHeads(grid [][]int, start point) int {
	var trails int

	queue := []point{start}
	// visited := map[point]struct{}{}
	var current point
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if grid[current.y][current.x] == 9 {
			trails++
		}

		// for _, next := range  {
		// if _, ok := visited[next]; !ok {
		queue = append(queue, move(grid, current, start)...)
		// visited[next] = struct{}{}
		// }
		// }
	}

	return trails
}

var directions = []point{
	{x: 1, y: 0},
	{x: -1, y: 0},
	{x: 0, y: 1},
	{x: 0, y: -1},
}

func move(grid [][]int, current, start point) []point {
	var result []point

	inBound := func(p point) bool {
		return p.x >= 0 && p.y >= 0 && p.y < len(grid) && p.x < len(grid[p.y])
	}

	for _, d := range directions {
		next := point{x: current.x + d.x, y: current.y + d.y}
		if inBound(next) && grid[next.y][next.x] == grid[current.y][current.x]+1 && next != start {
			result = append(result, next)
		}
	}

	return result
}
