package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Container [][]string
type Movement struct {
	amount, from, to int
}

func main() {
	input := getInput()
	container := getContainerState(input)
	movements := getMoveInstructions(input)

	for _, movement := range movements {
		container = moveCrates(container, movement)
	}

	log.Print(strings.Join(getTopOfStacks(container), ""))
}

func getTopOfStacks(container Container) []string {
	var topOfStacks []string
	for _, column := range container {
		topOfStacks = append(topOfStacks, column[0])
	}
	return topOfStacks
}

func getContainerState(input string) Container {
	var container Container
	crateLines := getCrateLines(input)
	var letters []string
	for i := 0; i < len(crateLines)-1; i++ {
		line := crateLines[i]
		columns := strings.Split(line, "")
		counter := 0
		for j := 1; j < len(columns); j += 4 {
			// Create new column
			if counter >= len(container) {
				container = append(container, []string(nil))
			}
			char := columns[j]
			if char != " " {
				container[counter] = append(container[counter], char)
			}
			counter++
		}
		letters = append(letters, columns[0])
	}

	return container
}

func getMoveInstructions(input string) []Movement {
	var movements []Movement
	addMovements := false
	for _, line := range strings.Split(input, "\n") {
		if !addMovements && line == "" {
			addMovements = true
			continue
		}

		if addMovements {
			movements = append(movements, parseMovement(line))
		}
	}

	return movements
}

func parseMovement(line string) Movement {
	words := strings.Split(line, " ")
	return Movement{
		amount: parseInt(words[1]),
		from:   parseInt(words[3]),
		to:     parseInt(words[5]),
	}
}

func parseInt(text string) int {
	number, _ := strconv.Atoi(text)
	return number
}

func moveCrates(container Container, movement Movement) Container {
	fromIndex := movement.from - 1
	toIndex := movement.to - 1

	columnToTakeFrom := container[fromIndex]
	// Create a copy of this slice
	cratesToMove := append([]string(nil), columnToTakeFrom[:0+movement.amount]...)
	columnToTakeFrom = columnToTakeFrom[0+movement.amount:]

	columnToPutInto := container[toIndex]
	columnToPutInto = pushToFront(columnToPutInto, cratesToMove)

	container[fromIndex] = columnToTakeFrom
	container[toIndex] = columnToPutInto

	return container
}

func pushToFront(slice, elements []string) []string {
	return append(reverseSlice(elements), slice...)
}

func reverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

func getCrateLines(input string) []string {
	var crateLines []string
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		crateLines = append(crateLines, line)
	}

	return crateLines
}

func getInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Failed to load input")
	}

	return string(data)
}
