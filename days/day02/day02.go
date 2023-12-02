package day02

import (
	"fmt"
	"strings"

	"github.com/mozoarella/adventofcode2023/days"
	"github.com/mozoarella/adventofcode2023/utils"
)

func Init() {
	days.DayMap[2] = &days.Day{
		PuzzleOne: solvePuzzleOne,
		PuzzleTwo: solvePuzzleTwo,
	}
}

type game struct {
	legal  bool
	rounds map[int]map[string]int
}

var (
	rules map[string]int = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games map[int]game = make(map[int]game)
)

func solvePuzzleOne() {
	inputText := utils.ReadFileToSlice("days/inputs/day02.txt")
	rawGames := parseGames(inputText)
	var legalGames []int
	for k, v := range rawGames {
		rounds := parseRounds(v)
		games[k] = game{
			legal:  true,
			rounds: rounds,
		}
	}
	//this probably could have been integrated with another step but this is fine.
	for k, v := range games {
		if !checkGameLegal(v) {
			v.legal = false
		} else {
			legalGames = append(legalGames, k)
		}
	}
	fmt.Println(utils.AddNumbersInSlice(legalGames))

}

func solvePuzzleTwo() {
	inputText := utils.ReadFileToSlice("days/inputs/day02.txt")
	rawGames := parseGames(inputText)
	var gamePowers []int
	for k, v := range rawGames {
		rounds := parseRounds(v)
		games[k] = game{
			legal:  true,
			rounds: rounds,
		}
	}
	for _, v := range games {
		gamePowers = append(gamePowers, calculateMinimumCubePower(v))
	}
	fmt.Println(utils.AddNumbersInSlice(gamePowers))

}

// This function yeets the "Round 1: " prefix and sticks the resulting string into a [int]string map with the round number.
// I consider myself lucky that the rounds were all in order so I could just use the index+1
func parseGames(input []string) map[int]string {
	parsedGames := make(map[int]string)
	for k, v := range input {
		id := k + 1
		_, results, _ := strings.Cut(v, ": ")
		parsedGames[id] = results
	}
	return parsedGames
}

// Turns rounds into a map of ints mapped to maps of colour name and amount (also ints)
func parseRounds(games string) map[int]map[string]int {
	// split the game rounds string into seperate rounds 'red 2, blue 4, green 5' 'red 5, blue 2, green 5'
	rawRounds := strings.Split(games, ";")
	rounds := make(map[int]map[string]int)
	for k, round := range rawRounds {
		roundColours := make(map[string]int)
		id := k + 1
		// split the colour string from the round into seperate 'red 2' 'blue 4' entries
		colours := strings.Split(round, ",")
		for _, c := range colours {
			colour := parseColour(c)
			//Take the array with [2 red] and turn it into red:2 :)
			roundColours[colour[1]] = utils.SilentAtoi(colour[0])
		}
		rounds[id] = roundColours

	}
	return rounds
}

func parseColour(colourAmount string) []string {
	// for whatever reason every string after the first one ended up with a leading space
	// so we use TrimSpace to get rid over any stragglers
	trimmedColour := strings.TrimSpace(colourAmount)
	perColour := strings.Split(trimmedColour, " ")
	return perColour
}

func checkGameLegal(game game) bool {
	gameLegal := true
	for _, round := range game.rounds {
		roundLegal := true
		for colour, amount := range round {
			if amount > rules[colour] {
				roundLegal = false
			}
		}
		if !roundLegal {
			gameLegal = false
		}
	}
	if !gameLegal {
		return false
	} else {
		return true
	}
}

func calculateMinimumCubePower(game game) int {
	var minimumColours map[string]int = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, round := range game.rounds {
		for colour, amount := range round {
			if amount > minimumColours[colour] {
				minimumColours[colour] = amount
			}
		}
	}
	total := minimumColours["red"] * minimumColours["blue"] * minimumColours["green"]
	return total
}
