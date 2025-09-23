package main

import (
	"fmt"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

var (
	cursorOK    = Green + "H" + Reset
	cursorError = Red + "H" + Reset
)

type position struct {
	row int
	col int
}

func (p *position) String() string {
	return fmt.Sprintf("%d-%d", p.row, p.col)
}

type cursor struct {
	row int
	col int
}

type board [3][3]string

func NewBoard() board {
	return [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
}

func (b *board) getEmptyPositions() []*position {
	emptyPositions := []*position{}

	for row, rowData := range b {
		for col, colData := range rowData {
			if colData == " " {
				emptyPositions = append(emptyPositions, &position{row, col})
			}
		}
	}

	return emptyPositions
}

func (b *board) GetSymbolFromPosition(p position) string {
	return b[p.row][p.col]
}

func (b *board) isBoardFull() bool {
	for _, r := range b {
		for _, c := range r {
			if c == " " {
				return false
			}
		}
	}

	return true
}

func (b *board) String(c *position) string {
	res := ""

	for row, rowData := range b {
		for col, colData := range rowData {
			cell := colData

			if c == nil {
				res += fmt.Sprintf("[%s]", cell)
				continue
			}

			if c.col == col && c.row == row {
				if cell == " " {
					cell = cursorOK
				} else {
					cell = cursorError
				}
			}

			res += fmt.Sprintf("[%s]", cell)
		}

		res += "\n"
	}

	return res
}

func (b *board) MarkField(c position, symbol string) error {
	if b[c.row][c.col] != " " {
		return fmt.Errorf("field is already taken")
	}

	b[c.row][c.col] = symbol
	return nil
}

func (b *board) getPosition(p position) string {
	return b[p.row][p.col]
}
