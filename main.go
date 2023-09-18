package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var board [3][3][3][3]rune
	fmt.Print(displayBoard(board))
}

func displayBoard(board [3][3][3][3]rune) string {
	var sb strings.Builder

	for sri := 0; sri < 3; sri++ {
		for ri := 0; ri < 3; ri++ {
			for sci := 0; sci < 3; sci++ {
				s, _ := json.Marshal(board[sri][sci][ri])
				sb.WriteString(strings.Trim(string(s), "[]"))
				sb.WriteString("  ")
			}
			sb.WriteRune('\n')
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
