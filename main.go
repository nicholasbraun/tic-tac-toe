package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	game *Game
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		game: NewGame(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.game.stage {
	case Setup:
		return updateSetupView(m, msg)
	case Playing:
		return updatePlayingView(m, msg)
	case Finished:
		return updateFinishedView(m, msg)
	}

	return m, nil
}

func updatePlayingView(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "left", "h":
			m.game.MoveCursorLeft()

		case "right", "l":
			m.game.MoveCursorRight()

		case "up", "k":
			m.game.MoveCursorUp()

		case "down", "j":
			m.game.MoveCursorDown()

		case "enter", " ":
			m.game.MakeMove()
		}
	}

	return m, nil
}

func updateFinishedView(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q", "n":
			return m, tea.Quit

		case "enter", "y":
			m.game = NewGame()
		}
	}

	return m, nil
}

func updateSetupView(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "y", "enter":
			m.game.humanOpponent = false
			m.game.stage = Playing

		case "n":
			m.game.humanOpponent = true
			m.game.stage = Playing
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.game.stage {
	case Setup:
		s := "Let's play TicTacToe.\nAre you playing against a computer? [Yn]"
		return s

	case Playing:
		s := "It's Player " + m.game.activePlayer.String() + "'s (" + m.game.CurrentSymbol() + ") turn.\n\n"

		s += m.game.board.String(m.game.cursor)

		s += "\nPress q to quit.\n"

		return s
	case Finished:
		s := "Game Finished.\n\n"

		s += m.game.board.String(nil)

		if m.game.winner != nil {
			if *m.game.winner == Player1 {
				s += "\nYou won! "
			} else {
				s += "\nYou lost! "
			}
			// s += "\nPlayer " + m.game.winner.String() + " has won!\n\n"
		} else {
			s += "\nIt's a draw.\n\n"
		}

		s += "New Game? [Yn]"

		return s
	default:
		return "something"
	}
}
