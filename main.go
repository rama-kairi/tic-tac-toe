package main

import "fmt"

var (
	player1    = 1
	player2    = 2
	emptySpace = 0
)

var board [3][3]int

// newBoard -  creates a new board
func newBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = emptySpace
		}
	}
}

// printBoard - prints the board
func printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 2 {
				fmt.Printf(" X |")
			} else if board[i][j] == 1 {
				fmt.Printf(" O |")
			} else {
				fmt.Printf("   |")
			}
		}
		fmt.Println()
		fmt.Println("----------")
	}
}

func main() {
	newBoard()
	fmt.Println("Welcome to Tic Tac Toe")

	currentPLayer := 1
	for {
		fmt.Println("Current Board :")
		printBoard()

		fmt.Printf("Hello Player %d, Please enter your move in row and column format (e.g. 1 2): ", currentPLayer)
		var row, col int
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}

		if row < 1 || row > 3 || col < 1 || col > 3 {
			fmt.Println("Invalid input, please try again")
			continue
		}

		updateBoard(row, col, currentPLayer)

		// Check if someone won the game
		isWon := checkWon()
		if isWon {
			updateBoard(row, col, currentPLayer)
			printBoard()
			break
		}

		// Check if the game is draw
		isDraw := checkDraw()
		if isDraw {
			updateBoard(row, col, currentPLayer)
			printBoard()
			fmt.Println("Game is draw")
			break
		}

		// Alternate the player
		if currentPLayer == 1 {
			currentPLayer = 2
		} else {
			currentPLayer = 1
		}
	}
}

// updateBoard - updates the board with the move
func updateBoard(row, col, player int) {
	board[row-1][col-1] = player
}

func checkWon() (isWon bool) {
	// Check row
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] &&
			board[i][1] == board[i][2] &&
			board[i][0] != emptySpace {
			fmt.Println("Congrats Player ", board[i][0], ", You won the game")
			return true
		}
	}

	// Check column
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] &&
			board[1][i] == board[2][i] &&
			board[0][i] != emptySpace {
			fmt.Println("Player ", board[0][i], " won the game")
			return true
		}
	}

	// Check diagonal
	if board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] &&
		board[0][0] != emptySpace {
		fmt.Println("Player ", board[0][0], " won the game")
		return true
	}

	return false
}

// Check Draw - check if the game is draw
func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == emptySpace {
				return false
			}
		}
	}
	return true
}
