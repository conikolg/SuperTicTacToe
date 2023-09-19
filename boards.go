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
			a[i][j] = '·'
		}
	}

	return Board{
		arr:    a,
		winner: 0,
	}
}

// toString method for Board struct
func (b Board) String() string {
	switch b.winner {
	case 'X':
		return "\\ /\n X \n/ \\"
	case 'O':
		return "/‾\\\n| |\n\\_/"
	}

	var str string
	for i := 0; i < len(b.arr); i++ {
		str += strings.Trim(string(b.arr[i]), "[]") + "\n"
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
	arr    [][]Board
	winner rune
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
		var repr [][]string
		for sci := 0; sci < 3; sci++ {
			miniboard := b.arr[sri][sci].String()
			repr = append(repr, strings.Split(miniboard, "\n"))
		}
		for ri := 0; ri < 3; ri++ {
			for sci := 0; sci < 3; sci++ {
				sb.WriteString(repr[sci][ri])
				if sci != 2 {
					sb.WriteString(" | ")
				} else {
					sb.WriteRune('\n')
				}
			}
		}

		if sri < 2 {
			sb.WriteString(strings.Repeat("—", 15) + "\n")
		}
	}

	return sb.String()[:sb.Len()-1]
}

func (b SuperBoard) set(i, j, k, l int, marker rune) {
	b.arr[i][j].set(k, l, marker)
}

func (b SuperBoard) Set(row, col int, marker rune) {
	i, j, k, l := row/3, col/3, row%3, col%3
	b.set(i, j, k, l, marker)
}

func (b SuperBoard) get(i, j, k, l int) rune {
	return b.arr[i][j].get(k, l)
}
