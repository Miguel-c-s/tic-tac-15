package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	blank   = 0
	player1 = 'X'
	player2 = 'O'
)

var (
	board         = [15][15]rune{}
	currentPlayer = player1
)

func main() {
	for {
		printBoard()
		fmt.Printf("Player %c, enter row and column (e.g. '3 2'): ", currentPlayer)
		row, col, err := readMove()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !isValidMove(row, col) {
			fmt.Println("Invalid move, try again")
			continue
		}

		board[row][col] = currentPlayer
		if hasWon(currentPlayer, row, col) {
			fmt.Printf("Player %c has won!\n", currentPlayer)
			break
		}

		if isBoardFull() {
			fmt.Println("The game is a draw!")
			break
		}

		// Switch players
		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
}

// Prints the current board to the console
func printBoard() {
	fmt.Println("    0  1  2  3  4  5  6  7  8  9  10 11 12 13 14")
	for i, row := range board {
		fmt.Printf("%2d", i)
		for _, col := range row {
			if col == blank {
				fmt.Printf("   ")
			} else {
				fmt.Printf("%3c", col)
			}
		}
		fmt.Println()
	}
}

// Reads the player's move from the console
func readMove() (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input")
	}

	row, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid row")
	}

	col, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid column")
	}

	if row > 14 || row < 0 || col > 14 || col < 0 {
		return 0, 0, fmt.Errorf("Value out of range")
	}

	return row, col, nil
}

// Returns true if the move is valid (i.e. the space is blank)
func isValidMove(row int, col int) bool {
	return board[row][col] == blank
}

// Returns true if the player has won with the given move
func hasWon(player rune, row int, col int) bool {
	// Check row
	count := 0
	for i := 0; i < 15; i++ {
		if board[row][i] == player {
			count++
		} else {
			count = 0
		}
		if count == 5 {
			return true
		}
	}

	// Check column
	count = 0
	for i := 0; i < 15; i++ {
		if board[i][col] == player {
			count++
		} else {
			count = 0
		}
		if count == 5 {
			return true
		}
	}

	// Check diagonal (top-left to bottom-right)
	count = 0
	for i := -10; i <= 10; i++ {
		if row+i < 0 || row+i >= 15 || col+i < 0 || col+i >= 15 {
			continue
		}
		if board[row+i][col+i] == player {
			count++
		} else {
			count = 0
		}
		if count == 5 {
			return true
		}
	}

	// Check diagonal (top-right to bottom-left)
	count = 0
	for i := -10; i <= 10; i++ {
		if row+i < 0 || row+i >= 15 || col-i < 0 || col-i >= 15 {
			continue
		}
		if board[row+i][col-i] == player {
			count++
		} else {
			count = 0
		}
		if count == 5 {
			return true
		}
	}

	return false
}

// Returns true if the board is full
func isBoardFull() bool {
	for _, row := range board {
		for _, col := range row {
			if col == blank {
				return false
			}
		}
	}
	return true
}
