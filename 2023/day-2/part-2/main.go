package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id, red, green, blue int
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
		total += parseLine(line)
	}

	return total
}

func parseLine(line string) int {
	line += ";"
	pos := strings.Index(line, ":")

	game := Game{
		red:   0,
		green: 0,
		blue:  0,
	}

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
			updateBoard(currColour, numCurrVal, &game)
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

	return game.red * game.blue * game.green
}

func updateBoard(colour string, count int, game *Game) {
	switch colour {
	case "red":
		game.red = max(game.red, count)
	case "green":
		game.green = max(game.green, count)
	case "blue":
		game.blue = max(game.blue, count)
	}
}
