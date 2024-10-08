package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Drink struct {
	Name        string
	Description string
	DrinkSizes  []DrinkSize
}

type DrinkSize struct {
	DrinkID uint
	Size    string  // "Small", "Medium", "Large"
	Price   float32 // Price for this size
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	filePath := "/home/dev/devops-practice/backend/data/drinks.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("failed to open JSON file: %w", err)
		os.Exit(2)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var drinks []Drink

	err = json.Unmarshal(byteValue, &drinks)
	if err != nil {
		log.Fatal("failed to parse JSON: %w", err)
		os.Exit(2)
	}
	items := make([]list.Item, len(drinks))

	for index, drink := range drinks {
		drink_item := item{
			title: drink.Name,
			desc:  drink.Description,
		}
		items[index] = drink_item
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Coffee2Go"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
