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
	g.evaluateGame()
	g.alternatePlayes()
	return nil
}

func (g *Game) alternatePlayes() {
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

func (g *Game) evaluateGame() {
	var row string
	var col string
	diagonalA := fmt.Sprintf("%s%s%s", g.board[0][0], g.board[1][1], g.board[2][2])
	diagonalB := fmt.Sprintf("%s%s%s", g.board[2][0], g.board[1][1], g.board[0][2])

	for _, player := range g.players {
		if diagonalA == fmt.Sprintf("%s%s%s", player, player, player) {
			g.winner = player
		}

		if diagonalB == fmt.Sprintf("%s%s%s", player, player, player) {
			g.winner = player
		}

		for i := 0; i < 3; i++ {
			row = fmt.Sprintf("%s%s%s", g.board[i][0], g.board[i][1], g.board[i][2])
			col = fmt.Sprintf("%s%s%s", g.board[0][i], g.board[1][i], g.board[2][i])
			if row == fmt.Sprintf("%s%s%s", player, player, player) {
				g.winner = player
			}

			if col == fmt.Sprintf("%s%s%s", player, player, player) {
				g.winner = player
			}
		}
	}
}
