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

type ExpectedOutcomeEncoding struct {
	input  string
	output int
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

func findRequiredMove(player2Move, expectedOutcome string) string {
	requiredOutcome := decodeExpectedOutpcome(expectedOutcome)
	possibleMoves := []string{"A", "B", "C"}

	for _, v := range possibleMoves {
		round := Round{decodeMove(v), decodeMove(player2Move)}
		score := getRoundScore(round)
		if score == requiredOutcome {
			return v
		}
	}

	log.Fatal("Failed to find required move")
	return ""
}

func parseInput(rawInput string) []Round {
	var rounds []Round
	rawRounds := strings.Split(rawInput, "\n")
	for _, round := range rawRounds {
		moves := strings.Split(round, " ")
		player2Move := moves[0]
		player1Move := findRequiredMove(player2Move, moves[1])
		rounds = append(rounds, Round{player1Move, player2Move})
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
	round.player1Move = decodeMove(round.player1Move)
	round.player2Move = decodeMove(round.player2Move)
	return getRoundScore(round) + getMoveScore(round.player1Move)
}

func decodeExpectedOutpcome(expectedOutcome string) int {
	decoder := []ExpectedOutcomeEncoding{
		{
			input:  "X",
			output: 0,
		},
		{
			input:  "Y",
			output: 3,
		},
		{
			input:  "Z",
			output: 6,
		},
	}

	for _, v := range decoder {
		if v.input == expectedOutcome {
			return v.output
		}
	}

	return 0
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
	}

	for _, v := range decoder {
		if v.input == encodedMove {
			return v.output
		}
	}

	return encodedMove
}

func getRoundScore(round Round) int {
	// Draw
	if round.player1Move == round.player2Move {
		return 3
	}

	// Loss
	if round.player1Move == "Rock" && round.player2Move == "Paper" || round.player1Move == "Paper" && round.player2Move == "Scissors" || round.player1Move == "Scissors" && round.player2Move == "Rock" {
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
