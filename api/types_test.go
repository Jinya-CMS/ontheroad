package api

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/dummy_data"
	"go.jinya.de/ontheroad/utils/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllTypesAction(t *testing.T) {
	req := httptest.NewRequest("GET", "/project", nil)
	res := tests.NewResponse(t, http.StatusOK, "")
	GetAllTypesAction(res, req, httprouter.Params{
		httprouter.Param{
			Key:   "id",
			Value: dummy_data.TestProject.Id,
		},
	})
}
