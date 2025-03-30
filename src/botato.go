package src

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
)

const port = 8080

// Botato is a simple bot that does nothing.
type Botato struct {
	Version         string
	SeleniumPath    string
	GeckoDriverPath string
}

func (bot Botato) Run() {
	log.Printf("Hello, I'm Botato version %s!", bot.Version)

	opts := []selenium.ServiceOption{ // TODO: make a factory to make options based on OS
		//selenium.StartFrameBuffer(),
		selenium.GeckoDriver(bot.GeckoDriverPath),
		selenium.Output(nil),
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(bot.SeleniumPath, port, opts...)
	if err != nil {
		log.Printf("Error starting the Selenium server: %v", err)
		return
	}

	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"} // TODO: factory to make browser capabilities
	// caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	if err := wd.Get("http://www.google.com"); err != nil {
		log.Printf("Error navigating to Google: %v", err)
		return
	}
	log.Println("Successfully navigated to Google!")
}
