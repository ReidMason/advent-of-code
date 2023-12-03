package main

import (
	"log"
	"os"
	"testing"
)

func loadInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func TestProcess(t *testing.T) {
	input :=
		`467..114..
.../......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
s..$.*....
.664.598..
`

	expected := 4361
	result := process(input)

	if result != expected {
		t.Errorf("Expected: %d found: %d", expected, result)
	}
}

func BenchmarkProcess(b *testing.B) {
	input := loadInput()

	b.ResetTimer()
	process(input)
}
