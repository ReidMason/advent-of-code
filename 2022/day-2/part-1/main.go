package main

import (
	"log"
	"os"
	"strings"
)

type Encoding struct {
	input  string
	output string
}

type MoveScore struct {
	move  string
	score int
}

type Round struct {
	player1Move, player2Move string
}

func main() {
	rawInput, err := getInput()
	if err != nil {
		log.Fatal("Failed to load inputs")
	}

	rounds := parseInput(rawInput)
	finalScore := playGame(rounds)
	log.Print(finalScore)
}

func parseInput(rawInput string) []Round {
	var rounds []Round
	rawRounds := strings.Split(rawInput, "\n")
	for _, round := range rawRounds {
		moves := strings.Split(round, " ")
		rounds = append(rounds, Round{moves[1], moves[0]})
	}

	return rounds
}

func playGame(rounds []Round) int {
	totalScore := 0
	for _, round := range rounds {
		totalScore += playRound(round)
	}

	return totalScore
}

func playRound(round Round) int {
	player1Move := decodeMove(round.player1Move)
	player2Move := decodeMove(round.player2Move)

	return getRoundScore(player1Move, player2Move) + getMoveScore(player1Move)
}

func decodeMove(encodedMove string) string {
	decoder := []Encoding{
		{
			input:  "A",
			output: "Rock",
		},
		{
			input:  "B",
			output: "Paper",
		},
		{
			input:  "C",
			output: "Scissors",
		},
		{
			input:  "X",
			output: "Rock",
		},
		{
			input:  "Y",
			output: "Paper",
		},
		{
			input:  "Z",
			output: "Scissors",
		},
	}

	for _, v := range decoder {
		if v.input == encodedMove {
			return v.output
		}
	}

	return encodedMove
}

func getRoundScore(player1Move, player2Move string) int {
	// Draw
	if player1Move == player2Move {
		return 3
	}

	// Loss
	if player1Move == "Rock" && player2Move == "Paper" || player1Move == "Paper" && player2Move == "Scissors" || player1Move == "Scissors" && player2Move == "Rock" {
		return 0
	}

	// Win
	return 6
}

func getMoveScore(move string) int {
	moveScores := []MoveScore{
		{
			move:  "Rock",
			score: 1,
		},
		{
			move:  "Paper",
			score: 2,
		},
		{
			move:  "Scissors",
			score: 3,
		},
	}

	for _, moveScore := range moveScores {
		if moveScore.move == move {
			return moveScore.score
		}
	}

	return 0
}

func getInput() (string, error) {
	data, err := os.ReadFile("input.txt")
	return string(data), err
}
