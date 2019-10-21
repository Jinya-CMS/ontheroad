package api

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/dummy_data"
	"go.jinya.de/ontheroad/utils/tests"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("Setting up the tests")
	dummy_data.FillProjects()
	defer dummy_data.ClearProjects()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGetAllSubsystemsAction(t *testing.T) {
	req := httptest.NewRequest("GET", "/project", nil)
	res := tests.NewResponse(t, 200, "")
	GetAllSubsystemsAction(res, req, httprouter.Params{
		httprouter.Param{
			Key:   "id",
			Value: dummy_data.TestProject.Id,
		},
	})
}
