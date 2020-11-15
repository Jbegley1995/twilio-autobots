package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

const (
	ORIGINS_FLAG          string = "origins"
	ORIGIN_PARAMETER_NAME string = "origin"
)

var (
	SearchNoOriginsError             error = errors.Errorf("Origins aren't being passed, please make sure you are passing origins. Use the flag \"%v\"", ORIGINS_FLAG)
	OriginIncorrectFormatError       error = errors.New("Origin is not passed in the correct format, please make sure to format as \"organization/repository\"")
	OrganizationIncorrectFormatError error = errors.New("Organization is empty, please make sure organization is being passed")
	RepositoryIncorrectFormatError   error = errors.New("Repository is empty, please make sure repository is being passed")
)

// SearchResponse is the structure to recieve a response from the search api.
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

// validateOrigin Splits the organization/repository input, and validates that it is correct.
func validateOrigin(repositoryOrigin string) error {
	repoInformation := strings.Split(repositoryOrigin, "/")
	if len(repoInformation) != 2 {
		return OriginIncorrectFormatError
	}
	if strings.Trim(repoInformation[0], "") == "" {
		return OrganizationIncorrectFormatError
	}
	if strings.Trim(repoInformation[1], "") == "" {
		return RepositoryIncorrectFormatError
	}
	return nil
}

func searchCommandAction(c *cli.Context) error {
	originsInput := c.StringSlice(ORIGINS_FLAG)
	if len(originsInput) == 0 {
		return SearchNoOriginsError
	}

	urlValues := url.Values{}
	for _, origin := range originsInput {
		if err := validateOrigin(origin); err != nil {
			return err
		}
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
	for repo, information := range searchResponse.Errors {
		fmt.Printf("Stars couldn't be retrieved correctly for %v: %v", repo, information)
	}

	return nil
}
