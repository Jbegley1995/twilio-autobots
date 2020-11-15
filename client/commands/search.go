package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

const (
	host      string = "http://localhost"
	port      string = "57865"
	subdomain string = "api/v1/search"

	// OriginsFlag is a flag that can be used to interact with the CLI when using the search functionality.
	OriginsFlag string = "origin"
	// OriginsFileFlag is a flag that can be used to interact with the CLI via a file when using the search functionality.
	OriginsFileFlag string = "origin-file"
	// OriginParameterName is the parameter that needs to be passed to the API.
	OriginParameterName string = "origin"
)

var (
	SearchNoOriginsError             error = errors.Errorf("Origins aren't being passed, please make sure you are passing origins. Use the flag \"%v\"", OriginsFlag)
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
			Name:  OriginsFlag,
			Usage: "A GitHub repository origin string in the form of \"organization/repository\"",
		},
		cli.StringFlag{
			Name:  OriginsFileFlag,
			Usage: "A path to a file, which contains a list of GitHub repository origin string in the form of \"organization/repository\". Each origin will be picked up on a new line.",
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

// searchReadFilesFromOriginFile opens up a file at the origin (or returns empty if there is none), and returns a list of origin strings.
func searchReadFilesFromOriginFile(originFile string) ([]string, error) {
	var (
		origins = []string{}
	)
	if len(originFile) == 0 {
		return origins, nil
	}

	file, err := os.Open(originFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		origins = append(origins, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return origins, nil
}

// searchCommandAction performs search functionality based on the origins passed. This will call the autobots API, in order to retrieve information from github.
func searchCommandAction(c *cli.Context) error {
	originsInput := c.StringSlice(OriginsFlag)

	inputFromFile, err := searchReadFilesFromOriginFile(c.String(OriginsFileFlag))
	if err != nil {
		return err
	}
	if len(inputFromFile) > 0 {
		originsInput = append(originsInput, inputFromFile...)
	}

	if len(originsInput) == 0 {
		return SearchNoOriginsError
	}

	urlValues := url.Values{}
	for _, origin := range originsInput {
		if err := validateOrigin(origin); err != nil {
			return err
		}
		urlValues.Add(OriginParameterName, origin)
	}

	// TODO: The way this url / request is working would be abstracted out and configurable.
	var url = fmt.Sprintf("%v:%v/%v?%v", host, port, subdomain, urlValues.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	searchResponse := SearchResponse{}
	decErr := json.NewDecoder(resp.Body).Decode(&searchResponse)
	if decErr != nil && decErr != io.EOF {
		return err
	}

	for repo, stars := range searchResponse.Data {
		fmt.Printf("[✓] %v currently has %v stars\n", repo, stars)
	}
	for repo, information := range searchResponse.Errors {
		fmt.Printf("[x] Stars couldn't be retrieved correctly for %v: %v\n", repo, information)
	}

	return nil
}
