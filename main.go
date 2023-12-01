package main

import (
	"flag"
	"fmt"

	"github.com/mozoarella/adventofcode2023/days"
)

var (
	dayFlag    *int = flag.Int("d", -1, "Pick the day for which you would like to run the code")
	puzzleFlag *int = flag.Int("p", -1, "Pick the part for which you would like to run the code")
)

func init() {
	flag.Parse()
}

func main() {

	if *dayFlag == -1 {
		for k, v := range days.DayMap {
			fmt.Printf("Day %v\n", k)
			v.RunSolution(*puzzleFlag)
		}
		return
	} else if val, ok := days.DayMap[*dayFlag]; ok {
		fmt.Printf("Day %v\n", *dayFlag)
		val.RunSolution(*puzzleFlag)
	} else {
		fmt.Println("Invalid day specified")
	}

}
