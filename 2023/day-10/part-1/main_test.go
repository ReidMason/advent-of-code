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
			input: `.....
.S-7.
.|.|.
.L-J.
.....`,
			expected: 4,
		},
		{
			input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`,
			expected: 8,
		},
		{
			input: `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`,
			expected: 8,
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
