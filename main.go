package main

import (
	"fmt"
)

var (
	p1BoardValue = 1
	p2BoardValue = 2
	emptySpace   = 0
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
		// var numStr []string
		for j := 0; j < 3; j++ {
			if board[i][j] == 2 {
				fmt.Printf("  X  | ")
			} else if board[i][j] == 1 {
				fmt.Printf("  O  | ")
			} else if board[i][j] == 0 {
				fmt.Printf(" %d-%d | ", i+1, j+1)
			} else {
				panic("Invalid board")
			}
		}
		fmt.Println()
		fmt.Println("--------------------")
	}
}

func main() {
	fmt.Println("Welcome to Tic Tac Toe")

	var playerOneName string
	fmt.Printf("Player 1, Please enter your name: ")
	_, errPOne := fmt.Scanf("%s", &playerOneName)
	if errPOne != nil {
		fmt.Println("Invalid input, please try again")
		return
	}

	var playerTwoName string
	fmt.Printf("Player 2, Please enter your name: ")
	_, errPTwo := fmt.Scanf("%s", &playerTwoName)
	if errPTwo != nil {
		fmt.Println("Invalid input, please try again")
		return
	}

	// Store the Current Player Name and Board Value
	cpMap := map[int]string{
		p1BoardValue: playerOneName,
		p2BoardValue: playerTwoName,
	}

	repMap := map[int]string{
		p1BoardValue: "X",
		p2BoardValue: "O",
	}

	fmt.Printf("Lets start with Player One - %s \n", playerOneName)
	cpBv := p1BoardValue

	fmt.Printf("Player %v will be represented by %s \n", cpMap[cpBv], repMap[cpBv])

	// Get Current Player Name from the map

	newBoard()

	for {
		currentPlayerName := cpMap[cpBv]
		fmt.Println("Current Board :")
		printBoard()

		fmt.Printf("Hello Player %s, Please enter your move in row and column format (e.g. 1 2): ", currentPlayerName)
		var row, col int
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}

		// Check if the move is valid
		if row < 1 || row > 3 || col < 1 || col > 3 {
			fmt.Println("Invalid input, please try again")
			continue
		}

		// Check if the place is not already occupied
		if board[row-1][col-1] != emptySpace {
			fmt.Println("Invalid input, please try again (place already occupied)")
			continue
		}

		// Update the board
		updateBoard(row, col, cpBv)

		// Check if someone won the game
		isWon := checkWon(currentPlayerName)
		if isWon {
			updateBoard(row, col, cpBv)
			printBoard()
			fmt.Println("Congrats Player ", currentPlayerName, ", You won the game")
			break
		}

		// Check if the game is draw
		isDraw := checkDraw()
		if isDraw {
			updateBoard(row, col, cpBv)
			printBoard()
			fmt.Println("Game is draw")
			break
		}

		// Alternate the player
		if cpBv == p1BoardValue {
			cpBv = p2BoardValue
		} else {
			cpBv = p1BoardValue
		}
	}
}

// updateBoard - updates the board with the move
func updateBoard(row, col, player int) {
	board[row-1][col-1] = player
}

// wonMessage - prints the message when someone wins the game
func wonMessage(player string) {
	fmt.Println("Congrats Player ", player, ", You won the game")
}

func checkWon(player string) (isWon bool) {
	// Check row
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] &&
			board[i][1] == board[i][2] &&
			board[i][0] != emptySpace {
			wonMessage(player)
			return true
		}
	}

	// Check column
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] &&
			board[1][i] == board[2][i] &&
			board[0][i] != emptySpace {
			wonMessage(player)
			return true
		}
	}

	// Check diagonal - 1
	if board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] &&
		board[0][0] != emptySpace {
		wonMessage(player)
		return true
	}
	// Check diagonal - 2
	if board[0][2] == board[1][1] &&
		board[1][1] == board[2][0] &&
		board[0][2] != emptySpace {
		wonMessage(player)
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
