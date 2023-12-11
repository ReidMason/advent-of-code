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
			input: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
			expected: 6,
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
