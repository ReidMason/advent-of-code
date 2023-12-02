package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id, red, gree, blue int
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
	total := 0
	lines := strings.Split(strings.TrimRight(input, "\n\r"), "\n")

	for _, line := range lines {
		id, possible := parseLine(line)
		if possible {
			total += id
		}
	}

	return total
}

func parseLine(line string) (int, bool) {
	line += ";"
	pos := 5

	id := ""
	for pos < len(line) {
		c := line[pos]
		if c == ':' {
			break
		}
		id += string(c)
		pos++
	}
	numId, _ := strconv.Atoi(id)

	pos += 2
	currVal := ""
	currColour := ""
	gettingVal := true
	for pos < len(line) {
		c := line[pos]

		if c == ' ' {
			gettingVal = !gettingVal
			pos++
			continue
		} else if c == ',' || c == ';' {
			numCurrVal, _ := strconv.Atoi(currVal)
			if !isValid(currColour, numCurrVal) {
				return numId, false
			}
			currColour = ""
			currVal = ""
			gettingVal = true

			pos += 2
			continue
		}

		if gettingVal {
			currVal += string(c)
		} else {
			currColour += string(c)
		}

		pos++
	}

	return numId, true
}

func isValid(colour string, count int) bool {
	switch colour {
	case "red":
		return count <= 12
	case "green":
		return count <= 13
	case "blue":
		return count <= 14
	}

	return false
}
