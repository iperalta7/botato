package main

import (
	"fmt"
	"os"

	"github.com/iperalta7/botato/src"
)

func main() {
	version, err := os.ReadFile("version.txt")
	if err != nil {
		fmt.Println("Error reading version file:", err)
		return
	}
	bot := src.Botato{
		Version: string(version),
	}
	bot.Run()
}
