package main

import (
	"log"
	"os"
	"strconv"
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
	pos := 0
	lines := strings.Split(input, "\n")

	seeds := getSeeds(lines[pos])
	pos += 3

	for pos < len(lines) {
		seeds, pos = buildMapping(lines, pos, seeds)
	}

	return minSlice(seeds)
}

func getSeeds(line string) []int {
	seeds := make([]int, 0)
	for _, str := range strings.Split(line, " ")[1:] {
		seedNum, _ := strconv.Atoi(str)
		seeds = append(seeds, seedNum)
	}

	return seeds
}

func get(mapping map[int]int, source int) int {
	val, ok := mapping[source]
	if !ok {
		return source
	}

	return val
}

func minSlice(vals []int) int {
	lowest := vals[0]
	for _, val := range vals {
		if val < lowest {
			lowest = val
		}
	}

	return lowest
}

func buildMapping(lines []string, pos int, seeds []int) ([]int, int) {
	processed := make([]int, 0)
	currLine := lines[pos]
	for currLine != "" {
		nums := strings.Split(currLine, " ")
		steps, _ := strconv.Atoi(nums[2])
		sourceStart, _ := strconv.Atoi(nums[1])
		destStart, _ := strconv.Atoi(nums[0])
		sourceEnd := sourceStart + steps - 1

		for i, seed := range seeds {
			found := false
			for _, p := range processed {
				if p == i {
					found = true
					break
				}
			}
			if found {
				continue
			}

			if seed >= sourceStart && seed <= sourceEnd {
				diff := seed - sourceStart
				seeds[i] = destStart + diff
				processed = append(processed, i)
			}
		}

		pos++
		currLine = lines[pos]
	}

	pos += 2

	return seeds, pos
}
