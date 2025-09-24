package tictactoe

import "strconv"

var Symbols map[player]string = map[player]string{Player1: "x", Player2: "o"}

type player int

const (
	_       player = iota
	Player1 player = iota
	Player2 player = iota
)

func (p player) String() string {
	return strconv.Itoa(int(p))
}
