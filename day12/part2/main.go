package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	var matrix [][]rune
	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	price := computeTotalPrice(matrix, true)

	fmt.Println("Total price:", price)
}

func computeTotalPrice(matrix [][]rune, debug bool) int {
	visited := make(map[[2]int]bool)
	var totalPrice int

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if visited[[2]int{x, y}] {
				continue
			}

			// Find all cells in this region
			ch := matrix[y][x]
			cells := make(map[[2]int]bool)
			var dfs func(x, y int)
			dfs = func(x, y int) {
				if x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix) {
					return
				}
				if visited[[2]int{x, y}] || matrix[y][x] != ch {
					return
				}
				visited[[2]int{x, y}] = true
				cells[[2]int{x, y}] = true
				dfs(x-1, y)
				dfs(x+1, y)
				dfs(x, y-1)
				dfs(x, y+1)
			}
			dfs(x, y)

			area := len(cells)
			sides := countSides(cells)
			price := area * sides
			if debug {
				fmt.Println("Character:", string(ch), "Area:", area, "Sides:", sides, "Price:", price)
			}
			totalPrice += price
		}
	}
	return totalPrice
}

func countSides(cells map[[2]int]bool) int {
	sides := 0
	rowMap := make(map[int][]int)
	colMap := make(map[int][]int)
	for cell := range cells {
		x, y := cell[0], cell[1]
		rowMap[y] = append(rowMap[y], x)
		colMap[x] = append(colMap[x], y)
	}

	// horizontal sides (top and bottom)
	for y, xCoords := range rowMap {
		var topEdges, bottomEdges []int
		for _, x := range xCoords {
			if !cells[[2]int{x, y - 1}] {
				topEdges = append(topEdges, x)
			}
			if !cells[[2]int{x, y + 1}] {
				bottomEdges = append(bottomEdges, x)
			}
		}
		sides += countSegments(topEdges)
		sides += countSegments(bottomEdges)
	}

	// vertical sides (left and right)
	for x, yCoords := range colMap {
		var leftEdges, rightEdges []int
		for _, y := range yCoords {
			if !cells[[2]int{x - 1, y}] {
				leftEdges = append(leftEdges, y)
			}
			if !cells[[2]int{x + 1, y}] {
				rightEdges = append(rightEdges, y)
			}
		}
		sides += countSegments(leftEdges)
		sides += countSegments(rightEdges)
	}

	return sides
}

// countSegments counts continuous segments in a list of coordinates
func countSegments(coords []int) int {
	if len(coords) == 0 {
		return 0
	}

	coordSet := make(map[int]bool)
	for _, c := range coords {
		coordSet[c] = true
	}

	segments := 0
	visited := make(map[int]bool)

	for _, c := range coords {
		if visited[c] {
			continue
		}
		segments++
		visited[c] = true

		// both directions
		for i := c - 1; coordSet[i]; i-- {
			visited[i] = true
		}
		for i := c + 1; coordSet[i]; i++ {
			visited[i] = true
		}
	}

	return segments
}
