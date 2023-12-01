package days

import "fmt"

type Day struct {
	PuzzleOne func()
	PuzzleTwo func()
}

var (
	DayMap map[int]*Day = make(map[int]*Day)
)

func (d *Day) RunSolution(puzzle int) {
	switch puzzle {
	case 1:
		fmt.Printf("Part %v\n", 1)
		d.PuzzleOne()
	case 2:
		fmt.Printf("Part %v\n", 2)
		d.PuzzleTwo()
	case -1:
		fmt.Printf("Part %v\n", 1)
		d.PuzzleOne()
		fmt.Printf("Part %v\n", 2)
		d.PuzzleTwo()
	default:
		fmt.Println("Invalid part specified")
	}
}
