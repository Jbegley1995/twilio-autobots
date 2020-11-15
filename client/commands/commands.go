package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/urfave/cli"
)

const (
	ORIGINS_FLAG          string = "origins"
	ORIGIN_PARAMETER_NAME string = "origin"
)

type SearchResponse struct {
	Data   map[string]int
	Errors map[string]string
}

// SearchCommand builds out the command for searching via the cli.
func SearchCommand() cli.Command {
	myFlags := []cli.Flag{
		cli.StringSliceFlag{
			Name:  ORIGINS_FLAG,
			Value: &cli.StringSlice{},
		},
	}

	return cli.Command{
		Name:   "search",
		Usage:  "Looks up the stars on a particular github repository",
		Action: searchCommandAction,
		Flags:  myFlags,
	}
}

func searchCommandAction(c *cli.Context) error {
	originsInput := c.StringSlice(ORIGINS_FLAG)
	urlValues := url.Values{}
	for _, origin := range originsInput {
		urlValues.Add(ORIGIN_PARAMETER_NAME, origin)
	}

	resp, err := http.Get("http://localhost:8080/api/v1/search?" + urlValues.Encode())
	if err != nil {
		return err
	}

	searchResponse := SearchResponse{}
	decErr := json.NewDecoder(resp.Body).Decode(&searchResponse)
	if decErr != nil && decErr != io.EOF {
		return err
	}

	for repo, stars := range searchResponse.Data {
		fmt.Printf("%v currently has %v stars\n", repo, stars)
	}
	return nil
}
