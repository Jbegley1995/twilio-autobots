package main

import (
	"fmt"
	"os"

	"github.com/jbegley1995/twilio-autobots/client/app"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
