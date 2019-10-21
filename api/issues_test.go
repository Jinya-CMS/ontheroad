package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/dummy_data"
	"go.jinya.de/ontheroad/utils/tests"
	"go.jinya.de/ontheroad/youtrack"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetIssuesAction(t *testing.T) {
	request := httptest.NewRequest("GET", "/issues", nil)
	response := tests.NewResponseWithHandleBody(t, http.StatusOK, func(data []byte) {
		var issues []youtrack.Issue
		err := json.Unmarshal(data, &issues)
		if err != nil {
			t.Error(err)
		}

		client := youtrack.Client{Project: dummy_data.TestProject}
		issuesFromYouTrack, err := client.GetIssues([]string{}, []string{}, []string{}, "", "")
		issueCount := len(issues)
		if issueCount != len(issuesFromYouTrack) {
			t.Error()
		}
	})

	GetIssuesAction(response, request, httprouter.Params{httprouter.Param{
		Key:   "id",
		Value: dummy_data.TestProject.Id,
	}})
}

func TestGetIssuesActionByVersion(t *testing.T) {
	client := youtrack.Client{Project: dummy_data.TestProject}
	versions, err := client.GetVersions()
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest("GET", fmt.Sprintf("/issues?versions=[\"%s\"]", url.QueryEscape(versions[0].Name)), nil)
	response := tests.NewResponseWithHandleBody(t, http.StatusOK, func(data []byte) {
		var issues []youtrack.Issue
		err := json.Unmarshal(data, &issues)
		if err != nil {
			t.Error(err)
		}

		for _, issue := range issues {
			if issue.FixVersion != versions[0].Name {
				t.Error()
			}
		}
	})

	GetIssuesAction(response, request, httprouter.Params{httprouter.Param{
		Key:   "id",
		Value: dummy_data.TestProject.Id,
	}})
}

func TestGetIssuesActionBySubsystem(t *testing.T) {
	client := youtrack.Client{Project: dummy_data.TestProject}
	subsystems, err := client.GetSubsystems()
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest("GET", fmt.Sprintf("/issues?subsystems=[\"%s\"]", url.QueryEscape(subsystems[0].Name)), nil)
	response := tests.NewResponseWithHandleBody(t, http.StatusOK, func(data []byte) {
		var issues []youtrack.Issue
		err := json.Unmarshal(data, &issues)
		if err != nil {
			t.Error(err)
		}

		contains := func(array []string, key string) bool {
			for _, item := range array {
				if key == item {
					return true
				}
			}

			return false
		}

		for _, issue := range issues {
			if !contains(issue.Subsystems, subsystems[0].Name) {
				t.Error()
			}
		}
	})

	GetIssuesAction(response, request, httprouter.Params{httprouter.Param{
		Key:   "id",
		Value: dummy_data.TestProject.Id,
	}})
}

func TestGetIssuesActionByTypes(t *testing.T) {
	client := youtrack.Client{Project: dummy_data.TestProject}
	types, err := client.GetTypes()
	if err != nil {
		t.Error(err)
	}

	request := httptest.NewRequest("GET", fmt.Sprintf("/issues?types=[\"%s\"]", url.QueryEscape(types[0].Name)), nil)
	response := tests.NewResponseWithHandleBody(t, http.StatusOK, func(data []byte) {
		var issues []youtrack.Issue
		err := json.Unmarshal(data, &issues)
		if err != nil {
			t.Error(err)
		}

		for _, issue := range issues {
			if issue.Type != types[0].Name {
				t.Error()
			}
		}
	})

	GetIssuesAction(response, request, httprouter.Params{httprouter.Param{
		Key:   "id",
		Value: dummy_data.TestProject.Id,
	}})
}
