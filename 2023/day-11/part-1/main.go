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
	rawLines := strings.Split(strings.TrimSpace(input), "\n")

	// Duplicate rows
	lines := make([][]string, 0)
	for _, line := range rawLines {
		newLine := make([]string, 0)
		for _, char := range line {
			newLine = append(newLine, string(char))
		}

		if !strings.Contains(line, "#") {
			lines = append(lines, newLine)
		}
		lines = append(lines, newLine)
	}

	// Duplicate columns
	cols := make([]int, 0)
	for i := 0; i < len(lines[0]); i++ {
		column := ""

		for j := 0; j < len(lines); j++ {
			val := lines[j][i]
			column += val
		}

		if !strings.Contains(column, "#") {
			cols = append(cols, i+len(cols))
		}
	}

	for i, row := range lines {
		newRow := row
		for _, col := range cols {
			newRow = append(newRow[:col+1], newRow[col:]...)
			newRow[col] = "."
		}
		lines[i] = newRow
	}

	// Find galaxies
	galaxies := make([][]int, 0)
	for i, row := range lines {
		for j, column := range row {
			if column == "#" {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	// Find distances
	total := 0
	for i, galaxy := range galaxies {
		for j, galaxy2 := range galaxies {
			if i == j {
				continue
			}

			dx := abs(galaxy[0] - galaxy2[0])
			dy := abs(galaxy[1] - galaxy2[1])
			distance := (dx + dy)
			total += distance
		}
	}

	return total / 2
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
