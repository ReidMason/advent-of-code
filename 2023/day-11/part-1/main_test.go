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
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input: `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
			expected: 374,
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		tc := tc
		result := process(tc.input)

		if result != tc.expected {
			t.Errorf("Expected: %d found: %d", tc.expected, result)
		}
	}
}

func BenchmarkProcess(b *testing.B) {
	input := loadInput()

	b.ResetTimer()
	process(input)
}
