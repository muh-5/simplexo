/*
* MIT License
* 
* Copyright (c) 2023 MUHAMMED HUSSAM
* 
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
* 
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
* 
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/

package main

import (
	/* standard library for go */
	"fmt"
	/* math library */
	"math"
	"os"
)

/* for playing another game */
var ok_game bool = true
/* chose for playing another game */
var game_chos rune = 'N'
/* for playing with ai */
var ok_ai bool
/* chose for playing with ai */
var ai_chos rune = 'Y'

/*
* who is winner
* 0 -> no one and game still run
* 1 -> player one
* 2 -> player two
* 3 -> draw
*/
var winner int = 0

/*
* 1 -> player one
* 2 -> player two
*/
var player int = 1

var play int

/* message that show for output or error */
var msg string

/* main board of the game */
var board [9]int = [9]int{
	0, 0, 0,
	0, 0, 0,
	0, 0, 0,
}

/* function to print board */
func print_board(board [9]int) {
	xo_char := [...]rune{
		'1', '2', '3',
		'4', '5', '6',
		'7', '8', '9',
	}
	for i := 0; i < 9; i++ {
		if (board[i] == 1) {
			xo_char[i] = 'X'
		} else if (board[i] == 2) {
			xo_char[i] = 'O'
		}
	}
	fmt.Println("-----------------")
	fmt.Printf("| %c | | %c | | %c |\n", xo_char[0], xo_char[1], xo_char[2])
	fmt.Println("-----------------")
	fmt.Printf("| %c | | %c | | %c |\n", xo_char[3], xo_char[4], xo_char[5])
	fmt.Println("-----------------")
	fmt.Printf("| %c | | %c | | %c |\n", xo_char[6], xo_char[7], xo_char[8])
	fmt.Println("-----------------")
}

/* 
* simple conditional function to check if any player win 
*
* return 1 when player one win
* return 2 when player two win
* return 0 if there is no winner
*
*/
func check_win(board [9]int) int {
	if (board[0] == 1 && board[1] == 1 && board[2] == 1) {
		/*
		* -----------
		* |x| |x| |x|
		* -----------
		* |4| |5| |6|
		* -----------
		* |7| |8| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 1;
	} else if (board[3] == 1 && board[4] == 1 && board[5] == 1) {
		/*
		* -----------
		* |1| |2| |3|
		* -----------
		* |x| |x| |x|
		* -----------
		* |7| |8| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 1;
	} else if (board[6] == 1 && board[7] == 1 && board[8] == 1) {
		/*
		* -----------
		* |1| |2| |3|
		* -----------
		* |4| |5| |6|
		* -----------
		* |x| |x| |x|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 1;
	} else if (board[0] == 1 && board[3] == 1 && board[6] == 1) {
		/*
		* -----------
		* |x| |2| |3|
		* -----------
		* |x| |5| |6|
		* -----------
		* |x| |8| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 1;
	} else if (board[1] == 1 && board[4] == 1 && board[7] == 1) {
		/*
		* -----------
		* |1| |x| |3|
		* -----------
		* |4| |x| |6|
		* -----------
		* |7| |x| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 1;
	} else if (board[2] == 1 && board[5] == 1 && board[8] == 1) {
		/*
		* -----------
		* |1| |2| |x|
		* -----------
		* |4| |5| |x|
		* -----------
		* |7| |8| |x|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 1;
	} else if (board[0] == 1 && board[4] == 1 && board[8] == 1) {
		/*
		* -----------
		* |x| |2| |3|
		* -----------
		* |4| |x| |6|
		* -----------
		* |7| |8| |x|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 1;
	} else if (board[2] == 1 && board[4] == 1 && board[6] == 1) {
		/*
		* -----------
		* |1| |2| |x|
		* -----------
		* |4| |x| |6|
		* -----------
		* |x| |8| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 1;
	} else if (board[0] == 2 && board[1] == 2 && board[2] == 2) {
		/*
		* -----------
		* |o| |o| |o|
		* -----------
		* |4| |5| |6|
		* -----------
		* |7| |8| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 2;
	} else if (board[3] == 2 && board[4] == 2 && board[5] == 2) {
		/*
		* -----------
		* |1| |2| |3|
		* -----------
		* |o| |o| |o|
		* -----------
		* |7| |8| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 2;
	} else if (board[6] == 2 && board[7] == 2 && board[8] == 2) {
		/*
		* -----------
		* |1| |2| |3|
		* -----------
		* |4| |5| |6|
		* -----------
		* |o| |o| |o|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 2;
	} else if (board[0] == 2 && board[3] == 2 && board[6] == 2) {
		/*
		* -----------
		* |o| |2| |3|
		* -----------
		* |o| |5| |6|
		* -----------
		* |o| |8| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 2;
	} else if (board[1] == 2 && board[4] == 2 && board[7] == 2) {
		/*
		* -----------
		* |1| |o| |3|
		* -----------
		* |4| |o| |6|
		* -----------
		* |7| |o| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 2;
	} else if (board[3] == 2 && board[5] == 2 && board[8] == 2) {
		/*
		* -----------
		* |1| |2| |o|
		* -----------
		* |4| |5| |o|
		* -----------
		* |7| |8| |o|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 2;
	} else if (board[0] == 2 && board[4] == 2 && board[8] == 2) {
		/*
		* -----------
		* |o| |2| |3|
		* -----------
		* |4| |o| |6|
		* -----------
		* |7| |8| |o|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 2;
	} else if (board[2] == 2 && board[4] == 2 && board[6] == 2) {
		/*
		* -----------
		* |1| |2| |o|
		* -----------
		* |4| |o| |6|
		* -----------
		* |o| |8| |9|
		* -----------
		*
		* note: the number show not the same in array
		* 	because array begin with 0
		*/
		return 2;
	} else {
		return 0;
	}
}

