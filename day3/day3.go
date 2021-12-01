package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	X int
	Y int
}

type Slope struct {
	Right int
	Down  int
}

func countTrees(rows []string, slope Slope) int {
	position := Position{
		X: 0,
		Y: 0,
	}
	count := 0
	height := len(rows)
	width := len(rows[0])
	for position.Y < height-1 {
		position.X += slope.Right
		position.Y += slope.Down
		if string(rows[position.Y][position.X%width]) == "#" {
			count++
		}
	}
	return count
}

func getRows() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows
}

func main() {
	rows := getRows()
	fmt.Println("Part One", countTrees(rows, Slope{Right: 3, Down: 1}))

	oneOne := countTrees(rows, Slope{
		Right: 1,
		Down:  1,
	})
	threeOne := countTrees(rows, Slope{
		Right: 3,
		Down:  1,
	})
	fiveOne := countTrees(rows, Slope{
		Right: 5,
		Down:  1,
	})
	sevenOne := countTrees(rows, Slope{
		Right: 7,
		Down:  1,
	})
	oneTwo := countTrees(rows, Slope{
		Right: 1,
		Down:  2,
	})

	fmt.Println("Part Two",
		oneOne*threeOne*fiveOne*sevenOne*oneTwo)

}
