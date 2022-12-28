package main

import "fmt"

var (
	// Create a Data Structure to hold the board
	// X - X - O
	// O - X - X
	// X - O - O
	board [3][3]int
	x     int = 1
	o     int = 2
)

func main() {
	fmt.Println("Tic Tac Toe Game")
	newBoard()
	printBoard()
	fmt.Printf("Board: %v", board)
}

// newBoard() - Initialize the board
func newBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = 0
		}
	}
}

// printBoard() - Print the board
func printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v | ", board[i][j])
		}
		fmt.Println()
		fmt.Println("-----------")
	}
}
