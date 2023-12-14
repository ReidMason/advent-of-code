package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	res := process(input)
	log.Println(res)
}

const multiplier = 1000000

func process(input string) int {
	rawLines := strings.Split(strings.TrimSpace(input), "\n")

	// Build 2d array
	log.Println("Building 2d array")
	lines := make([][]string, 0)
	for _, line := range rawLines {
		newLine := make([]string, 0)
		for _, char := range line {
			newLine = append(newLine, string(char))
		}

		lines = append(lines, newLine)
	}

	// Find galaxies
	log.Println("Finding galaxies")
	galaxies := make([][]int, 0)
	for i, row := range lines {
		for j, column := range row {
			if column == "#" {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	// Find empty rows
	log.Println("Finding empty rows")
	rowAdditions := make([]int, 0)
	for i, row := range lines {
		valid := true
		for _, char := range row {
			if char == "#" {
				valid = false
				break
			}
		}

		if valid {
			rowAdditions = append(rowAdditions, i)
		}
	}

	// Find empty columns
	log.Println("Finding empty columns")
	colAdditions := make([]int, 0)
	for i := 0; i < len(lines[0]); i++ {
		column := ""

		for j := 0; j < len(lines); j++ {
			val := lines[j][i]
			column += val
		}

		if !strings.Contains(column, "#") {
			colAdditions = append(colAdditions, i)
		}
	}

	for i, galaxy := range galaxies {
		x := galaxy[0]
		y := galaxy[1]

		for _, rowAddition := range rowAdditions {
			if galaxy[0] > rowAddition {
				x += multiplier - 1
			}
		}

		for _, colAddition := range colAdditions {
			if galaxy[1] > colAddition {
				y += multiplier - 1
			}
		}

		galaxies[i] = []int{x, y}
	}

	// Find distances
	log.Println("Finding distances")
	total := 0
	for i, galaxy := range galaxies {
		for j, galaxy2 := range galaxies {
			if i == j {
				continue
			}

			dx := abs(galaxy[0] - galaxy2[0])
			dy := abs(galaxy[1] - galaxy2[1])
			distance := dx + dy
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
