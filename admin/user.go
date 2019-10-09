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
	type createUserTmplData struct {
		Error    string
		HasError bool
		User     *database.User
	}

	if validityErrors != nil {
		tmplData := createUserTmplData{
			Error:    validityErrors.Error(),
			HasError: true,
			User:     user,
		}

		httpUtils.RenderAdmin("templates/admin/user/add.html.tmpl", tmplData, w)
		return
	}

	err := database.CreateUser(user)
	if err != nil {
		var tmplData createUserTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = createUserTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = createUserTmplData{
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
	type editUserTmplData struct {
		Error    string
		User     *database.User
		HasError bool
	}

	user, err := database.GetUser(params.ByName("id"))
	var tmplData editUserTmplData
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = editUserTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = editUserTmplData{
				Error:    err.Error(),
				HasError: true,
				User:     user,
			}
		}

		httpUtils.RenderAdmin("templates/admin/user/edit.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/user/edit.html.tmpl", editUserTmplData{
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
	type editUserTmplData struct {
		Error    string
		HasError bool
		User     *database.User
	}

	user, err := database.GetUser(params.ByName("id"))
	if err != nil {
		var tmplData editUserTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = editUserTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = editUserTmplData{
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
		tmplData := editUserTmplData{
			Error:    validityErrors.Error(),
			HasError: true,
			User:     user,
		}

		httpUtils.RenderAdmin("templates/admin/user/edit.html.tmpl", tmplData, w)
		return
	}

	err = database.UpdateUser(user)
	if err != nil {
		var tmplData editUserTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = editUserTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				User:     user,
			}
		} else {
			tmplData = editUserTmplData{
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
