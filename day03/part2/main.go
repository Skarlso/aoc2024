package main

import (
	"fmt"
	"os"
)

type lexer struct {
	input    []byte
	position int // current position in input (points to current char)
	skipMul  bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	l := &lexer{
		input: content,
	}

	sum := 0
	for l.Next() {
		if v, ok := l.readNextNumber(); ok {
			sum += v
		}
	}

	fmt.Println(sum)
}

func (l *lexer) readNextNumber() (int, bool) {
	// if we encounter m and the next two letters are ul, we have mul.
	// we can be strict about this
	// Next() is increasing the position so we don't have to care about that here.
	c := l.input[l.position]

	switch c {
	case 'm':
		if l.skipMul {
			break
		}
		if l.position+3 < len(l.input)-1 && string(l.input[l.position:l.position+3]) == "mul" {
			// scan until closing bracket.
			var parameters string

			if l.input[l.position+3] != '(' {
				return -1, false
			}

			parameters += string(l.input[l.position+3])

			index := l.position + 4 // after the (
			for {
				parameters += string(l.input[index])
				if index >= len(l.input) {
					break
				}
				if l.input[index] == ')' {
					break
				}

				index++
			}

			var (
				a, b int
			)

			// check how many things got parsed properly
			if _, err := fmt.Sscanf(parameters, "(%d,%d)", &a, &b); err != nil {
				return -1, false
			}

			return a * b, true
		}
	case 'd':
		// all dos and don'ts always appear with correct ().
		if l.position+1 < len(l.input)-1 && "d"+string(l.input[l.position+1]) == "do" {
			l.skipMul = false
		}

		if l.position+5 < len(l.input)-1 && "d"+string(l.input[l.position+1:l.position+5]) == "don't" {
			l.skipMul = true
		}
	}

	return -1, false
}

func (l *lexer) Next() bool {
	// reached the end
	if l.position >= len(l.input)-1 {
		return false
	}

	// look ahead and see what we got
	l.position++

	return true
}
