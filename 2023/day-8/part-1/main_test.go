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
			input: `RL

      AAA = (BBB, CCC)
      BBB = (DDD, EEE)
      CCC = (ZZZ, GGG)
      DDD = (DDD, DDD)
      EEE = (EEE, EEE)
      GGG = (GGG, GGG)
      ZZZ = (ZZZ, ZZZ)`,
			expected: 2,
		},
		{
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
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
