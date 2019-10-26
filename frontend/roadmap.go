package frontend

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/database"
	httputils "go.jinya.de/ontheroad/utils/http"
	"go.jinya.de/ontheroad/youtrack"
	"net/http"
	"sort"
	"time"
)

func count(issues []youtrack.Issue, predicate func(issue youtrack.Issue) bool) int {
	count := 0
	for _, item := range issues {
		if predicate(item) {
			count += 1
		}
	}

	return count
}

type templateVersion struct {
	Name        string
	Issues      []youtrack.Issue
	Released    bool
	ReleaseDate time.Time
}

var templateData struct {
	Error          string
	Versions       []templateVersion
	CurrentProject *database.Project
	HasError       bool
	Projects       []database.Project
}

func RoadmapViewWithoutKey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	projects, err := database.GetAllProjects()
	if err != nil {
		templateData.Error = err.Error()
		templateData.HasError = true
		httputils.RenderSingle("templates/frontend/error/generic.html.tmpl", templateData, w)
		return
	}

	RoadmapView(w, r, httprouter.Params{httprouter.Param{
		Key:   "key",
		Value: projects[0].Key,
	}})
}

func RoadmapView(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	templateData.HasError = false
	projectKey := params.ByName("key")
	projects, err := database.GetAllProjects()
	if err != nil {
		templateData.Error = err.Error()
		templateData.HasError = true
		httputils.RenderSingle("templates/frontend/error/generic.html.tmpl", templateData, w)
		return
	}

	templateData.Projects = projects
	var project *database.Project
	if projectKey == "" {
		project = &projects[0]
	} else {
		project, err = database.GetProjectByKey(projectKey)
		if err != nil {
			templateData.Error = err.Error()
			templateData.HasError = true
			httputils.RenderSingle("templates/frontend/error/generic.html.tmpl", templateData, w)
			return
		}
	}

	templateData.CurrentProject = project
	client := youtrack.Client{Project: *project}
	versions, err := client.GetVersions()
	if err != nil {
		templateData.Error = err.Error()
		templateData.HasError = true
		httputils.RenderSingle("templates/frontend/error/generic.html.tmpl", templateData, w)
		return
	}

	issues, err := client.GetIssues([]string{}, []string{}, []string{"Feature", "Improvement"}, "Fix Versions", "DESC")
	if err != nil {
		templateData.Error = err.Error()
		templateData.HasError = true
		httputils.RenderSingle("templates/frontend/error/generic.html.tmpl", templateData, w)
		return
	}

	templateVersions := make([]templateVersion, len(versions))
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Ordinal < versions[j].Ordinal
	})

	for vIdx, version := range versions {
		issueCount := count(issues, func(issue youtrack.Issue) bool {
			return issue.FixVersion == version.Name
		})
		issueList := make([]youtrack.Issue, issueCount)

		idx := 0
		for _, issue := range issues {
			if issue.FixVersion == version.Name {
				issueList[idx] = issue
				idx++
			}
		}

		templateVersions[vIdx] = templateVersion{
			Name:        version.Name,
			Released:    version.Released,
			ReleaseDate: version.ReleaseDate,
			Issues:      issueList,
		}
	}

	templateData.Versions = templateVersions
	templateData.HasError = false
	httputils.RenderFrontend("templates/frontend/roadmap/index.html.tmpl", templateData, r, w)
}
