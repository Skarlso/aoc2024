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
	type regionInfo struct {
		count     int // area
		perimeter int // number of boundary edges (fences)
	}

	// This needs to be a slice of maps or regions because multiple regions can exist for the same character.
	region := make(map[rune]regionInfo)
	regions := make([]map[rune]regionInfo, 0)

	visited := make(map[[2]int]bool)

	var dfs func(x, y int, ch rune)
	dfs = func(x, y int, ch rune) {
		// bounds check and already-visited guard
		if x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix) {
			return
		}
		if visited[[2]int{x, y}] {
			return
		}
		if matrix[y][x] != ch {
			return
		}

		info := region[ch]
		// area
		info.count++

		// perimeter: for each of 4 sides, if neighbor is OOB or different letter -> +1
		// left
		if x-1 < 0 || matrix[y][x-1] != ch {
			info.perimeter++
		}
		// right
		if x+1 >= len(matrix[0]) || matrix[y][x+1] != ch {
			info.perimeter++
		}
		// up
		if y-1 < 0 || matrix[y-1][x] != ch {
			info.perimeter++
		}
		// down
		if y+1 >= len(matrix) || matrix[y+1][x] != ch {
			info.perimeter++
		}

		region[ch] = info
		visited[[2]int{x, y}] = true

		dfs(x-1, y, ch)
		dfs(x+1, y, ch)
		dfs(x, y-1, ch)
		dfs(x, y+1, ch)
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			ch := matrix[y][x]
			if !visited[[2]int{x, y}] {
				dfs(x, y, ch)
				if len(region) > 0 {
					regions = append(regions, region)
				}
				region = make(map[rune]regionInfo)
			}
		}
	}

	var price int
	for _, region := range regions {
		for char, info := range region {
			priceRegion := info.count * info.perimeter
			if debug {
				fmt.Println("Character:", string(char), "Area:", info.count, "Perimeter:", info.perimeter, "Price:", priceRegion)
			}
			price += priceRegion
		}
	}
	return price
}
