package main

import (
	"fmt"
	"strings"
)

/*
defines the Board "class", utils for operations dealing with
just a single, simple, TicTacToe board.
*/
type Board struct {
	arr    [3][3]rune
	winner rune
}

// toString method for Board struct
func (b Board) String() string {
	var str string
	for i := 0; i < len(b.arr); i++ {
		str += fmt.Sprint(b.arr[i], '\n')
	}
	str = str[:len(str)-1]

	return str
}

func (b Board) getWinner() rune {
	return b.winner
}

func (b Board) computeWinner() rune {
	fmt.Println("The computeWinner method on a Board hasn't been implemented yet!")
	return b.getWinner()
}

func (b Board) set(i, j int, marker rune) {
	b.arr[i][j] = marker
}

/*
defines the SuperBoard "class", utils for operations dealing with
the entire 3x3 grid of smaller TicTacToe boards.
*/
type SuperBoard struct {
	arr [3][3]Board
}

// toString method for SuperBoard struct
func (b SuperBoard) String() string {
	var sb strings.Builder

	for sri := 0; sri < 3; sri++ {
		for ri := 0; ri < 3; ri++ {
			for sci := 0; sci < 3; sci++ {
				s := fmt.Sprint(b.arr[sri][sci].arr[ri])
				s = strings.Trim(s, "[]")
				sb.WriteString(fmt.Sprint(s, "  "))
			}
			sb.WriteRune('\n')
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
