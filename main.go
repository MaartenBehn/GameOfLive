package main

import (
	"fmt"
)

func main() {

	fmt.Print("\033[H\033[2J") // Clears consol
	fmt.Println("Welcome to Conways Game of Live.")
	fmt.Println("Please input the bord size:")

	var bordsize int = 10
	//fmt.Scan(&bordsize)

	bord := make([][]bool, bordsize)
	for i := range bord {
		bord[i] = make([]bool, bordsize)
	}

	bord[3][3] = true
	bord[4][4] = true
	bord[5][2] = true
	bord[5][3] = true
	bord[5][4] = true

	var running bool = true
	var step int
	for running {
		PrintBord(bord, step)

		var validInput bool
		for !validInput {
			fmt.Println("Enter keyword:")
			var input string
			fmt.Scan(&input)

			switch input {
			case "exit":
				running = false
				validInput = true
				break
			case "step":
				step++
				UpdateBord(&bord)
				validInput = true
				break
			case "s":
				step++
				UpdateBord(&bord)
				validInput = true
				break
			default:
				PrintBord(bord, step)
				fmt.Print("No valid input! \nValid keywords are: step (s), exit \n")
				break
			}
		}

	}
}

// PrintBord prints the bord to the consol using O for false and X for true
func PrintBord(bord [][]bool, step int) {
	fmt.Print("\033[H\033[2J") // Clears consol

	fmt.Printf("%s%d%s", "Conways Game of Live \nCurrent step: ", step, "\n--------------------\n") // Prints header

	// Prints bord
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

	fmt.Println("--------------------") // lower spertion line
}

// UpdateBord updates all cells
func UpdateBord(bord *[][]bool) {

	// Creating a bord copy
	bordCopy := make([][]bool, len(*bord))
	for i := range bordCopy {
		bordCopy[i] = make([]bool, len((*bord)[i]))
		copy(bordCopy[i], (*bord)[i])
	}

	for y, collum := range bordCopy {
		for x, cell := range collum {
			activeCount := GetActivNeigborCount(bordCopy, y, x) // Checking how many Neigbors are active

			(*bord)[y][x] = (activeCount == 2 && cell) || activeCount == 3 // Getting the cell out of the pointer list without creating a copy an appling Conways laws.
		}
	}
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
