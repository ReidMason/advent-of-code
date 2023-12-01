package main

import (
	"log"
	"os"
	"strings"
)

type Rucksack struct {
	compartment1 []string
	compartment2 []string
}

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatal("Failed to load inputs")
	}

	totalItemPriority := 0
	rucksacks := parseInputs(input)
	for _, rucksack := range rucksacks {
		commonItem := findCommonItem(rucksack)
		totalItemPriority += getItemPriority(commonItem)
	}
	log.Print(totalItemPriority)
}

func parseInputs(input string) []Rucksack {
	var rucksacks []Rucksack
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line := strings.Split(line, "")
		midpoint := len(line) / 2
		rucksack := Rucksack{
			compartment1: line[0:midpoint],
			compartment2: line[midpoint:],
		}
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}

func findCommonItem(rucksack Rucksack) string {
	for _, item1 := range rucksack.compartment1 {
		for _, item2 := range rucksack.compartment2 {
			if item1 == item2 {
				return item2
			}
		}
	}

	return ""
}

func getItemPriority(item string) int {
	asciiOffset := 96
	b := int([]byte(item)[0]) - asciiOffset

	capitalAsciiOffset := 58
	if b < 0 {
		b += capitalAsciiOffset
	}

	return b
}

func getInput() (string, error) {
	data, err := os.ReadFile("input.txt")
	return string(data), err
}
