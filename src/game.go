package src

import (
	"errors"
	"fmt"
)

//Game struct
type Game struct {
	players       [2]string
	currentPlayer string
	rolls         int
	winner        string
	board         [3][3]string
}

//NewGame is the way of creating a new game
func NewGame() Game {
	return Game{
		players:       [2]string{"X", "O"},
		currentPlayer: "X",
		rolls:         0,
	}
}

//Play is a method to play a game
func (g *Game) Play(row, column int) error {
	if g.winner != "" {
		return errors.New("Game is over")
	}
	if g.board[row][column] != "" {
		return errors.New("Position taken")
	}
	g.board[row][column] = g.currentPlayer
	g.rolls++
	g.evaluate()
	g.alternatePlayers()
	return nil
}

func (g *Game) alternatePlayers() {
	if g.currentPlayer == g.players[0] {
		g.currentPlayer = g.players[1]
	} else {
		g.currentPlayer = g.players[0]
	}
}

//ShowBoard returns the actual board
func (g *Game) ShowBoard() string {
	return fmt.Sprintf("%s", g.board)
}

//Status return the atual game status
func (g *Game) Status() string {
	if g.winner != "" {
		return fmt.Sprintf("%s won", g.winner)
	}

	if g.rolls == 9 {
		return "Tie"
	}

	return fmt.Sprintf("%s turn", g.currentPlayer)
}

func (g *Game) evaluate() {
	if isDiagonal1Equal(&g.board) {
		g.winner = g.board[0][0]
	}

	if isDiagonal2Equal(&g.board) {
		g.winner = g.board[2][0]
	}

	for i := 0; i < 3; i++ {
		if isRowEqual(&g.board[i]) {
			g.winner = g.board[i][0]
			break
		}

		if isColEqual(&g.board, i) {
			g.winner = g.board[0][i]
			break
		}
	}
}

func isDiagonal1Equal(board *[3][3]string) bool {
	return board[0][0] == board[1][1] && (board[0][0] == board[2][2])
}

func isDiagonal2Equal(board *[3][3]string) bool {
	return board[2][0] == board[1][1] && (board[2][0] == board[0][2])
}

func isRowEqual(row *[3]string) bool {
	return row[0] != "" && row[0] == row[1] && (row[0] == row[2])
}

func isColEqual(board *[3][3]string, i int) bool {
	return board[0][i] != "" && board[0][i] == board[1][i] && (board[0][i] == board[2][i])
}
