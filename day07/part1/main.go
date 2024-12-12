package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type eq struct {
	result   int
	operands []int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")

	eqs := make([]eq, 0)
	for _, l := range split {
		s := strings.Split(l, ": ")
		ops := strings.Split(s[1], " ")
		res, _ := strconv.Atoi(s[0])

		opsNums := make([]int, 0)
		for _, o := range ops {
			n, _ := strconv.Atoi(o)
			opsNums = append(opsNums, n)
		}
		eqs = append(eqs, eq{
			result:   res,
			operands: opsNums,
		})
	}

	total := 0
	for _, eq := range eqs {
		if n := solve(eq); n > 0 {
			total += n
		}
	}

	fmt.Println(total)
}

// I don't need all just once... ffs.
func solve(e eq) int {
	// generate all possible combinations, and then evaluate them.
	combos := generateOperatorCombinations(e.operands)
	var result int

	for _, c := range combos {
		if e.result == eval(c) {
			return e.result
		}
	}

	return result
}

func eval(exp string) int {
	s := strings.Fields(exp)
	left, _ := strconv.Atoi(s[0])

	for i := 1; i < len(s)-1; i += 2 {
		if s[i] == "*" {
			right, _ := strconv.Atoi(s[i+1])
			left *= right
		}
		if s[i] == "+" {
			right, _ := strconv.Atoi(s[i+1])
			left += right
		}
	}

	return left
}

// generateOperatorCombinations generates all possible combinations of operators
func generateOperatorCombinations(numbers []int) []string {
	if len(numbers) <= 1 {
		return []string{fmt.Sprint(numbers[0])}
	}

	operators := []string{"*", "+"}

	// First convert the first number to string
	current := []string{fmt.Sprint(numbers[0])}

	// For each remaining number, add each possible operator followed by the number
	for _, num := range numbers[1:] {
		var newResults []string

		// Try both operators with this number
		for _, op := range operators {
			for _, prev := range current {
				expr := prev + " " + op + " " + fmt.Sprint(num)
				newResults = append(newResults, expr)
			}
		}

		current = newResults
	}

	return current
}
