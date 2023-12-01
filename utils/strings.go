package utils

import "strconv"

func SilentAtoi(input string) int {
	i, _ := strconv.Atoi(input)
	return i
}

func ReplaceCharsAtIndexes(mapping map[int]string, text string) string {
	output := []rune(text)
	for k, v := range mapping {
		output[k] = []rune(v)[0]
	}
	return string(output)
}
