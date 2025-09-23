package main

import (
	"fmt"
	"strconv"
	"strings"
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

func stringPosToPosition(posStr string) (*position, error) {
	spitString := strings.Split(posStr, "-")

	row, err := strconv.Atoi(spitString[0])
	if err != nil {
		return nil, err
	}

	col, err := strconv.Atoi(spitString[1])
	if err != nil {
		return nil, err
	}

	return &position{row, col}, nil
}

type board [3][3]string

func newBoard() board {
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

func (b *board) getSymbolFromPosition(p position) string {
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

func (b *board) markField(c position, symbol string) error {
	if b[c.row][c.col] != " " {
		return fmt.Errorf("field is already taken")
	}

	b[c.row][c.col] = symbol
	return nil
}

func (b *board) getPosition(p position) string {
	return b[p.row][p.col]
}
