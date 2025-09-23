package main

import (
	"fmt"
	"math"
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
		board:        newBoard(),
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
		if g.board.getSymbolFromPosition(positions[0]) == Symbols[g.activePlayer] &&
			g.board.getSymbolFromPosition(positions[1]) == Symbols[g.activePlayer] &&
			g.board.getSymbolFromPosition(positions[2]) == Symbols[g.activePlayer] {
			return true, &g.activePlayer
		}
	}

	if g.board.isBoardFull() {
		return true, nil
	}

	return false, nil
}

func (g *Game) makeHumanMove() (finished bool, winner *player) {
	g.board.markField(*g.cursor, g.CurrentSymbol())

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

func (g *Game) calculateScores() map[string]int {
	scores := map[string]int{}
	empties := g.board.getEmptyPositions()

	maxVal := func(m map[string]int) int {
		v := math.MinInt
		for _, s := range m {
			if s > v {
				v = s
			}
		}
		return v
	}
	minVal := func(m map[string]int) int {
		v := math.MaxInt
		for _, s := range m {
			if s < v {
				v = s
			}
		}
		return v
	}

	for _, p := range empties {
		ng := NewGame()
		ng.board = g.board
		ng.cursor = p
		ng.activePlayer = g.activePlayer

		finished, _ := ng.makeHumanMove() // switches player if not finished

		s := 0

		// terminal bonus/penalty
		if finished {
			if ng.winner == nil {
				// draw => no adjustment
			} else if *ng.winner == Player1 {
				s += 10000
			} else {
				s -= 10000
			}
			scores[p.String()] = s
			continue
		}

		child := ng.calculateScores()
		if len(child) > 0 {
			if ng.activePlayer == Player1 {
				s += maxVal(child) // P1 will try to maximize
			} else {
				s += minVal(child) // P2 will try to minimize
			}
		}
		scores[p.String()] = s
	}
	return scores
}

func (g *Game) makeComputerMove() (finished bool, winner *player) {
	currentSymbol := g.CurrentSymbol()

	if p := g.checkForWinningMove(currentSymbol); p != nil {
		g.board.markField(*p, currentSymbol)
	} else if p := g.checkForWinningMove(g.OpponentSymbol()); p != nil {
		g.board.markField(*p, currentSymbol)
	} else {
		scores := g.calculateScores()

		min := math.MaxInt
		max := math.MinInt

		maxPos := ""
		minPos := ""
		for posStr, score := range scores {
			if score > max {
				max = score
				maxPos = posStr
			}

			if score < min {
				min = score
				minPos = posStr
			}
		}

		if g.activePlayer == Player1 {
			p, err := stringPosToPosition(maxPos)
			if err != nil {
				panic("could not get position from string")
			}
			g.board.markField(*p, Symbols[Player1])
		} else {
			p, err := stringPosToPosition(minPos)
			if err != nil {
				panic("could not get position from string")
			}
			g.board.markField(*p, Symbols[Player2])
		}
	}

	return g.handleMoveDone()
}
