package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_it_allows_mark_a_non_take_position(t *testing.T) {
	game := NewGame()
	error := game.Play(1, 1)
	assert.NoError(t, error)
}

func Test_the_board_is_a_3_x_3_grid(t *testing.T) {
	game := NewGame()
	assert.Equal(t, "[[  ] [  ] [  ]]", game.ShowBoard())
}

func Test_it_saves_the_players_turn(t *testing.T) {
	game := NewGame()
	game.Play(0, 1)
	assert.Equal(t, "[[ X ] [  ] [  ]]", game.ShowBoard())
}

func Test_it_does_not_allow_take_a_taken_position(t *testing.T) {
	game := NewGame()
	game.Play(0, 0)
	error := game.Play(0, 0)
	assert.EqualError(t, error, "Position taken")
}

func Test_it_alternates_between_players(t *testing.T) {
	game := NewGame()
	game.Play(0, 0)
	game.Play(0, 1)
	game.Play(0, 2)
	assert.Equal(t, "[[X O X] [  ] [  ]]", game.ShowBoard())
}

func Test_it_find_the_winner_in_the_top_row(t *testing.T) {
	game := NewGame()
	game.Play(0, 0)
	game.Play(1, 0)
	game.Play(0, 1)
	game.Play(1, 1)
	game.Play(0, 2)
	assert.Equal(t, "X won", game.Status())

	game = NewGame()
	game.Play(1, 0)
	game.Play(0, 0)
	game.Play(1, 1)
	game.Play(0, 1)
	game.Play(2, 2)
	game.Play(0, 2)
	assert.Equal(t, "O won", game.Status())
}

func Test_it_find_the_winner_in_the_middle_row(t *testing.T) {
	game := NewGame()
	game.Play(1, 0)
	game.Play(0, 0)
	game.Play(1, 1)
	game.Play(0, 1)
	game.Play(1, 2)
	assert.Equal(t, "X won", game.Status())

	game = NewGame()
	game.Play(0, 0)
	game.Play(1, 0)
	game.Play(0, 1)
	game.Play(1, 1)
	game.Play(2, 2)
	game.Play(1, 2)
	assert.Equal(t, "O won", game.Status())
}

func Test_it_find_the_winner_in_the_bottom_row(t *testing.T) {
	game := NewGame()
	game.Play(2, 0)
	game.Play(0, 0)
	game.Play(2, 1)
	game.Play(0, 1)
	game.Play(2, 2)
	assert.Equal(t, "X won", game.Status())

	game = NewGame()
	game.Play(0, 0)
	game.Play(2, 0)
	game.Play(0, 1)
	game.Play(2, 1)
	game.Play(1, 1)
	game.Play(2, 2)
	assert.Equal(t, "O won", game.Status())
}

func Test_it_find_the_winner_in_the_first_column(t *testing.T) {
	game := NewGame()
	game.Play(0, 0) //[ X  O    ]
	game.Play(0, 1) //[ X  O    ]
	game.Play(1, 0) //[ X       ]
	game.Play(1, 1)
	game.Play(2, 0)
	assert.Equal(t, "X won", game.Status())

	game = NewGame()
	game.Play(0, 1) //[ O  X  X ]
	game.Play(0, 0) //[ O  X    ]
	game.Play(1, 1) //[ O       ]
	game.Play(1, 0)
	game.Play(0, 2)
	game.Play(2, 0)
	assert.Equal(t, "O won", game.Status())
}

func Test_it_find_the_winner_in_the_second_column(t *testing.T) {
	game := NewGame()
	game.Play(0, 1) //[ O  X    ]
	game.Play(0, 0) //[ 0  X    ]
	game.Play(1, 1) //[    X    ]
	game.Play(1, 0)
	game.Play(2, 1)
	assert.Equal(t, "X won", game.Status())

	game = NewGame()
	game.Play(0, 0) //[ X  O  X ]
	game.Play(0, 1) //[ X  O    ]
	game.Play(1, 0) //[    O    ]
	game.Play(1, 1)
	game.Play(0, 2)
	game.Play(2, 1)
	assert.Equal(t, "O won", game.Status())
}

func Test_it_find_the_winner_in_the_third_column(t *testing.T) {
	game := NewGame()
	game.Play(0, 2) //[ O     X ]
	game.Play(0, 0) //[ 0     X ]
	game.Play(1, 2) //[       X ]
	game.Play(1, 0)
	game.Play(2, 2)
	assert.Equal(t, "X won", game.Status())

	game = NewGame()
	game.Play(0, 0) //[ X  X  O ]
	game.Play(0, 2) //[ X     O ]
	game.Play(1, 0) //[       O ]
	game.Play(1, 2)
	game.Play(0, 1)
	game.Play(2, 2)
	assert.Equal(t, "O won", game.Status())
}

func Test_it_find_the_winner_diagonally(t *testing.T) {
	game := NewGame()
	game.Play(0, 0) // [ X  O    ]
	game.Play(0, 1) // [ O  X    ]
	game.Play(1, 1) // [       X ]
	game.Play(1, 0)
	game.Play(2, 2)
	assert.Equal(t, "X won", game.Status())

	game = NewGame()
	game.Play(0, 1) // [ O  X  X ]
	game.Play(0, 0) // [ X  O    ]
	game.Play(1, 0) // [       O ]
	game.Play(1, 1)
	game.Play(0, 2)
	game.Play(2, 2)
	assert.Equal(t, "O won", game.Status())

	game = NewGame()
	game.Play(0, 2) // [    O  X ]
	game.Play(0, 1) // [ O  X    ]
	game.Play(1, 1) // [ X       ]
	game.Play(1, 0)
	game.Play(2, 0)
	assert.Equal(t, "X won", game.Status())

	game = NewGame()
	game.Play(0, 1) // [    X  O ]
	game.Play(0, 2) // [    O  X ]
	game.Play(1, 2) // [ O     X ]
	game.Play(1, 1)
	game.Play(2, 2)
	game.Play(2, 0)
	assert.Equal(t, "O won", game.Status())
}

func Test_it_identifies_a_tie_game(t *testing.T) {
	game := NewGame()
	game.Play(0, 0) //[ X  X  O ]
	game.Play(0, 2) //[ O  O  X ]
	game.Play(0, 1) //[ X  O  X ]
	game.Play(1, 0)
	game.Play(1, 2)
	game.Play(1, 1)
	game.Play(2, 0)
	game.Play(2, 1)
	game.Play(2, 2)
	assert.Equal(t, "Tie", game.Status())
}

func Test_it_identifies_the_players_turn(t *testing.T) {
	game := NewGame()
	assert.Equal(t, "X turn", game.Status())
	game.Play(0, 0)
	assert.Equal(t, "O turn", game.Status())
}

func Test_it_does_not_play_a_finished_game(t *testing.T) {
	game := NewGame()
	game.Play(0, 0) //[ X  O    ]
	game.Play(0, 1) //[ X  O    ]
	game.Play(1, 0) //[ X       ]
	game.Play(1, 1)
	game.Play(2, 0)
	error := game.Play(2, 2)
	assert.EqualError(t, error, "Game is over")
}
