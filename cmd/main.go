package main

import (
	"fmt"
	"os"

	"example.com/tictactoe"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	game *tictactoe.Game
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		game: tictactoe.NewGame(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.game.GetStage() {
	case tictactoe.Setup:
		return updateSetupView(m, msg)
	case tictactoe.Playing:
		return updatePlayingView(m, msg)
	case tictactoe.Finished:
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
			m.game = tictactoe.NewGame()
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
			m.game.Start(false)
		case "n":
			m.game.Start(true)
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.game.GetStage() {
	case tictactoe.Setup:
		s := "Let's play TicTacToe.\nAre you playing against a computer? [Yn]"
		return s

	case tictactoe.Playing:
		s := "It's " + m.game.GetActiveSymbol() + "'s turn.\n\n"

		s += m.game.GetBoard(true)

		s += "\nPress q to quit.\n"

		return s
	case tictactoe.Finished:
		s := "Game Finished.\n\n"

		s += m.game.GetBoard(false)

		if m.game.GetWinner() != nil {
			if m.game.IsOpponentHuman() {
				s += "\nPlayer " + m.game.GetWinner().String() + " won! "
			} else {
				if *m.game.GetWinner() == tictactoe.Player1 {
					s += "\nYou won! "
				} else {
					s += "\nYou lost! "
				}
			}
		} else {
			s += "\nIt's a draw.\n\n"
		}

		s += "New Game? [Yn]"

		return s
	default:
		return "something"
	}
}
