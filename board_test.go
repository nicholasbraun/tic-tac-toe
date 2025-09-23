package main

import (
	"testing"
)

func TestBoard(t *testing.T) {
	t.Run("set up an empty board", func(t *testing.T) {
		b := NewBoard()

		got := b.String(nil)
		want := "[ ][ ][ ]\n[ ][ ][ ]\n[ ][ ][ ]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("show a green cursor on the board when the field is empty", func(t *testing.T) {
		b := NewBoard()

		got := b.String(&position{1, 1})
		want := "[ ][ ][ ]\n[ ][" + cursorOK + "][ ]\n[ ][ ][ ]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("show a red cursor on the board when the field is taken", func(t *testing.T) {
		b := NewBoard()
		b.MarkField(position{1, 1}, "x")

		got := b.String(&position{1, 1})
		want := "[ ][ ][ ]\n[ ][" + cursorError + "][ ]\n[ ][ ][ ]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("mark a field with x", func(t *testing.T) {
		b := NewBoard()
		b.MarkField(position{0, 1}, "x")

		got := b.String(nil)
		want := "[ ][x][ ]\n[ ][ ][ ]\n[ ][ ][ ]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("marking a used field returns an error", func(t *testing.T) {
		b := NewBoard()
		err := b.MarkField(position{0, 1}, "x")
		if err != nil {
			t.Fatalf("should not return an error when marking an empty field")
		}

		err = b.MarkField(position{0, 1}, "o")

		if err == nil {
			t.Fatalf("expected an error when marking used field")
		}

		got := b.String(nil)
		want := "[ ][x][ ]\n[ ][ ][ ]\n[ ][ ][ ]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
