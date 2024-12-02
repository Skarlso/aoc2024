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
		if checkSafe(nums) {
			sum++
		}
	}

	fmt.Println(sum)
}

func checkSafe(nums []int) bool {
	if isSafe(nums) {
		return true
	}

	for i := range nums {
		newList := make([]int, 0)
		newList = append(newList, nums[:i]...)
		newList = append(newList, nums[i+1:]...)
		if isSafe(newList) {
			return true
		}
	}

	return false
}

func isSafe(nums []int) bool {
	cmp := func(a, b int) bool { return a < b }
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] == 0 {
			continue
		}
		if nums[i]-nums[i+1] > 0 {
			cmp = func(a, b int) bool { return a > b }
		}
		break
	}

	for i := 1; i < len(nums); i++ {
		diff := abs(nums[i] - nums[i-1])

		if diff < 1 || diff > 3 || !cmp(nums[i-1], nums[i]) {
			return false
		}
	}

	return true
}

func convert(line string) []int {
	var nums []int

	for _, v := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}

	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
