package tictactoe

import (
	"testing"
)

func TestBoard(t *testing.T) {
	t.Run("set up an empty board", func(t *testing.T) {
		b := newBoard()

		want := board{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		}
		assertBoard(t, b, want, nil)
	})

	t.Run("show a green cursor on the board when the field is empty", func(t *testing.T) {
		b := newBoard()

		want := board{
			{" ", " ", " "},
			{" ", cursorOK, " "},
			{" ", " ", " "},
		}
		assertBoard(t, b, want, &position{1, 1})
	})

	t.Run("show a red cursor on the board when the field is taken", func(t *testing.T) {
		b := newBoard()
		b.markField(position{1, 1}, "x")

		want := board{
			{" ", " ", " "},
			{" ", cursorError, " "},
			{" ", " ", " "},
		}
		assertBoard(t, b, want, &position{1, 1})
	})

	t.Run("mark a field with x", func(t *testing.T) {
		b := newBoard()
		b.markField(position{0, 1}, "x")

		want := board{
			{" ", "x", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		}
		assertBoard(t, b, want, nil)
	})

	t.Run("marking a used field returns an error", func(t *testing.T) {
		b := newBoard()
		err := b.markField(position{0, 1}, "x")
		if err != nil {
			t.Fatalf("should not return an error when marking an empty field")
		}

		err = b.markField(position{0, 1}, "o")

		if err == nil {
			t.Fatalf("expected an error when marking used field")
		}

		want := board{
			{" ", "x", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		}
		assertBoard(t, b, want, nil)
	})
}

func assertBoard(t testing.TB, got, want board, cursorPosition *position) {
	t.Helper()

	gotStr := got.String(cursorPosition)
	wantStr := want.String(nil)

	if gotStr != wantStr {
		t.Errorf("want board: %v, got: %v", wantStr, gotStr)
	}
}
