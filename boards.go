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
	arr    [][]rune
	winner rune
}

func NewBoard() Board {
	var a [][]rune = make([][]rune, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]rune, 3)
		for j := 0; j < 3; j++ {
			a[i][j] = ' '
		}
	}

	return Board{
		arr:    a,
		winner: 0,
	}
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

func (b Board) get(i, j int) rune {
	return b.arr[i][j]
}

/*
defines the SuperBoard "class", utils for operations dealing with
the entire 3x3 grid of smaller TicTacToe boards.
*/
type SuperBoard struct {
	arr [][]Board
}

func NewSuperBoard() SuperBoard {
	var a [][]Board = make([][]Board, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]Board, 3)
		for j := 0; j < 3; j++ {
			a[i][j] = NewBoard()
		}
	}

	return SuperBoard{arr: a}
}

// toString method for SuperBoard struct
func (b SuperBoard) String() string {
	var sb strings.Builder

	for sri := 0; sri < 3; sri++ {
		for ri := 0; ri < 3; ri++ {
			for sci := 0; sci < 3; sci++ {
				s := string(b.arr[sri][sci].arr[ri])
				s = strings.Trim(s, "[]")
				sb.WriteString(s)
				if sci < 2 {
					sb.WriteString(" | ")
				}
			}
			if ri < 2 {
				sb.WriteRune('\n')
			}
		}
		if sri < 2 {
			sb.WriteString("\n" + strings.Repeat("-", 15) + "\n")
		}
	}

	return sb.String()
}

func (b SuperBoard) set(i, j, k, l int, marker rune) {
	b.arr[i][j].set(k, l, marker)
}

func (b SuperBoard) get(i, j, k, l int) rune {
	return b.arr[i][j].get(k, l)
}
