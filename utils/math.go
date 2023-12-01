package utils

func AddNumbersInSlice(numberslice []int) int {
	runningTotal := 0
	for _, i := range numberslice {
		runningTotal = runningTotal + i
	}
	return runningTotal
}
