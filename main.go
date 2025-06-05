package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var board [3][3]rune
	for i := range board {
		for j := range board[i] {
			board[i][j] = EMPTY
		}
	}

	reader := bufio.NewReader(os.Stdin)

	// game loop
	for {
		printBoard(board)

		// Player move
		var row, col int
		for {
			fmt.Print("Enter your move 'row-id col-id'")
			input, _ := reader.ReadString('\n')
			parts := strings.Fields(input)
			if len(parts) != 2 {
				fmt.Println("Invalid input. Format: row col (e.g., 0 2)")
				continue
			}
			r, err1 := strconv.Atoi(parts[0])
			c, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil || r < 0 || r > 2 || c < 0 || c > 2 || board[r][c] != EMPTY {
				fmt.Println("Invalid move. Try again.")
				continue
			}
			row, col = r, c
			break
		}
		board[row][col] = PLAYER

		if isGameOver(board) {
			break
		}

		// AI move
		fmt.Println("AI is thinking...")
		move := findBestMove(board)
		board[move.row][move.col] = AI

		if isGameOver(board) {
			break
		}
	}
}
