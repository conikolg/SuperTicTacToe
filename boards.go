package main

import (
	"strings"
)

/*
defines the Board "class", utils for operations dealing with
just a single, simple, TicTacToe board.
*/
const EMPTY rune = '·'

type Board struct {
	arr    [][]rune
	winner rune
}

func NewBoard() Board {
	var a [][]rune = make([][]rune, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]rune, 3)
		for j := 0; j < 3; j++ {
			a[i][j] = EMPTY
		}
	}

	return Board{
		arr:    a,
		winner: 0,
	}
}

// toString method for Board struct
func (b *Board) String() string {
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

func (b *Board) GetWinner() rune {
	return b.winner
}

func (b *Board) computeWinner() rune {
	// check the rows
	for i := 0; i < len(b.arr); i++ {
		if b.arr[i][0] == b.arr[i][1] && b.arr[i][1] == b.arr[i][2] && b.arr[i][0] != EMPTY {
			b.winner = b.arr[i][0]
			return b.winner
		}
	}

	// check the cols
	for i := 0; i < len(b.arr); i++ {
		if b.arr[0][i] == b.arr[1][i] && b.arr[1][i] == b.arr[2][i] && b.arr[0][i] != EMPTY {
			b.winner = b.arr[0][i]
			return b.winner
		}
	}

	// check diagonals
	if b.arr[0][0] == b.arr[1][1] && b.arr[1][1] == b.arr[2][2] && b.arr[1][1] != EMPTY {
		b.winner = b.arr[1][1]
		return b.winner
	}
	if b.arr[2][0] == b.arr[1][1] && b.arr[1][1] == b.arr[0][2] && b.arr[1][1] != EMPTY {
		b.winner = b.arr[1][1]
		return b.winner
	}

	// check if any open spaces left
	for i := 0; i < len(b.arr)*len(b.arr); i++ {
		if b.arr[i/3][i%3] == EMPTY {
			return 0
		}
	}

	// stalemate
	b.winner = '-'
	return b.winner
}

func (b *Board) set(i, j int, marker rune) bool {
	if b.arr[i][j] == EMPTY {
		b.arr[i][j] = marker
		b.computeWinner()
		return true
	}
	return false
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

	sb.WriteString("  1-3   4-6   7-9\n")
	for sri := 0; sri < 3; sri++ {
		var repr [][]string
		for sci := 0; sci < 3; sci++ {
			miniboard := b.arr[sri][sci].String()
			repr = append(repr, strings.Split(miniboard, "\n"))
		}
		for ri := 0; ri < 3; ri++ {
			sb.WriteRune(rune('a' + sri*3 + ri))
			sb.WriteRune(' ')
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
			sb.WriteString("  " + strings.Repeat("—", 15) + "\n")
		}
	}

	return sb.String()[:sb.Len()-1]
}

func (b *SuperBoard) set(i, j, k, l int, marker rune) bool {
	return b.arr[i][j].set(k, l, marker)
}

func (b *SuperBoard) Set(row, col int, marker rune) bool {
	i, j, k, l := row/3, col/3, row%3, col%3
	return b.set(i, j, k, l, marker)
}

func (b *SuperBoard) GetBoard(i int) Board {
	return b.arr[i/3][i%3]
}

func (b *SuperBoard) GetWinner() rune {
	return b.winner
}

func (b *SuperBoard) ComputeWinner() rune {
	// check the rows
	for i := 0; i < len(b.arr); i++ {
		if b.arr[i][0].GetWinner() == b.arr[i][1].GetWinner() && b.arr[i][1].GetWinner() == b.arr[i][2].GetWinner() && b.arr[i][0].GetWinner() != 0 {
			b.winner = b.arr[i][0].GetWinner()
			return b.winner
		}
	}

	// check the cols
	for i := 0; i < len(b.arr); i++ {
		if b.arr[0][i].GetWinner() == b.arr[1][i].GetWinner() && b.arr[1][i].GetWinner() == b.arr[2][i].GetWinner() && b.arr[0][i].GetWinner() != 0 {
			b.winner = b.arr[0][i].GetWinner()
			return b.winner
		}
	}

	// check diagonals
	if b.arr[0][0].GetWinner() == b.arr[1][1].GetWinner() && b.arr[1][1].GetWinner() == b.arr[2][2].GetWinner() && b.arr[1][1].GetWinner() != 0 {
		b.winner = b.arr[1][1].GetWinner()
		return b.winner
	}
	if b.arr[2][0].GetWinner() == b.arr[1][1].GetWinner() && b.arr[1][1].GetWinner() == b.arr[0][2].GetWinner() && b.arr[1][1].GetWinner() != 0 {
		b.winner = b.arr[1][1].GetWinner()
		return b.winner
	}

	// check if any open spaces left
	for i := 0; i < len(b.arr)*len(b.arr); i++ {
		if b.arr[i/3][i%3].GetWinner() == 0 {
			return 0
		}
	}

	// stalemate
	b.winner = '-'
	return b.winner
}
