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

func process(input string) int {
	replacer := strings.NewReplacer("=", "", "(", "", ")", "", ",", "")
	input = replacer.Replace(input)

	text := strings.Fields(input)
	instructions := text[0]
	text = text[1:]

	return step(text, "AAA", instructions, 0, 0)
}

func step(input []string, target string, dir string, dirpos int, steps int) int {
	if target == "ZZZ" {
		return steps
	}

	if dirpos >= len(dir) {
		dirpos = 0
	}

	inc := 1
	if dir[dirpos] == 'R' {
		inc = 2
	}

	for i := 0; i < len(input); i += 3 {
		val := input[i]
		if val == target {
			newTarget := input[i+inc]
			return step(input, newTarget, dir, dirpos+1, steps+1)
		}
	}

	return steps
}
