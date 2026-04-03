package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// 1. Model — app state
type model struct {
	count int
}

// 2. Init — called once at startup
func (m model) Init() tea.Cmd {
	return nil // no initial command
}

// 3. Update — handle messages, return new state + optional cmd
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m.count++
		case "down", "j":
			m.count--
		}
	}

	return m, nil
}

// 4. View — render model as a string
func (m model) View() string {
	return fmt.Sprintf(
		"Counter: %d\n\nPress up/k to increment, down/j to decrement, q to quit.\n",
		m.count,
	)
}

func main() {
	p := tea.NewProgram(model{count: 0})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
