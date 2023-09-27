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
	var difficulty string

	// Ask player what difficulty they want
	fmt.Print("Do you want to play on 'Easy' mode or 'Hard' mode? ")
	fmt.Scan(&difficulty)
	difficulty = strings.ToLower(difficulty)
	for difficulty != "easy" && difficulty != "hard" {
		fmt.Print("Invalid choice. Do you want to play on 'Easy' mode or 'Hard' mode? ")
		fmt.Scan(&difficulty)
		difficulty = strings.ToLower(difficulty)
	}

	// Start the game
	for game.GetWinner() == 0 {
		if p1turn {
			// let player make a move
			fmt.Println(game)
			row, col = DoUserMove(&game, allowed_square)
		} else {
			fmt.Println("The computer is choosing a location...")

			if difficulty == "easy" {
				row, col = doBasicComputerMove(&game, allowed_square)
			} else {
				row, col = doAdvComputerMove(&game, allowed_square)
			}

			// Clear screen
			fmt.Print("\033[H\033[2J")
			fmt.Printf("Computer chose %c%c.\n", row+'a', col+'1')
		}

		// compute where the next person may go
		game.ComputeWinner()
		next_allowed_square := RowColToNextBoardIdx(row, col)
		if game.GetBoard(next_allowed_square).winner != 0 {
			allowed_square = -1
		} else {
			allowed_square = next_allowed_square
		}

		// swap turns
		p1turn = !p1turn
	}

	fmt.Print("\033[H\033[2J")
	fmt.Println(game)
	fmt.Println()
	switch game.winner {
	case 'X':
		fmt.Println("Congratulations! You won!")
	case 'O':
		fmt.Println("Congratulations! You lost to a bot that picks randomly!!!")
	default:
		fmt.Println("It's a stalemate!")
	}
}

func RowColToNextBoardIdx(row, col int) int {
	return row%3*3 + col%3
}

func RowColToCurrentBoardIdx(row, col int) int {
	return row/3*3 + col/3
}

func DoUserMove(game *SuperBoard, allowed_square int) (int, int) {
	var loc string

	// Informational prompt conditional on next active square
	if allowed_square == -1 {
		fmt.Println("You may play in any section of the board for your next move.")
	} else {
		fmt.Printf("You must play in square #%d for your next move. \n", allowed_square+1)
	}

	// Compute user-friendly locations
	var minRow, minCol, maxRow, maxCol int
	if allowed_square != -1 {
		minRow, minCol = allowed_square/3*3, allowed_square%3*3
		maxRow, maxCol = minRow+2, minCol+2
	} else {
		minRow, minCol = 0, 0
		maxRow, maxCol = 8, 8
	}
	minRow += 'a'
	minCol += '1'
	maxRow += 'a'
	maxCol += '1'

	// get choice
	fmt.Printf("Enter a location (%c%c-%c%c): ", minRow, minCol, maxRow, maxCol)
	fmt.Scan(&loc)
	loc = strings.ToLower(loc)

	for {
		// error if not specified in valid format
		if len(loc) != 2 || loc[0] < 'a' || loc[0] > 'i' || loc[1] < '1' || loc[1] > '9' {
			fmt.Printf("Invalid location. Please enter a location (%c%c-%c%c): ", minRow, minCol, maxRow, maxCol)
			fmt.Scan(&loc)
			continue
		}

		// convert to coords
		var row int = int(loc[0] - 'a')
		var col int = int(loc[1] - '1')

		// error if not in the right square
		square := RowColToCurrentBoardIdx(row, col)
		if allowed_square != -1 {
			if square != allowed_square {
				fmt.Printf("That is in square #%d. Please choose a location in square #%d.\n", square+1, allowed_square+1)
				fmt.Printf("Please enter a location (%c%c-%c%c): ", minRow, minCol, maxRow, maxCol)
				fmt.Scan(&loc)
				continue
			}
		}

		// must have chosen a square that is still in play
		if game.GetBoard(square).winner != 0 {
			fmt.Printf("Square #%d has already been completed, you cannot make any more moves there.", square+1)
			fmt.Printf("Please enter a location (%c%c-%c%c): ", minRow, minCol, maxRow, maxCol)
			fmt.Scan(&loc)
			continue
		}

		// must have chosen an empty location
		success := game.Set(row, col, 'X')
		if !success {
			fmt.Println("Chosen location is already filled.")
			fmt.Printf("Please enter a location (%c%c-%c%c): ", minRow, minCol, maxRow, maxCol)
			fmt.Scan(&loc)
			continue
		}

		// return the move made
		return row, col
	}
}

func doBasicComputerMove(game *SuperBoard, allowed_square int) (int, int) {
	// Pick an available square, if needed
	for allowed_square == -1 {
		square := rand.Intn(9)
		if game.GetBoard(square).winner == 0 {
			allowed_square = square
		}
	}

	// Pick a space in the square
	minRow, minCol := allowed_square/3*3, allowed_square%3*3
	row := rand.Intn(3) + minRow
	col := rand.Intn(3) + minCol
	// Retry as needed
	for !game.Set(row, col, 'O') {
		row = rand.Intn(3) + minRow
		col = rand.Intn(3) + minCol
	}

	return row, col
}

func doAdvComputerMove(game *SuperBoard, allowed_square int) (int, int) {
	return doBasicComputerMove(game, allowed_square)
}
