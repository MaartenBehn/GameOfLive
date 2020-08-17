package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Conways Game of Live.")
	fmt.Println("Please input the bord size:")

	var bordsize int
	fmt.Scan(&bordsize)

	bord := make([][]bool, bordsize)
	for i := range bord {
		bord[i] = make([]bool, bordsize)
	}

	PrintBord(bord)
}

// PrintBord prints the bord to the consol using O for false and X for true
func PrintBord(bord [][]bool) {
	for _, collum := range bord {
		for _, cell := range collum {
			if cell {
				fmt.Print("X")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Print("\n")
	}
}

// UpdateBord updates all cells
func UpdateBord(bord [][]bool) {
	/*for y, collum := range bord {
		for x, cell := range collum {
			activeCount := GetActivNeigborCount(bord, y, x)

		}

	}*/
}

// GetActivNeigborCount returns the ammount of active neigbors al cell has.
func GetActivNeigborCount(bord [][]bool, y int, x int) int {
	var activeCount int

	for dy := y - 1; dy <= y+1; dy++ {
		for dx := x - 1; dx <= x+1; dx++ {

			if dy == y && dx == x {
				continue
			}

			var ty, tx int = dy, dx

			if ty < 0 {
				ty = len(bord) - 1
			} else if ty >= len(bord) {
				ty = 0
			}

			if tx < 0 {
				tx = len(bord[0]) - 1
			} else if tx >= len(bord[0]) {
				tx = 0
			}

			if bord[ty][tx] {
				activeCount++
			}
		}
	}

	return activeCount
}
