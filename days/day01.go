package days

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mozoarella/adventofcode2023/utils"
)

var (
	numberMap map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func init() {
	DayMap[1] = &Day{
		PuzzleOne: solvePuzzleOne,
		PuzzleTwo: solvePuzzleTwo,
	}
}

func solvePuzzleOne() {
	inputText := utils.ReadFileToSlice("days/inputs/day01.txt")
	var allValues []int
	for _, v := range inputText {
		allValues = append(allValues, firstAndLastNumber(v))
	}
	fmt.Println(utils.AddNumbersInSlice(allValues))
}

func solvePuzzleTwo() {
	inputText := utils.ReadFileToSlice("days/inputs/day01.txt")
	var allValues []int
	for _, v := range inputText {
		v = utils.ReplaceCharsAtIndexes(findWrittenNumbersInString(v), v)
		firstAndLastNumber(v)
		allValues = append(allValues, firstAndLastNumber(v))
	}
	fmt.Println(utils.AddNumbersInSlice(allValues))
}

func firstAndLastNumber(input string) int {
	numberMatch := regexp.MustCompile(`\d`)
	numbersInString := numberMatch.FindAllString(input, -1)
	return utils.SilentAtoi(numbersInString[0] + numbersInString[len(numbersInString)-1])
}

func convertLettersToNumbers(input string) string {
	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	return replacer.Replace(input)
}

func findWrittenNumbersInString(input string) map[int]string {

	matches := make(map[int]string)
	for i := 0; i < len(input); i++ {
		start := i
		end := len(input)
		for k := range numberMap {
			index := strings.Index(input[start:end], k)
			if index == 0 {
				matches[i] = convertLettersToNumbers(k)
			}
		}
	}
	return matches
}
