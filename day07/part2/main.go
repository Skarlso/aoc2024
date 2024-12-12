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
		if n := generateOperatorCombinations(eq); n > 0 {
			total += n
		}
	}

	fmt.Println(total)
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
		if s[i] == "||" {
			l := strconv.Itoa(left)
			l += s[i+1]
			left, _ = strconv.Atoi(l)
		}
	}

	return left
}

func generateOperatorCombinations(e eq) int {
	operators := []string{"*", "+", "||"}
	current := []string{fmt.Sprint(e.operands[0])}

	for _, num := range e.operands[1:] {
		var newResults []string

		for _, op := range operators {
			for _, prev := range current {
				expr := prev + " " + op + " " + fmt.Sprint(num)
				if eval(expr) == e.result {
					return e.result
				}

				newResults = append(newResults, expr)
			}
		}

		current = newResults
	}

	return 0
}
