package app

import (
	"github.com/jbegley1995/twilio-autobots/client/commands"
	"github.com/urfave/cli"
)

func Run(args []string) error {
	app := cli.NewApp()
	app.Name = "Twilio Autobots CLI"
	app.Usage = "Allows you to query the Twilio Autobots API"

	app.Commands = []cli.Command{
		commands.SearchCommand(),
	}

	if err := app.Run(args); err != nil {
		return err
	}

	return nil
}
