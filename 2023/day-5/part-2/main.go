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

type Seed struct {
	start int
	span  int
}

func getSeeds(line string) []Seed {
	seeds := make([]Seed, 0)
	seedStrs := strings.Split(line, " ")[1:]
	for i := 0; i < len(seedStrs); i += 2 {
		seedNum, _ := strconv.Atoi(seedStrs[i])
		seedRange, _ := strconv.Atoi(seedStrs[i+1])
		seeds = append(seeds, Seed{seedNum, seedRange})
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

func minSlice(vals []Seed) int {
	lowest := vals[0].start
	for _, val := range vals {
		if val.start < lowest {
			lowest = val.start
		}
	}

	return lowest
}

func buildMapping(lines []string, pos int, seeds []Seed) ([]Seed, int) {
	processed := make([]Seed, 0)
	currLine := lines[pos]

	for currLine != "" {
		nums := strings.Split(currLine, " ")
		steps, _ := strconv.Atoi(nums[2])
		sourceStart, _ := strconv.Atoi(nums[1])
		destStart, _ := strconv.Atoi(nums[0])
		sourceEnd := sourceStart + steps - 1

		seedCopy := make([]Seed, len(seeds))
		copy(seedCopy, seeds)

		x := len(seedCopy)
		for x > 0 {
			newSeedCopy, seed := pop(seedCopy)
			seedCopy = newSeedCopy
			x = len(seedCopy)

			seedStart := seed.start
			seedEnd := seed.start + seed.span

			if seedStart >= sourceStart && seedEnd <= sourceEnd {
				diff := seed.start - sourceStart

				seeds = remove(seeds, seed)
				seed.start = destStart + diff
				processed = append(processed, seed)

			} else if seedStart >= sourceStart && seedStart <= sourceEnd {
				seeds = remove(seeds, seed)

				endDiff := seedEnd - sourceEnd - 1
				newSeed1 := Seed{sourceEnd + 1, endDiff}
				seeds = append(seeds, newSeed1)
				seedCopy = append(seedCopy, newSeed1)

				newSeed2 := Seed{seedStart, seed.span - endDiff}
				diff := seed.start - sourceStart
				newSeed2.start = destStart + diff
				processed = append(processed, newSeed2)

			} else if seedEnd > sourceStart && seedEnd <= sourceEnd {
				seeds = remove(seeds, seed)

				newDiff := seedEnd - sourceStart
				newSeed1 := Seed{sourceStart, newDiff}

				diff := newSeed1.start - sourceStart
				newSeed1.start = destStart + diff
				processed = append(processed, newSeed1)

				newDiff = sourceStart - seedStart
				newSeed2 := Seed{seedStart, newDiff}
				seeds = append(seeds, newSeed2)
				seedCopy = append(seedCopy, newSeed2)
			}
		}

		pos++
		currLine = lines[pos]
	}

	pos += 2

	for _, seed := range seeds {
		processed = append(processed, seed)
	}

	return processed, pos
}

func pop(seeds []Seed) ([]Seed, Seed) {
	i := 0
	res := seeds[i]
	seeds = append(seeds[:i], seeds[i+1:]...)

	return seeds, res
}

func remove(seeds []Seed, seed Seed) []Seed {
	newSeeds := make([]Seed, 0)
	for _, oldSeed := range seeds {
		theSame := oldSeed.start == seed.start && oldSeed.span == seed.span
		if !theSame {
			newSeeds = append(newSeeds, oldSeed)
		}
	}

	return newSeeds
}
