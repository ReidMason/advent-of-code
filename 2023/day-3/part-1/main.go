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

var idxs = []int{-1, 1}

func process(input string) int {
	lineLength := strings.Index(input, "\n")
	idxs = append(idxs, []int{-lineLength - 1, -lineLength, -lineLength + 1, lineLength - 1, lineLength, lineLength + 1}...)

	input = strings.ReplaceAll(input, "\n", "")

	total := 0
	currNum := ""
	validNum := false
	for i, c := range input {
		if c < '0' || c > '9' {
			if validNum {
				numVal, _ := strconv.Atoi(currNum)
				total += numVal
				validNum = false
			}
			currNum = ""
			continue
		}

		currNum += string(c)
		if validNum {
			continue
		}

		for _, j := range idxs {
			validNum = validSymbol(input, i+j)
			if validNum {
				break
			}
		}
	}

	return total
}

func validSymbol(text string, idx int) bool {
	if idx < 0 || idx >= len(text) {
		return false
	}

	c := text[idx]
	return c != '.' && (c < '0' || c > '9')
}
