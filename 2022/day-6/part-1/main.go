package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()

	log.Print(findMarkerIndex(input))
}

func findMarkerIndex(input string) int {
	markerLength := 4
	slice := strings.Split(input, "")
	for i := 0; i < len(slice)-markerLength-1; i++ {
		if !sliceHasMatchingEntries(slice[i : i+markerLength]) {
			return i + markerLength
		}
	}

	return 0
}

func sliceHasMatchingEntries(slice []string) bool {
	var processedEntries []string
	for _, entry := range slice {
		for _, v := range processedEntries {
			if entry == v {
				return true
			}
		}
		processedEntries = append(processedEntries, entry)
	}

	return false
}

func getInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Failed to load input")
	}

	return string(data)
}
