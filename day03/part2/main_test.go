package main

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	content := []byte(`[]sdfv.,,/.,349ufmul(4,5)do()dsadfsadon't()mul(3,4)slkjfm,.cm.,,mul(4,5),do()mul(4,3)`)
	l := &lexer{
		input: content,
	}

	sum := 0
	for l.Next() {
		if v, ok := l.readNextNumber(); ok {
			sum += v
		}
	}

	if sum != 32 {
		t.Fatal("sum was not 32 but: ", sum)
	}
}

func TestLexer2(t *testing.T) {
	content := []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)
	l := &lexer{
		input: content,
	}

	sum := 0
	for l.Next() {
		if v, ok := l.readNextNumber(); ok {
			fmt.Println(v)
			sum += v
		}
	}

	if sum != 48 {
		t.Fatal("sum was not 32 but: ", sum)
	}
}
