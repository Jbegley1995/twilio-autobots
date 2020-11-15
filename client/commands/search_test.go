package commands_test

import (
	"os"
	"testing"

	"github.com/jbegley1995/twilio-autobots/client/app"
	"github.com/jbegley1995/twilio-autobots/client/commands"
)

func TestSearchEmpty(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "search")

	if err := app.Run(args); err != commands.SearchNoOriginsError {
		t.Errorf("Error actual = %v, and Expected = %v.", err, commands.SearchNoOriginsError)
	}
}

func TestSearchInvalidOriginOrganizationEmpty(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "search")
	args = append(args, "--origins=/twilio-autobots")

	if err := app.Run(args); err != commands.OrganizationIncorrectFormatError {
		t.Errorf("Error actual = %v, and Expected = %v.", err, commands.OrganizationIncorrectFormatError)
	}
}

func TestSearchInvalidOriginRepositoryEmpty(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "search")
	args = append(args, "--origins=jbegley1995/")

	if err := app.Run(args); err != commands.RepositoryIncorrectFormatError {
		t.Errorf("Error actual = %v, and Expected = %v.", err, commands.RepositoryIncorrectFormatError)
	}
}

func TestSearchInvalidOriginEmpty(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "search")
	args = append(args, "--origins=")

	if err := app.Run(args); err != commands.OriginIncorrectFormatError {
		t.Errorf("Error actual = %v, and Expected = %v.", err, commands.OriginIncorrectFormatError)
	}
}

func TestSearchValidSingle(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "search")
	args = append(args, "--origins=jbegley1995/twilio-autobots")

	if err := app.Run(args); err != nil {
		t.Error(err)
	}
}

func TestSearchValidDouble(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "search")
	args = append(args, "--origins=jbegley1995/twilio-autobots")
	args = append(args, "--origins=google/go-github")

	if err := app.Run(args); err != nil {
		t.Error(err)
	}
}
