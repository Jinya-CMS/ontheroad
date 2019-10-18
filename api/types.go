package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/database"
	"go.jinya.de/ontheroad/youtrack"
	"net/http"
)

func GetAllTypesAction(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	projectId := params.ByName("id")
	project, err := database.GetProject(projectId)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	youtrackClient := youtrack.Client{Project: *project}
	types, err := youtrackClient.GetTypes()
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(types)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(result)
}
