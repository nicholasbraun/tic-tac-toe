package tictactoe

import (
	"fmt"
	"strconv"
	"strings"
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
