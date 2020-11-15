package main

import (
	"log"
	"os"

	"github.com/jbegley1995/twilio-autobots/client/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Twilio Autobots CLI"
	app.Usage = "Allows you to query the Twilio Autobots API"

	app.Commands = []cli.Command{
		commands.SearchCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		return
	}

}
