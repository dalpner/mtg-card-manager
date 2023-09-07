package main

import (
	"log"
	"mtg-card-manager/cardSearch"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(cardSearch.InitialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
