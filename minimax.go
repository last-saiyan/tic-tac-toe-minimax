package main

func findBestMove(board [3][3]rune) Move {
	bestVal := -1000
	bestMove := Move{-1, -1}

	for i := range board {
		for j := range board[i] {
			// for each empty cell
			if board[i][j] == EMPTY {
				board[i][j] = AI
				// calculate the value for the move
				moveVal := minimax(board, false)
				// for debugging
				// fmt.Printf("i %d, j %d, moveVal %d, \n", i, j, moveVal)
				board[i][j] = EMPTY
				// get the move with the best value
				if moveVal > bestVal {
					bestMove = Move{i, j}
					bestVal = moveVal
				}
			}
		}
	}
	return bestMove
}

// isMax:
// true if it’s AI’s turn (maximize), coz ai requires +10 from evaluate
// false if it’s the player’s turn (minimize). coz human requires -10 from evaluate
func minimax(board [3][3]rune, isMax bool) int {
	score := evaluate(board)
	// recursive function base case
	// stop if no more moves/if there is a winner
	if score == 10 || score == -10 {
		return score
	}
	if !isMovesLeft(board) {
		return 0
	}

	if isMax {
		best := -1000
		for i := range board {
			for j := range board[i] {
				if board[i][j] == EMPTY {
					board[i][j] = AI
					curr := minimax(board, false)
					if curr > best {
						best = curr
					}
					board[i][j] = EMPTY
				}
			}
		}
		return best
	} else {
		best := 1000
		for i := range board {
			for j := range board[i] {
				if board[i][j] == EMPTY {
					board[i][j] = PLAYER
					curr := minimax(board, true)
					if curr < best {
						best = curr
					}
					board[i][j] = EMPTY
				}
			}
		}
		return best
	}
}
