package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ordering struct {
	x, y int
}

type update struct {
	origin []int
	// store the index so that it's easy to compare.
	numbers map[int]int // bool is easier to code with in the checks
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	readPageNumbers := false
	rules := make([]ordering, 0)
	updates := make([]update, 0)
	for _, l := range split {
		if l == "" {
			readPageNumbers = true

			continue
		}

		if !readPageNumbers {
			orderSplit := strings.Split(l, "|")
			x, _ := strconv.Atoi(orderSplit[0])
			y, _ := strconv.Atoi(orderSplit[1])
			o := ordering{
				x: x,
				y: y,
			}
			rules = append(rules, o)

			continue
		}

		pagesSplit := strings.Split(l, ",")
		original := make([]int, 0)
		m := map[int]int{}
		for i, p := range pagesSplit {
			n, _ := strconv.Atoi(p)
			original = append(original, n)
			m[n] = i
		}

		updates = append(updates, update{
			origin:  original,
			numbers: m,
		})
	}

	// good := make([][]int, 0)
	// this will obviously not work but let's try.
	// sum := 0
	// TODO: This most likely could be done in one loop... Meh.
	bad := make([]update, 0)
	for _, u := range updates {
		goodUpdate := true
		for _, r := range rules {
			// if both numbers are present, the rule applies
			// ALL the rules must apply that include numbers for this update
			index1, ok1 := u.numbers[r.x]
			index2, ok2 := u.numbers[r.y]
			if ok1 && ok2 {
				if index1 > index2 {
					goodUpdate = false
					break
				}
			}
		}

		if !goodUpdate {
			bad = append(bad, u)
		}
	}

	// Make this a QUEUE and start popping off bad ones from it.
	// If it's bad, do the change, requeue it.

	var (
		current update
		sum     int
	)
outer:
	for len(bad) != 0 {
		current, bad = bad[0], bad[1:]

		for _, r := range rules {
			index1, ok1 := current.numbers[r.x]
			index2, ok2 := current.numbers[r.y]
			if ok1 && ok2 && index1 > index2 {
				current.origin[index1], current.origin[index2] = current.origin[index2], current.origin[index1]
				current.numbers[r.x] = index2
				current.numbers[r.y] = index1
				bad = append(bad, current)
				continue outer
			}
		}

		sum += current.origin[len(current.origin)/2]

	}

	fmt.Println(sum)

}
