package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
	"github.com/mholt/binding"
	"go.jinya.de/ontheroad/database"
	httpUtils "go.jinya.de/ontheroad/utils/http"
	"net/http"
	"strconv"
)

type projectActionTmplData struct {
	Error    string
	Project  *database.Project
	HasError bool
}

func ListProject(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type listData struct {
		Error           string
		HasError        bool
		Projects        []database.Project
		TotalCount      int
		CurrentOffset   int
		Empty           bool
		Keyword         string
		NextEnabled     bool
		PreviousEnabled bool
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	keyword := r.URL.Query().Get("keyword")

	projects, totalCount, err := database.GetProjects(offset, limit, keyword)
	if err != nil {
		tmplData := listData{
			Error:           err.Error(),
			Projects:        nil,
			TotalCount:      0,
			CurrentOffset:   0,
			Empty:           true,
			HasError:        true,
			Keyword:         keyword,
			NextEnabled:     false,
			PreviousEnabled: false,
		}

		httpUtils.RenderAdmin("templates/admin/project/list.html.tmpl", tmplData, w)
		return
	}

	tmplData := listData{
		Error:           "",
		Projects:        projects,
		TotalCount:      totalCount,
		CurrentOffset:   offset,
		Empty:           totalCount == 0,
		HasError:        false,
		Keyword:         keyword,
		NextEnabled:     totalCount >= limit+offset,
		PreviousEnabled: offset > 0,
	}

	httpUtils.RenderAdmin("templates/admin/project/list.html.tmpl", tmplData, w)
}

func AddProjectView(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	httpUtils.RenderAdmin("templates/admin/project/add.html.tmpl", nil, w)
}

func AddProjectAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.ContentLength == 0 {
		AddProjectView(w, r, nil)
		return
	}

	project := new(database.Project)
	validityErrors := binding.Bind(r, project)

	if validityErrors != nil {
		tmplData := projectActionTmplData{
			Error:    validityErrors.Error(),
			HasError: true,
			Project:  project,
		}

		httpUtils.RenderAdmin("templates/admin/project/add.html.tmpl", tmplData, w)
		return
	}

	err := database.CreateProject(project)
	if err != nil {
		var tmplData projectActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = projectActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Project:  project,
			}
		} else {
			tmplData = projectActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Project:  project,
			}
		}

		httpUtils.RenderAdmin("templates/admin/project/add.html.tmpl", tmplData, w)
		return
	}

	http.Redirect(w, r, "/admin/project", http.StatusSeeOther)
}

func DetailsProjectView(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	project, err := database.GetProject(params.ByName("id"))
	var tmplData projectActionTmplData
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = projectActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Project:  project,
			}
		} else {
			tmplData = projectActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Project:  project,
			}
		}

		httpUtils.RenderAdmin("templates/admin/project/details.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/project/details.html.tmpl", projectActionTmplData{
		Error:    "",
		Project:  project,
		HasError: false,
	}, w)
}

func EditProjectView(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	project, err := database.GetProject(params.ByName("id"))
	var tmplData projectActionTmplData
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = projectActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Project:  project,
			}
		} else {
			tmplData = projectActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Project:  project,
			}
		}

		httpUtils.RenderAdmin("templates/admin/project/edit.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/project/edit.html.tmpl", projectActionTmplData{
		Error:    "",
		Project:  project,
		HasError: false,
	}, w)
}

func EditProjectAction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.ContentLength == 0 {
		EditProjectView(w, r, params)
		return
	}

	project, err := database.GetProject(params.ByName("id"))
	if err != nil {
		var tmplData projectActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = projectActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Project:  project,
			}
		} else {
			tmplData = projectActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Project:  project,
			}
		}

		httpUtils.RenderAdmin("templates/admin/project/edit.html.tmpl", tmplData, w)
		return
	}

	validityErrors := binding.Bind(r, project)

	if validityErrors != nil {
		tmplData := projectActionTmplData{
			Error:    validityErrors.Error(),
			HasError: true,
			Project:  project,
		}

		httpUtils.RenderAdmin("templates/admin/project/edit.html.tmpl", tmplData, w)
		return
	}

	err = database.UpdateProject(project)
	if err != nil {
		var tmplData projectActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = projectActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Project:  project,
			}
		} else {
			tmplData = projectActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Project:  project,
			}
		}

		httpUtils.RenderAdmin("templates/admin/project/edit.html.tmpl", tmplData, w)
		return
	}

	http.Redirect(w, r, "/admin/project", http.StatusSeeOther)
}

func DeleteProjectView(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	project, err := database.GetProject(params.ByName("id"))
	var tmplData projectActionTmplData
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = projectActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Project:  project,
			}
		} else {
			tmplData = projectActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Project:  project,
			}
		}

		httpUtils.RenderAdmin("templates/admin/project/delete.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/project/delete.html.tmpl", projectActionTmplData{
		Error:    "",
		Project:  project,
		HasError: false,
	}, w)
}

func DeleteProjectAction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := database.DeleteProject(params.ByName("id"))
	if err != nil {
		project, _ := database.GetProject(params.ByName("id"))
		var tmplData projectActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = projectActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Project:  project,
			}
		} else {
			tmplData = projectActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Project:  project,
			}
		}

		httpUtils.RenderAdmin("templates/admin/project/delete.html.tmpl", tmplData, w)
		return
	}

	http.Redirect(w, r, "/admin/project", http.StatusSeeOther)
}
