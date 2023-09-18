package main

import (
	"fmt"
)

func main() {
	var game SuperBoard = NewSuperBoard()

	game.set(1, 1, 0, 0, 'O')
	game.set(1, 1, 1, 1, 'X')
	game.set(1, 1, 2, 2, 'O')
	game.set(2, 2, 0, 0, 'X')
	fmt.Println(game)
}
