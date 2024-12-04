package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := Day4Part2(string(input))
	fmt.Println(result)
}

func Day4Part1(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	width := len(lines[0])

	chars := make([]string, 0)
	for _, line := range lines {
		chars = append(chars, strings.Split(line, "")...)
	}

	score := 0
	directions := []int{-width, width, -1, 1, -width - 1, -width + 1, width - 1, width + 1}
	for i, char := range chars {
		if char != "X" {
			continue
		}

		for _, direction := range directions {
			if spellsXmas(chars, width, i, direction, "X") {
				score++
			}
		}
	}

	return score
}

func spellsXmas(chars []string, width int, pos int, direction int, target string) bool {
	if pos < 0 || pos >= len(chars) {
		return false
	}

	current := chars[pos]
	if current != target {
		return false
	}

	nextTarget := getnextchar(current)
	if nextTarget == "" {
		return true
	}

	onLeftEdge := pos%width == 0
	if onLeftEdge && (direction == -1 || direction == -width-1 || direction == width-1) {
		return false
	}

	onRightEdge := (pos+1)%width == 0
	if onRightEdge && (direction == 1 || direction == -width+1 || direction == width+1) {
		return false
	}

	newIdx := pos + direction
	return spellsXmas(chars, width, newIdx, direction, nextTarget)
}

func getnextchar(current string) string {
	switch current {
	case "X":
		return "M"
	case "M":
		return "A"
	case "A":
		return "S"
	case "S":
		return ""
	default:
		return "NONE"
	}
}

func Day4Part2(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	width := len(lines[0])

	chars := make([]string, 0)
	for _, line := range lines {
		chars = append(chars, strings.Split(line, "")...)
	}

	score := 0
	diagonals := []int{-width - 1, width + 1, width - 1, -width + 1}
	for i := width; i < len(chars)-width; i++ {
		char := chars[i]
		if char != "A" || (i+1)%width == 0 || i%width == 0 {
			continue
		}

		tiles := make([]string, 4)
		for j, diagonal := range diagonals {
			tiles[j] = chars[i+diagonal]
		}

		if !((tiles[0] == "M" && tiles[1] == "S") || (tiles[0] == "S" && tiles[1] == "M")) ||
			!((tiles[2] == "M" && tiles[3] == "S") || (tiles[2] == "S" && tiles[3] == "M")) {
			continue
		}

		score++
	}

	return score
}
