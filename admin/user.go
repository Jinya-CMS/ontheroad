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

type userActionTmplData struct {
	Error    string
	User     *database.User
	HasError bool
}

func ListUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type listData struct {
		Error           string
		HasError        bool
		Users           []database.User
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

	users, totalCount, err := database.GetUsers(offset, limit, keyword)
	if err != nil {
		tmplData := listData{
			Error:           err.Error(),
			Users:           nil,
			TotalCount:      0,
			CurrentOffset:   0,
			Empty:           true,
			HasError:        true,
			Keyword:         keyword,
			NextEnabled:     false,
			PreviousEnabled: false,
		}

		httpUtils.RenderAdmin("templates/admin/user/list.html.tmpl", tmplData, w)
		return
	}

	tmplData := listData{
		Error:           "",
		Users:           users,
		TotalCount:      totalCount,
		CurrentOffset:   offset,
		Empty:           totalCount == 0,
		HasError:        false,
		Keyword:         keyword,
		NextEnabled:     totalCount >= limit+offset,
		PreviousEnabled: offset > 0,
	}

	httpUtils.RenderAdmin("templates/admin/user/list.html.tmpl", tmplData, w)
}

func AddUserView(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	httpUtils.RenderAdmin("templates/admin/user/add.html.tmpl", nil, w)
}

func AddUserAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.ContentLength == 0 {
		AddUserView(w, r, nil)
		return
	}

	user := new(database.User)
	validityErrors := binding.Bind(r, user)

	if validityErrors != nil {
		tmplData := userActionTmplData{
			Error:    validityErrors.Error(),
			HasError: true,
			User:     user,
		}

		httpUtils.RenderAdmin("templates/admin/user/add.html.tmpl", tmplData, w)
		return
	}

	err := database.CreateUser(user)
	if err != nil {
		var tmplData userActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = userActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = userActionTmplData{
				Error:    err.Error(),
				HasError: true,
				User:     user,
			}
		}

		httpUtils.RenderAdmin("templates/admin/user/add.html.tmpl", tmplData, w)
		return
	}

	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}

func EditUserView(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	user, err := database.GetUser(params.ByName("id"))
	var tmplData userActionTmplData
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = userActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = userActionTmplData{
				Error:    err.Error(),
				HasError: true,
				User:     user,
			}
		}

		httpUtils.RenderAdmin("templates/admin/user/edit.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/user/edit.html.tmpl", userActionTmplData{
		Error:    "",
		User:     user,
		HasError: false,
	}, w)
}

func EditUserAction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.ContentLength == 0 {
		EditUserView(w, r, params)
		return
	}

	user, err := database.GetUser(params.ByName("id"))
	if err != nil {
		var tmplData userActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = userActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = userActionTmplData{
				Error:    err.Error(),
				HasError: true,
				User:     user,
			}
		}

		httpUtils.RenderAdmin("templates/admin/user/edit.html.tmpl", tmplData, w)
		return
	}

	validityErrors := binding.Bind(r, user)

	if validityErrors != nil {
		tmplData := userActionTmplData{
			Error:    validityErrors.Error(),
			HasError: true,
			User:     user,
		}

		httpUtils.RenderAdmin("templates/admin/user/edit.html.tmpl", tmplData, w)
		return
	}

	err = database.UpdateUser(user)
	if err != nil {
		var tmplData userActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = userActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = userActionTmplData{
				Error:    err.Error(),
				HasError: true,
				User:     user,
			}
		}

		httpUtils.RenderAdmin("templates/admin/user/edit.html.tmpl", tmplData, w)
		return
	}

	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}

func DeleteUserView(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	user, err := database.GetUser(params.ByName("id"))
	var tmplData userActionTmplData
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = userActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = userActionTmplData{
				Error:    err.Error(),
				HasError: true,
				User:     user,
			}
		}

		httpUtils.RenderAdmin("templates/admin/user/delete.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/user/delete.html.tmpl", userActionTmplData{
		Error:    "",
		User:     user,
		HasError: false,
	}, w)
}

func DeleteUserAction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := database.DeleteUser(params.ByName("id"))
	if err != nil {
		user, _ := database.GetUser(params.ByName("id"))
		var tmplData userActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = userActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = userActionTmplData{
				Error:    err.Error(),
				HasError: true,
				User:     user,
			}
		}

		httpUtils.RenderAdmin("templates/admin/user/delete.html.tmpl", tmplData, w)
		return
	}

	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}
