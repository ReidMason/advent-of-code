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

	result := Day6Part1(strings.TrimSpace(string(input)))
	fmt.Println(result)
}

const (
	up = iota
	right
	down
	left
)

func Day6Part1(input string) int {
	guardRotation := up
	input = strings.Replace(input, "^", "G", 1)

	rows := strings.Split(input, "\n")
	grid := make([][]string, 0)
	for _, row := range rows {
		grid = append(grid, strings.Split(row, ""))
	}

	guardX, guardY := findGuard(grid)
	grid[guardY][guardX] = "."

	visitedCells := make(map[string]bool)

	for {
		fmt.Println("Guard is at:", guardX, guardY)
		visitedCells[fmt.Sprintf("%d,%d", guardX, guardY)] = true
		frontX, frontY := findPosInFront(grid, guardX, guardY, guardRotation)
		if frontX == -1 || frontY == -1 {
			fmt.Println("Guard is at the edge of the grid")
			break
		}

		inFrontContent := grid[frontY][frontX]
		if inFrontContent == "#" {
			guardRotation = rotateRotation(guardRotation)
			fmt.Println("Guard is rotating facing:", guardRotation)
			continue
		}

		guardX, guardY = frontX, frontY
	}

	return len(visitedCells)
}

func rotateRotation(rotation int) int {
	return (rotation + 1) % 4
}

func findPosInFront(grid [][]string, x, y int, direction int) (int, int) {
	switch direction {
	case up:
		if y-1 >= 0 {
			return x, y - 1
		}
	case down:
		if y+1 < len(grid) {
			return x, y + 1
		}
	case left:
		if x-1 >= 0 {
			return x - 1, y
		}
	case right:
		if x+1 < len(grid[y]) {
			return x + 1, y
		}
	}
	return -1, -1
}

func findGuard(grid [][]string) (int, int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == "G" {
				return x, y
			}
		}
	}
	return -1, -1
}
