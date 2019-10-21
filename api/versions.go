package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/database"
	"go.jinya.de/ontheroad/youtrack"
	"net/http"
)

func GetAllVersionsAction(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	projectId := params.ByName("id")
	project, err := database.GetProject(projectId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	youtrackClient := youtrack.Client{Project: *project}
	versions, err := youtrackClient.GetVersions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(versions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(result)
}
