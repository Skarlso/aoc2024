package main

import "testing"

func TestMain(t *testing.T) {
	grid := [][]int{
		{0, 1, 2, 3},
		{1, 2, 3, 4},
		{8, 7, 6, 5},
		{9, 8, 7, 6},
	}

	sum := validTrailHeads(grid, point{x: 0, y: 0})

	if sum != 1 {
		t.Fatalf("should have equaled 2 but was: %d", sum)
	}
}
