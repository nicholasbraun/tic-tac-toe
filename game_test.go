package main

import (
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
	t.Run("it sets the game up correctly", func(t *testing.T) {
		g := NewGame()

		if g.activePlayer != Player1 {
			t.Errorf("initial active player should be Player1. got %q", g.activePlayer)
		}

		wantBoard := board{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		}

		assertBoard(t, g.board, wantBoard)

		gotCursor := g.cursor

		if gotCursor.col != 0 && gotCursor.row != 0 {
			t.Errorf("initial cursor should be (0,0). got (%d,%d)", gotCursor.row, gotCursor.col)
		}
	})

	t.Run("it moves the cursor correctly", func(t *testing.T) {
		g := NewGame()

		g.MoveCursorRight()

		wantBoard := "[ ][" + cursorOK + "][ ]\n[ ][ ][ ]\n[ ][ ][ ]\n"
		gotBoard := g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorRight()

		wantBoard = "[ ][ ][" + cursorOK + "]\n[ ][ ][ ]\n[ ][ ][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorRight()

		wantBoard = "[ ][ ][" + cursorOK + "]\n[ ][ ][ ]\n[ ][ ][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorDown()

		wantBoard = "[ ][ ][ ]\n[ ][ ][" + cursorOK + "]\n[ ][ ][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorDown()

		wantBoard = "[ ][ ][ ]\n[ ][ ][ ]\n[ ][ ][" + cursorOK + "]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorDown()

		wantBoard = "[ ][ ][ ]\n[ ][ ][ ]\n[ ][ ][" + cursorOK + "]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorLeft()

		wantBoard = "[ ][ ][ ]\n[ ][ ][ ]\n[ ][" + cursorOK + "][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorLeft()

		wantBoard = "[ ][ ][ ]\n[ ][ ][ ]\n[" + cursorOK + "][ ][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorLeft()

		wantBoard = "[ ][ ][ ]\n[ ][ ][ ]\n[" + cursorOK + "][ ][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorUp()

		wantBoard = "[ ][ ][ ]\n[" + cursorOK + "][ ][ ]\n[ ][ ][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorUp()

		wantBoard = "[" + cursorOK + "][ ][ ]\n[ ][ ][ ]\n[ ][ ][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}

		g.MoveCursorUp()

		wantBoard = "[" + cursorOK + "][ ][ ]\n[ ][ ][ ]\n[ ][ ][ ]\n"
		gotBoard = g.board.String(g.cursor)

		if gotBoard != wantBoard {
			t.Errorf("want board: %q, got: %q", wantBoard, gotBoard)
		}
	})

	t.Run("it ends the game when it is a draw", func(t *testing.T) {
		g := NewGame()

		g.board = board{
			{"x", "x", "o"},
			{"o", "x", "x"},
			{" ", "o", "o"},
		}
		g.activePlayer = Player1
		g.cursor = &position{2, 0}
		g.MakeMove()

		wantBoard := board{
			{"x", "x", "o"},
			{"o", "x", "x"},
			{"x", "o", "o"},
		}

		assertBoard(t, g.board, wantBoard)

		if g.stage != Finished {
			t.Errorf("the game should be finished. got: %v", g.stage)
		}

		if g.winner != nil {
			t.Errorf("the game should be a draw. got winner: %v", g.winner)
		}
	})

	t.Run("it ends the game when a player has won", func(t *testing.T) {
		g := NewGame()

		g.board = board{
			{"x", "x", " "},
			{"o", "o", " "},
			{" ", " ", " "},
		}
		g.activePlayer = Player1
		g.cursor = &position{0, 2}
		g.MakeMove()

		wantBoard := board{
			{"x", "x", "x"},
			{"o", "o", " "},
			{" ", " ", " "},
		}

		assertBoard(t, g.board, wantBoard)

		if g.stage != Finished {
			t.Errorf("game should be finished. got stage %d", g.stage)
		}

		if g.winner == nil {
			t.Fatalf("there should be a winner")
		}

		if *g.winner != Player1 {
			t.Errorf("player1 should be winner. got: %d", *g.winner)
		}
	})
}

func TestComputer(t *testing.T) {
	t.Run("it runs a simulation of all possible moves against the computer and never wins", func(t *testing.T) {
		g := NewGame()
		g.humanOpponent = false

		simulateGameRound(t, *g, 1)
	})

	t.Run("it calculates scores of possible next positions", func(t *testing.T) {
		g := NewGame()
		g.humanOpponent = false

		g.board = board{
			{"x", "x", "o"},
			{"o", "o", " "},
			{"x", " ", " "},
		}
		g.activePlayer = Player1

		got := g.calculateScores()
		fmt.Println(got)

		want12 := 0
		got12 := got["1-2"]

		if got12 != want12 {
			t.Errorf("position 1-2 should have a score of %d. got %d", want12, got12)
		}

		want21 := -10000
		got21 := got["2-1"]

		if got21 != want21 {
			t.Errorf("position 2-1 should have a score of %d. got %d", want21, got21)
		}

		want22 := -10000
		got22 := got["2-2"]

		if got22 != want22 {
			t.Errorf("position 2-2 should have a score of %d. got %d", want22, got22)
		}
	})
}

func assertBoard(t testing.TB, got, want board) {
	t.Helper()

	gotStr := got.String(nil)
	wantStr := want.String(nil)

	if gotStr != wantStr {
		t.Errorf("want board: %v, got: %v", wantStr, gotStr)
	}
}

func simulateGameRound(t testing.TB, game Game, round int) {
	if round > 5 {
		t.Fatalf("can't have more than 5 rounds")
	}

	for _, emptyPosition := range game.board.getEmptyPositions() {
		newGame := NewGame()
		newGame.activePlayer = game.activePlayer
		newGame.board = game.board
		newGame.humanOpponent = game.humanOpponent
		newGame.states = game.states
		newGame.cursor = emptyPosition

		isFinished, winner := newGame.MakeMove()

		if isFinished && winner != nil && *winner == Player1 {
			t.Fatalf("player 1 must not be able to win the game: \n%v", newGame.getStates())
		}

		if !newGame.board.isBoardFull() && newGame.stage != Finished {
			simulateGameRound(t, *newGame, round+1)
		}
	}
}
