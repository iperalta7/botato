package main

import (
	"log"
	"os"

	"github.com/iperalta7/botato/src"
	"github.com/joho/godotenv"
)

func main() {
	version, err := os.ReadFile("version.txt")
	if err != nil {
		log.Println("Error reading version file:", err)
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	seleniumPath := os.Getenv("SELENIUM_PATH")
	if seleniumPath == "" {
		log.Println("SELENIUM_PATH environment variable is not set")
		return
	}

	geckoDriverPath := os.Getenv("GECKO_DRIVER_PATH")
	if geckoDriverPath == "" {
		log.Println("GECKO_DRIVER_PATH environment variable is not set")
		return
	}

	bot := src.Botato{
		Version:         string(version),
		SeleniumPath:    seleniumPath,
		GeckoDriverPath: geckoDriverPath,
	}
	bot.Run()
}
