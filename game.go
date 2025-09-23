package main

import (
	"fmt"
	"strconv"
)

type player int

const (
	Player1 player = iota
	Player2 player = iota
)

const emptyField = " "

var winningPositions = [8][3]position{
	// collumns
	{position{0, 0}, position{1, 0}, position{2, 0}},
	{position{0, 1}, position{1, 1}, position{2, 1}},
	{position{0, 2}, position{1, 2}, position{2, 2}},

	// rows
	{position{0, 0}, position{0, 1}, position{0, 2}},
	{position{1, 0}, position{1, 1}, position{1, 2}},
	{position{2, 0}, position{2, 1}, position{2, 2}},

	// diagonals
	{position{0, 0}, position{1, 1}, position{2, 2}},
	{position{0, 2}, position{1, 1}, position{2, 0}},
}

func (p player) String() string {
	return strconv.Itoa(int(p))
}

var Symbols map[player]string = map[player]string{Player1: "x", Player2: "o"}

type Stage int

const (
	Setup    Stage = iota
	Playing  Stage = iota
	Finished Stage = iota
)

type Game struct {
	activePlayer  player
	humanOpponent bool
	board         board
	stage         Stage
	cursor        *position
	winner        *player
	states        []board
}

func (g *Game) getStates() string {
	res := ""

	for _, state := range g.states {
		res += fmt.Sprintf("%v\n", state.String(nil))
	}

	return res
}

func (g *Game) CurrentSymbol() string {
	return Symbols[g.activePlayer]
}

func (g *Game) OpponentSymbol() string {
	if g.activePlayer == Player1 {
		return Symbols[Player2]
	}

	return Symbols[Player1]
}

func NewGame() *Game {
	return &Game{
		activePlayer: Player1,
		board:        NewBoard(),
		stage:        Setup,
		cursor:       &position{0, 0},
	}
}

func (g *Game) MoveCursorRight() {
	if g.cursor.col < len(g.board)-1 {
		g.cursor.col++
	}
}

func (g *Game) MoveCursorLeft() {
	if g.cursor.col > 0 {
		g.cursor.col--
	}
}

func (g *Game) MoveCursorDown() {
	if g.cursor.row < len(g.board[0])-1 {
		g.cursor.row++
	}
}

func (g *Game) MoveCursorUp() {
	if g.cursor.row > 0 {
		g.cursor.row--
	}
}

func (g *Game) isGameFinished() (isFinished bool, winner *player) {
	for _, positions := range winningPositions {
		if g.board.GetSymbolFromPosition(positions[0]) == Symbols[g.activePlayer] &&
			g.board.GetSymbolFromPosition(positions[1]) == Symbols[g.activePlayer] &&
			g.board.GetSymbolFromPosition(positions[2]) == Symbols[g.activePlayer] {
			return true, &g.activePlayer
		}
	}

	if g.board.isBoardFull() {
		return true, nil
	}

	return false, nil
}

func (g *Game) makeHumanMove() (finished bool, winner *player) {
	g.board.MarkField(*g.cursor, g.CurrentSymbol())

	return g.handleMoveDone()
}

func (g *Game) MakeMove() (finished bool, winner *player) {
	if g.humanOpponent {
		return g.makeHumanMove()
	}

	finished, winner = g.makeHumanMove()
	if finished {
		return finished, winner
	}

	return g.makeComputerMove()
}

func (g *Game) switchPlayer() {
	if g.activePlayer == Player1 {
		g.activePlayer = Player2
	} else {
		g.activePlayer = Player1
	}
}

func (g *Game) checkForWinningMove(symbol string) *position {
	for _, positions := range winningPositions {
		if g.board.getPosition(positions[0]) == symbol &&
			g.board.getPosition(positions[1]) == symbol &&
			g.board.getPosition(positions[2]) == emptyField {

			return &positions[2]
		}

		if g.board.getPosition(positions[0]) == symbol &&
			g.board.getPosition(positions[1]) == emptyField &&
			g.board.getPosition(positions[2]) == symbol {

			return &positions[1]
		}

		if g.board.getPosition(positions[0]) == emptyField &&
			g.board.getPosition(positions[1]) == symbol &&
			g.board.getPosition(positions[2]) == symbol {

			return &positions[0]
		}
	}

	return nil
}

func (g *Game) checkForCorners() *position {
	corners := [4]position{
		{0, 0}, {0, 2}, {2, 0}, {2, 2},
	}

	if g.board.getPosition(corners[0]) == emptyField &&
		g.board.getPosition(corners[1]) == emptyField &&
		g.board.getPosition(corners[2]) == emptyField &&
		g.board.getPosition(corners[3]) == emptyField {
		return &corners[0]
	}

	return nil
}

func (g *Game) handleMoveDone() (isFinished bool, winner *player) {
	g.states = append(g.states, g.board)

	isFinished, winner = g.isGameFinished()
	if isFinished {
		g.winner = winner
		g.stage = Finished

		return
	}

	g.switchPlayer()

	return
}

func (g *Game) makeComputerMove() (finished bool, winner *player) {
	currentSymbol := g.CurrentSymbol()

	if p := g.checkForWinningMove(currentSymbol); p != nil {
		g.board.MarkField(*p, currentSymbol)
		goto afterMove
	}

	if p := g.checkForWinningMove(g.OpponentSymbol()); p != nil {
		g.board.MarkField(*p, currentSymbol)
		goto afterMove
	}

	if g.board.getPosition(position{1, 1}) == emptyField {
		g.board.MarkField(position{1, 1}, currentSymbol)
		goto afterMove
	}

	if emptyCorner := g.checkForCorners(); emptyCorner != nil {
		g.board.MarkField(*emptyCorner, currentSymbol)
		goto afterMove
	}

	// TODO: find best move instead of first empty position
	if emptyPositions := g.board.getEmptyPositions(); len(emptyPositions) > 0 {
		g.board.MarkField(*emptyPositions[0], currentSymbol)
		goto afterMove
	}

afterMove:
	return g.handleMoveDone()
}

func (g Game) simulateGames(p player) []*position {
	positions := []*position{}

	for _, emptyPosition := range g.board.getEmptyPositions() {
		newGame := NewGame()
		newGame.activePlayer = g.activePlayer
		newGame.board = g.board
		newGame.humanOpponent = g.humanOpponent
		newGame.states = g.states
		newGame.cursor = emptyPosition
		newGame.MakeMove()

		if newGame.stage == Finished && newGame.winner != nil && *newGame.winner != p {
		}

		// fmt.Printf("round %d\n%v\n", round, newGame.board.String(nil))

		if !newGame.board.isBoardFull() && newGame.stage != Finished {
			newGame.simulateGames(p)
		}
	}

	return positions
}
