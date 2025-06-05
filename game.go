package main

import "fmt"

const (
	EMPTY  = ' '
	PLAYER = 'X'
	AI     = 'O'
)

type Move struct {
	row, col int
}

func printBoard(board [3][3]rune) {
	fmt.Println("Board:")
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print("_ ")
			} else {
				fmt.Printf("%c ", cell)
			}
		}
		fmt.Println()
	}
}

// calls evaluate sees if game is over and prints winner
func isGameOver(board [3][3]rune) bool {
	score := evaluate(board)
	if score == 10 {
		printBoard(board)
		fmt.Println("AI wins!")
		return true
	} else if score == -10 {
		printBoard(board)
		fmt.Println("You win!")
		return true
	} else if !isMovesLeft(board) {
		printBoard(board)
		fmt.Println("It's a draw!")
		return true
	}
	return false
}

// checks for any empty cells in the board
func isMovesLeft(board [3][3]rune) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == EMPTY {
				return true
			}
		}
	}
	return false
}

// returns 10 if AI wins
// returns -10 if human wins
// else 0
func evaluate(board [3][3]rune) int {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			if board[i][0] == AI {
				return +10
			} else if board[i][0] == PLAYER {
				return -10
			}
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			if board[0][i] == AI {
				return +10
			} else if board[0][i] == PLAYER {
				return -10
			}
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == AI {
			return +10
		} else if board[0][0] == PLAYER {
			return -10
		}
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == AI {
			return +10
		} else if board[0][2] == PLAYER {
			return -10
		}
	}
	return 0
}
