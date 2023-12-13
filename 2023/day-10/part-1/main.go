package main

import (
	"log"
	"os"
	"strings"
)

var validStrs = []string{
	"|7F",
	"|JL",
	"-FL",
	"-J7",
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	res := process(input)
	log.Println(res)
}

func process(input string) int {
	// Build the grid
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([]string, 0)
	width := 0
	for _, line := range lines {
		width = len(line)
		grid = append(grid, strings.Split(line, "")...)
	}

	// Find the start position
	startIdx := 0
	for i, char := range grid {
		if char == "S" {
			startIdx = i
			break
		}
	}

	scoreGrid := make([]int, len(grid))
	score := check(startIdx, grid, scoreGrid, width, 1)

	return score / 2
}

func check(idx int, grid []string, scoreGrid []int, width, score int) int {
	scores := make([]int, 0)
	scoreGrid[idx] = score

	indexes := make([]int, 0)

	upIdx := idx - width
	if upIdx < 0 {
		upIdx = -1
	}
	indexes = append(indexes, upIdx)

	downIdx := idx + width
	if downIdx >= len(grid) {
		downIdx = -1
	}
	indexes = append(indexes, downIdx)

	leftIdx := idx - 1
	if (idx+1)%width == 1 {
		leftIdx = -1
	}
	indexes = append(indexes, leftIdx)

	rightIdx := idx + 1
	if (idx+1)%width == 0 {
		rightIdx = -1
	}
	indexes = append(indexes, rightIdx)

	for i, newIdx := range indexes {
		if newIdx == -1 {
			continue
		}

		nextPosScore := scoreGrid[newIdx]
		gridValue := grid[newIdx]
		validStr := validStrs[i]
		canMove := strings.Index(validStr, gridValue) != -1 && nextPosScore == 0
		if canMove {
			scores = append(scores, check(newIdx, grid, scoreGrid, width, score+1))
		}
	}

	maxScore := score
	for _, childScore := range scores {
		if maxScore < childScore {
			maxScore = childScore
		}
	}

	return maxScore
}
