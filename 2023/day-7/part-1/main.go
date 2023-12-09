package main

import (
	"log"
	"os"
	"sort"
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

type Hand struct {
	hand       string
	bet, score int
}

var cards = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

func process(input string) int {
	hands := make([]Hand, 0)

	blocks := strings.Fields(input)
	for i := 1; i < len(blocks); i += 2 {
		bet, _ := strconv.Atoi(blocks[i])
		hands = append(hands, Hand{
			hand: blocks[i-1],
			bet:  bet,
		})
	}

	for i, hand := range hands {
		hands[i].score = getScore(hand.hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i]
		h2 := hands[j]

		if h1.score == h2.score {
			for i := range h1.hand {
				i1 := findIndex(cards, string(h1.hand[i]))
				i2 := findIndex(cards, string(h2.hand[i]))
				if i1 == i2 {
					continue
				}

				return i1 > i2
			}
		}

		return h1.score < h2.score
	})

	total := 0
	for i, hand := range hands {
		total += hand.bet * (i + 1)
	}

	return total
}

func getScore(hand string) int {
	maxCount := 0
	for _, card := range cards {
		count := strings.Count(hand, card)
		maxCount += count * count
	}

	return maxCount
}

func findIndex(arr []string, val string) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}

	return -1
}
