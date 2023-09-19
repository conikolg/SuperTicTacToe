package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var game SuperBoard = NewSuperBoard()
	var p1turn bool = true

	for game.winner == 0 {
		if p1turn {
			fmt.Println(game)
			fmt.Print("Enter a location (a1-i9): ")
			var loc string
			fmt.Scan(&loc)

			// for now, assume loc is a valid location
			row := loc[0] - 'a'
			col := loc[1] - '1'
			game.Set(int(row), int(col), 'X')
		} else {
			fmt.Println("The computer is choosing a location...")
			row := rand.Intn(9)
			col := rand.Intn(9)
			fmt.Printf("Computer chose row=%d col=%d \n", row, col)
			game.Set(int(row), int(col), 'O')
		}

		p1turn = !p1turn
	}
}
