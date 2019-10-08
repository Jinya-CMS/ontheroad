package admin

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/database"
	httpUtils "go.jinya.de/ontheroad/utils/http"
	"net/http"
	"strconv"
)

func ListUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type listData struct {
		Error         string
		Users         []database.User
		TotalCount    int
		CurrentOffset int
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	users, totalCount, err := database.GetUsers(offset, limit, r.URL.Query().Get("keyword"))
	if err != nil {
		tmplData := listData{
			Error:         err.Error(),
			Users:         nil,
			TotalCount:    0,
			CurrentOffset: 0,
		}

		httpUtils.RenderAdmin("admin/user/list.html.tmpl", tmplData, w)
		return
	}

	tmplData := listData{
		Error:         "",
		Users:         users,
		TotalCount:    totalCount,
		CurrentOffset: offset,
	}

	httpUtils.RenderAdmin("admin/user/list.html.tmpl", tmplData, w)
}
