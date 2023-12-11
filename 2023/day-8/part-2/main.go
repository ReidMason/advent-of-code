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

	startingTargets := findStartingTargets(text)
	scores := make([]int, 0, len(startingTargets))
	for _, target := range startingTargets {
		score := step(text, target, instructions, 0, 0)
		scores = append(scores, score)
	}

	return lcmm(scores)
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}

func lcmm(nums []int) int {
	if len(nums) == 2 {
		return lcm(nums[0], nums[1])
	}
	var arg0 = nums[0]
	nums = nums[1:]
	return lcm(arg0, lcmm(nums))
}

func findStartingTargets(input []string) []string {
	startingTargets := make([]string, 0)

	for i := 0; i < len(input); i += 3 {
		val := input[i]
		if strings.HasSuffix(val, "A") {
			startingTargets = append(startingTargets, val)
		}
	}

	return startingTargets
}

func step(input []string, target string, dir string, dirpos int, steps int) int {
	if strings.HasSuffix(target, "Z") {
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
