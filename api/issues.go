package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/database"
	"go.jinya.de/ontheroad/youtrack"
	"net/http"
)

func GetIssuesAction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	versionsFromQuery := r.URL.Query().Get("versions")
	if versionsFromQuery == "" {
		versionsFromQuery = "[]"
	}

	typesFromQuery := r.URL.Query().Get("types")
	if typesFromQuery == "" {
		typesFromQuery = "[]"
	}

	subsystemsFromQuery := r.URL.Query().Get("subsystems")
	if subsystemsFromQuery == "" {
		subsystemsFromQuery = "[]"
	}

	projectId := params.ByName("id")
	orderColumn := r.URL.Query().Get("orderBy")
	orderDirection := r.URL.Query().Get("orderDirection")

	project, err := database.GetProject(projectId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	var versions []string
	var types []string
	var subsystems []string

	err = json.Unmarshal([]byte(versionsFromQuery), &versions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal([]byte(typesFromQuery), &types)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal([]byte(subsystemsFromQuery), &subsystems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	client := youtrack.Client{Project: *project}
	issues, err := client.GetIssues(versions, subsystems, types, orderColumn, orderDirection)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(issues)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(result)
}
