package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var game SuperBoard = NewSuperBoard()
	var p1turn bool = true
	var allowed_square int = -1
	var row, col int

	for game.winner == 0 {
		if p1turn {
			// let player make a move
			fmt.Println(game)
			row, col = DoUserMove(game, allowed_square)
		} else {
			// Compute possible move for computer
			fmt.Println("The computer is choosing a location...")
			minRow, minCol := allowed_square/3*3, allowed_square%3*3
			row = rand.Intn(3) + minRow
			col = rand.Intn(3) + minCol

			// Retry making the move as needed
			for !game.Set(row, col, 'O') {
				row = rand.Intn(3) + minRow
				col = rand.Intn(3) + minCol
			}
			fmt.Printf("Computer chose row=%d col=%d.\n", row, col)
		}

		// compute where the next person may go
		next_allowed_square := RowColToNextBoardIdx(row, col)
		if game.GetBoard(next_allowed_square).winner != 0 {
			allowed_square = -1
		} else {
			allowed_square = next_allowed_square
		}

		// swap turns
		p1turn = !p1turn
	}
}

func RowColToNextBoardIdx(row, col int) int {
	return row%3*3 + col%3
}

func RowColToCurrentBoardIdx(row, col int) int {
	return row/3*3 + col/3
}

func DoUserMove(game SuperBoard, allowed_square int) (int, int) {
	var loc string

	// Informational prompt conditional on next active square
	if allowed_square == -1 {
		fmt.Println("You may choose any section of the board for your next move.")
	} else {
		fmt.Printf("You must choose a location in square #%d for your next move. \n", allowed_square+1)
	}
	// get choice
	fmt.Print("Enter a location (a1-i9): ")
	fmt.Scan(&loc)
	loc = strings.ToLower(loc)

	for {
		// error if not specified in valid format
		if len(loc) != 2 || loc[0] < 'a' || loc[0] > 'z' || loc[1] < '1' || loc[1] > '9' {
			fmt.Print("Invalid location given. Please enter a location (a1-i9): ")
			fmt.Scan(&loc)
			continue
		}

		// convert to coords
		var row int = int(loc[0] - 'a')
		var col int = int(loc[1] - '1')

		// error if not in the right square
		if allowed_square != -1 {
			square := RowColToCurrentBoardIdx(row, col)
			if square != allowed_square {
				fmt.Printf("That is in square #%d, but you must choose a location in square #%d.\n", square+1, allowed_square+1)
				fmt.Print("Please enter a location (a1-i9): ")
				fmt.Scan(&loc)
				continue
			}
		}

		// must have chosen an empty location
		success := game.Set(row, col, 'X')
		if !success {
			fmt.Print("Chosen location is already filled.\nPlease enter a location (a1-i9): ")
			fmt.Scan(&loc)
			continue
		}

		// return the move made
		return row, col
	}

}
