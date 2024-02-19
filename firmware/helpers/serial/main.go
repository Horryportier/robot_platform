package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"go.bug.st/serial"
)


var (
    exit = false
)

type model struct {
    port serial.Port
}

func initialModel(port_name string, baud_rate int) model {
    mode := &serial.Mode {
        BaudRate: baud_rate,
    }
    port, err := serial.Open(port_name, mode)
    if err != nil {
        log.Fatal(err)
        exit = true
    }

	return model{
        port,
	}
}
func (m model) Init() tea.Cmd {
    bytes_writen, err := m.port.Write([]byte("off"))
    println("bytes_writen", bytes_writen)
    if err != nil {
        log.Println(err)
        exit = true
    }
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:
        
        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit
        }
    }

    if exit {
        return m, tea.Quit
    }

    return m, nil
}

func (m model) View() string {
    buffer := []byte{}
    bytes_read, err :=  m.port.Read(buffer)
    println(bytes_read)
    if err != nil {
        log.Println(err)
        exit = true
    }
    return string(buffer[:])
}

func main() {
    baut_rate, err := strconv.Atoi(os.Args[2])
    if err != nil {
        log.Println(err)
    }
    p := tea.NewProgram(initialModel(os.Args[1],  baut_rate))
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