/*
* function to check if all elements are taken
* return 1 if all taken
* return 0 if not
*/
func all_taken(board [9]int) int {
	for i := 0; i < 9; i++ {
		if (board[i] == 0) {
			return 0;
		}
	}
	return 1;
}

/*
* AI function to return value of possible choice
* its used minimax algorithm 
* if this play make player2 (AI) win return 1
* if this play make player1 win return -1
* if there is no winner return 0
* then add this value to the possible next play of next player
*/
func ai_ret(board [9]int, player int) float64 {
	/* score of this possible choice */
	var score float64
	var next_player int

	if (check_win(board) == 2) {
		score = 1
	} else if (check_win(board) == 1) {
		score = -1
	} else {
		score = 0
	}
	if (player == 1) {
		next_player = 2
	} else {
		next_player = 1;
	}
	for i := 0; i < 9; i++ {
		if (board[i] == 0) {
			board[i] = next_player
			score += ai_ret(board, next_player)
		}
	}
	return score
}

/* AI function to chose position to play */
func ai_move(board [9]int) int {
	var score float64 = math.Inf(-1)
	var temp_score float64
	var index int

	/* check if any position make player 2 win */
	for i:= 0; i < 9; i++ {
		if (board[i] == 0) {
			board[i] = 2
			if (check_win(board) == 2) {
				return i;
			}
			board[i] = 0;
		}
	}

	/* 
	* check if any position make player 1 win 
	* this you block player 1 position to win
	*/
	for i := 0; i < 9; i++ {
		if (board[i] == 0) {
			board[i] = 1
			if (check_win(board) == 1) {
				return i;
			}
			board[i] = 0;
		}
	}
	
	/* if there is no winner check possible choices by using ai_ret */
	for i := 0; i < 9; i++ {
		if (board[i] == 0) {
			board[i] = 2;
			temp_score = ai_ret(board, 2)
			if (temp_score > score) {
				score = temp_score;
				index = i
			}
		}
	}
	return index;
}

func main() {
	for (ok_game) {
		fmt.Println("simple XO game")
		fmt.Printf("\ndo you want play with ai?(Y/n): ")
		fmt.Scanf("%c", &ai_chos)
		if (ai_chos == 'N' || ai_chos == 'n') {
			ok_ai = false
		} else {
			ok_ai = true
		}

		/* if there is no winner continue */
		for (winner == 0) {
			fmt.Print("\033[H\033[2J")
			fmt.Println("simple XO game")

			/* show message after clear screen */
			fmt.Println(msg)
			/* clear message */
			msg = ""

			print_board(board)

			/* check who will play */
			if (player == 1) {
				fmt.Printf("player %d: ", player)
				fmt.Scan(&play)
				if (play == 0) {
					fmt.Println("good bye, see you soon :)")
					os.Exit(0)
				}
				play = play - 1
				msg = fmt.Sprintf("player %d play at %d", 
				player, play + 1)
			} else {
				/* check if AI will play or another player */
				if (ok_ai) {
					/* 
					* get AI move 
					* by using ai_move function 
					*/
					play = ai_move(board)

					/* print AI play */
					msg = fmt.Sprintf("ai play at %d", 
					play + 1)
				} else {
					fmt.Printf("player %d: ", player)
					fmt.Scan(&play)

					if (play == 0) {
						fmt.Println("good bye, see you soon :)")
						os.Exit(0)
					}
					play = play - 1
					msg = fmt.Sprintf("player %d play at %d",
					player, play + 1)
				}
			}
			/* check if play is out of range */
			if (play < 0 || play > 9 ) {
				msg = "wrong input"
			} else if (board[play] != 0) {
				msg = "this place already used"
			} else {
				/* add play to board */
				board[play] =  player

				/* change player for next play */
				if (player == 1) {
					player = 2
				} else {
					player = 1
				}

				/* check if any player is win */
				if (check_win(board) == 1) {
					winner = 1
				} else if (check_win(board) == 2) {
					winner = 2
				} else if (all_taken(board) == 1) {
					/* 
					* this mean no one win
					* and there is no space in board
					*/
					winner = 3;
				}
			}
		}
		fmt.Println(msg)
		if (winner == 1) {
			fmt.Println("player one win")
		} else if (winner == 2) {
			fmt.Println("player two win")
		} else {
			fmt.Println("draw")
		}
		print_board(board)

		fmt.Printf("\n play another game?(y/N): ")
		fmt.Scanf("%c",&game_chos)
		if (game_chos == 'Y' || game_chos == 'y') {
			/* clean to begin another game */
			ok_game = true
			winner = 0
			board = [9]int{
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			}
			msg = ""
		} else {
			ok_game = false
		}
	}
}
