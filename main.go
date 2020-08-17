package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	Clear()
	fmt.Println("Welcome to Conways Game of Live.")
	fmt.Println("Please input the bord size:")

	var bordsize int = 10
	fmt.Scan(&bordsize)

	bord := make([][]bool, bordsize)
	for i := range bord {
		bord[i] = make([]bool, bordsize)
	}

	var running bool = true
	var step int

	PrintBord(bord, step)
	for running {
		fmt.Println("Enter keyword:")
		var input string
		fmt.Scan(&input)

		if input == "exit" {

			running = false

		} else if input == "help" {

			PrintBord(bord, step)
			fmt.Print("Keywords are: \nhelp      -> Shows this.\nstep      -> Updates Game by one step. (s also works)\nset       -> Set the value of a cell.\ntemplate  -> Place premade templates.\nexit      -> Stops programm.\n")

		} else if input == "step" || input == "s" {

			step++
			UpdateBord(&bord)
			PrintBord(bord, step)

		} else if input == "set" {

			PrintBord(bord, step)

			var x, y, arg0 int
			fmt.Println("Enter <x> <y> <value> (0 = dead, 1 = alive):")
			fmt.Scan(&x, &y, &arg0)

			var value bool = true
			if arg0 == 0 {
				value = false
			}
			bord[y][x] = value
			PrintBord(bord, step)

		} else if input == "template" {

			PrintBord(bord, step)
			fmt.Print("Tmplates are: \n1 -> Flyer \nEnter number:\n")

			var number int
			fmt.Scan(&number)
			switch number {
			case 1:
				if bordsize >= 3 {
					bord[0][1] = true
					bord[1][2] = true
					bord[2][0] = true
					bord[2][1] = true
					bord[2][2] = true
				}
				break
			}

			PrintBord(bord, step)

		} else {

			PrintBord(bord, step)
			fmt.Print("No valid input! \nValid keywords are: help, step (s), set, template, exit \n")

		}
	}
}

// PrintBord prints the bord to the consol using O for false and X for true
func PrintBord(bord [][]bool, step int) {
	Clear()

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

// Clear clears the Consol
func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
