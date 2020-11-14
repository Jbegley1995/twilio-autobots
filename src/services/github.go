package services

import (
	"context"
	"errors"
	"strings"

	"github.com/google/go-github/v32/github"
)

var (
	OriginIncorrectFormatError       error = errors.New("Origin is not passed in the correct format, please make sure to format as \"organization/repository\"")
	OrganizationIncorrectFormatError error = errors.New("Organization is empty, please make sure organization is being passed")
	RepositoryIncorrectFormatError   error = errors.New("Repository is empty, please make sure repository is being passed")
)

type origin struct {
	Organization string
	Repository   string
}

// GithubService The base client to use to communicate with the Github API.
type GithubService struct {
	Client *github.Client
}

// Github Retreives a github service.
func Github() *GithubService {
	var service = GithubService{}
	service.Client = github.NewClient(nil)

	return &service
}

// splitOrigin Splits the organization/repository input, and returns a useable type.
func splitOrigin(repositoryOrigin string) (*origin, error) {
	repoInformation := strings.Split(repositoryOrigin, "/")
	if len(repoInformation) != 2 {
		return nil, OriginIncorrectFormatError
	}
	if strings.Trim(repoInformation[0], "") == "" {
		return nil, OrganizationIncorrectFormatError
	}
	if strings.Trim(repoInformation[1], "") == "" {
		return nil, RepositoryIncorrectFormatError
	}
	return &origin{
		Organization: repoInformation[0],
		Repository:   repoInformation[1],
	}, nil
}

// GetByOrigin Get's a repository given an origin identifier.
func (service *GithubService) GetByOrigin(originIdentifier string) (repo *github.Repository, err error) {
	var (
		input *origin
	)

	if input, err = splitOrigin(originIdentifier); err != nil {
		return nil, err
	}

	if repo, _, err = service.Client.Repositories.Get(context.Background(), input.Organization, input.Repository); err != nil {
		return nil, err
	}

	return repo, nil
}
