package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := getInput()
	if err != nil {
		log.Fatal("Failed to get input data")
	}

	input := parseInput(rawInput)

	findTopThreeCalorieCounts(input)
}

func findTopThreeCalorieCounts(input []string) {
	var highestCalorieCounts []int

	totalCalories := 0
	for i, v := range input {
		if v != "" {
			calories, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Failed to parse calories input", err)
			}
			totalCalories += calories
		}

		if v == "" || i == len(input)-1 {
			if len(highestCalorieCounts) < 3 {
				highestCalorieCounts = append(highestCalorieCounts, totalCalories)
			} else if lowestValue, lowestIndex := minIntInSlice(highestCalorieCounts); lowestValue < totalCalories {
				highestCalorieCounts[lowestIndex] = totalCalories
			}
			totalCalories = 0
		}
	}

	totalOfTopThree := sumIntSlice(highestCalorieCounts)
	log.Printf("The top three Elves carrying the most calories have %d in total", totalOfTopThree)
}

func parseInput(input string) []string {
	result := strings.Split(input, "\n")
	for i, v := range result {
		result[i] = strings.TrimSpace(v)
	}

	return result
}

func minIntInSlice(slice []int) (int, int) {
	lowestValue := -1
	lowestIndex := -1
	for i, v := range slice {
		if v < lowestValue || lowestValue == -1 {
			lowestValue = v
			lowestIndex = i
		}
	}

	return lowestValue, lowestIndex
}

func sumIntSlice(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}

	return total
}

func getInput() (string, error) {
	data, err := os.ReadFile("input.txt")
	return string(data), err
}
