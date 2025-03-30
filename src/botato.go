package src

import "fmt"

// Botato is a simple bot that does nothing.
type Botato struct {
	Version string
}

func (bot Botato) Run() {
	fmt.Printf("Hello, I'm Botato version %s!", bot.Version)
}
