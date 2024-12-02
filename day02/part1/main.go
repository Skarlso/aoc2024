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

	sum := 0
	for _, l := range split {
		nums := convert(l)
		safe := false
		if nums[1] > nums[0] {
			safe = checkIncreasing(nums)
		} else {
			safe = checkDecreasing(nums)
		}

		if safe {
			sum++
		}
	}

	fmt.Println(sum)
}

func convert(line string) []int {
	var nums []int

	for _, v := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}

	return nums
}

func checkIncreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		// not increasing
		if nums[i] <= nums[i-1] {
			return false
		}

		// increasing by too much
		if nums[i]-nums[i-1] > 3 {
			return false
		}
	}

	// fmt.Println("safe increasing: ", nums)

	return true
}

func checkDecreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		// not decreasing
		if nums[i] >= nums[i-1] {
			return false
		}

		// decreasing by too much
		if nums[i-1]-nums[i] > 3 {
			return false
		}
	}

	// fmt.Println("safe decreasing: ", nums)

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
