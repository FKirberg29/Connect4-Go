///////////////////////////////////////////////////////
//
// Fabian Kirberg
//
//////////////////////////////////////////////////////

package connect4

import (
	"fmt"
	"os"
	"strings"
)

const MAX_SIZE int = 16

type Connect4 struct {
	num_rows    int
	num_columns int
	win_length  int
	totalSlots  int
	board       []int
}

// Constructor
func New(num_rows int, num_columns int, win_length int) *Connect4 {
	c4 := new(Connect4)
	c4.num_rows = num_rows
	c4.num_columns = num_columns
	c4.win_length = win_length
	c4.totalSlots = num_rows * num_columns
	c4.board = make([]int, c4.totalSlots)
	return c4
}

// Helper method to print column labels
func (c4 Connect4) Header() {
	fmt.Println(" A B C D E F G H I J K L M N O P"[0 : c4.num_columns*2])
}

// //Helper method to create and fill board with 0s
// func (c4 Connect4) MakeBoard() {
// 	//c4.board = Array.new(c4.num_rows){Array.new(c4.num_columns, 0)}
// }

// Helper method that prints the formatted board and header
func (c4 Connect4) Printboard() {
	for i := 0; i < c4.totalSlots; i += c4.num_columns {
		rowSlice := c4.board[i : i+c4.num_columns]
		fmt.Println(rowSlice)
	}
	c4.Header()
}

// Helper method that places player token in the lowest possible row of the selected column
func (c4 Connect4) MakeMove(player int, input int) {
	for i := c4.num_rows - 1; i >= 0; i-- {
		if c4.board[i*c4.num_columns+input] == 0 {
			c4.board[i*c4.num_columns+input] = player
			break
		}
	}
	c4.Printboard()
}

// Helper method that checks if a player has won the game
func (c4 Connect4) XInARow(player int) {
	//Use .each to loop through every value of the board
	for i := 0; i < c4.num_rows; i++ {
		for j := 0; j < c4.num_columns; j++ {
			var up_right_counter = 1
			var down_right_counter = 1
			var horizontal_counter = 1
			var vertical_counter = 1
			var currentIndex = c4.board[i*c4.num_columns+j]

			//Only check for a win if the current board value is a player's token
			if currentIndex == player {

				// Up Right Diagonal Check. Only runs if board has enough room for a win.
				if i >= c4.win_length-1 {
					for up_right_counter < c4.win_length && currentIndex == player && c4.board[((i*c4.num_columns)+j)-(up_right_counter*c4.num_columns)+up_right_counter] == player {
						up_right_counter += 1
					}
				}

				// Down Right Diagonal and Vertical Check. Only runs if board has enough room for a win.
				if i <= c4.num_rows-c4.win_length {
					for down_right_counter < c4.win_length && currentIndex == player && c4.board[((i*c4.num_columns)+j)+(down_right_counter*c4.num_columns)+down_right_counter] == player {
						down_right_counter += 1
					}
					for vertical_counter < c4.win_length && currentIndex == player && c4.board[((i*c4.num_columns)+j)+(vertical_counter*c4.num_columns)] == player {
						vertical_counter += 1
					}
				}

				// Horizontal Check
				for horizontal_counter < c4.win_length && currentIndex == player && (((i*c4.num_columns)+j)+horizontal_counter) < c4.totalSlots && c4.board[((i*c4.num_columns)+j)+horizontal_counter] == player {
					horizontal_counter += 1
				}
			}

			//If any of the directional checks reached the winLength, declare a winner and end the game.
			if up_right_counter == c4.win_length || down_right_counter == c4.win_length || horizontal_counter == c4.win_length || vertical_counter == c4.win_length {
				fmt.Println("Congratulations, Player ", player, ". You win.")
				os.Exit(0)
			}
		}
	}
}

// Player column selection loop
func (c4 Connect4) PlaceToken(player int) {
	var inputString string
	var input int

	//Get user input
	fmt.Printf("Player %d, which Column?\n", player)
	fmt.Scanln(&inputString)
	inputString = strings.ToUpper(inputString)
	if len(inputString) > 0 {
		input = int(inputString[0])
	}

	//If user inputs "q" or "Q", quit
	if inputString == "Q" {
		fmt.Println("Goodbye.")
		os.Exit(0)
	}

	//If character chosen is invalid
	if inputString == "" || input <= 64 || input > 64+c4.num_columns {
		fmt.Println("Invalid input, please try again.")
		c4.PlaceToken(player)
	}

	//If character chosen is within the bounds of the game
	if input > 64 && input <= 64+c4.num_columns {
		if c4.board[input-65] != 0 {
			fmt.Println("This column is full, please select another.")
			c4.PlaceToken(player)
		}
		//Place the token, print the board, and check for a win.
		c4.MakeMove(player, input-65)
		c4.XInARow(player)
	}

	// If a player successfully places their token, swap players and request new input.
	if player == 1 {
		c4.PlaceToken(2)
	} else {
		c4.PlaceToken(1)
	}
}

func (c4 Connect4) PlayGame() {
	fmt.Printf("Connect 4 with %d rows, %d columns, and a win length of %d.\n", c4.num_rows, c4.num_columns, c4.win_length)
	//Print the initial board
	c4.Printboard()
	//Begin the game with player 1
	c4.PlaceToken(1)
}
