package main

import (
	"fmt"
	"os"
	"slices"
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
	numbers := strings.Split(string(content), " ")
	fmt.Println("numbers: ", numbers)

	for blinks := 0; blinks < 75; blinks++ {
		for i := 0; i < len(numbers); i++ {
			// fmt.Println("num: ", numbers[i])
			if numbers[i] == "0" {
				numbers[i] = "1"
			} else if len(numbers[i])%2 == 0 {
				front := numbers[i][0 : len(numbers[i])/2]
				back := numbers[i][len(numbers[i])/2:]
				// fmt.Println("front, back", front, back)
				front = removeLeadingZeros(front)
				back = removeLeadingZeros(back)

				numbers[i] = front
				numbers = slices.Insert(numbers, i+1, back)
				i++
			} else {
				n, _ := strconv.Atoi(numbers[i])
				n *= 2024
				numbers[i] = strconv.Itoa(n)
			}
		}

		fmt.Println(len(numbers)) // figure out what the len will be...
	}

}

func removeLeadingZeros(s string) string {
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return "0"
	}

	return s
}
