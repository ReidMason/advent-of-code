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
			input: `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`,
			expected: 4,
		},
		{
			input: `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`,
			expected: 4,
		},
		{
			input: `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`,
			expected: 8,
		},
		{
			input: `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`,
			expected: 10,
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
