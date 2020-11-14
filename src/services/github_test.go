package services

import (
	"testing"
)

var (
	githubService = Github()
)

func TestGetWithEmptyInput(t *testing.T) {
	if githubService == nil {
		t.Fatal("Service not set up properly.")
	}

	_, err := githubService.GetByOrigin("")

	if err != OriginIncorrectFormatError {
		t.Errorf("Error actual = %v, and Expected = %v.", err, OriginIncorrectFormatError)
	}
}

func TestGetWithEmptyOrganization(t *testing.T) {
	if githubService == nil {
		t.Fatal("Service not set up properly.")
	}

	_, err := githubService.GetByOrigin("/twilio-autobots")

	if err != OrganizationIncorrectFormatError {
		t.Errorf("Error actual = %v, and Expected = %v.", err, OrganizationIncorrectFormatError)
	}
}

func TestGetWithEmptyRepository(t *testing.T) {
	if githubService == nil {
		t.Fatal("Service not set up properly.")
	}

	_, err := githubService.GetByOrigin("jbegley1995/")

	if err != RepositoryIncorrectFormatError {
		t.Errorf("Error actual = %v, and Expected = %v.", err, RepositoryIncorrectFormatError)
	}
}

func TestGetWithIncorrectOrganization(t *testing.T) {
	if githubService == nil {
		t.Fatal("Service not set up properly.")
	}

	repo, err := githubService.GetByOrigin("googleasdf/go-github")

	if repo != nil || err == nil {
		t.Errorf("Expected an error when trying to fetch for an organization that doesn't exist.")
	}
}

func TestGetWithIncorrectRepository(t *testing.T) {
	if githubService == nil {
		t.Fatal("Service not set up properly.")
	}

	repo, err := githubService.GetByOrigin("google/go-github-asdf")

	if repo != nil || err == nil {
		t.Errorf("Expected an error when trying to fetch for an organization that doesn't exist.")
	}
}
