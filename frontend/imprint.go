package frontend

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/database"
	httpUtils "go.jinya.de/ontheroad/utils/http"
	"net/http"
)

func Imprint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	imprint, _ := database.GetConfiguration("imprint")
	projects, _ := database.GetAllProjects()
	httpUtils.RenderFrontend("templates/frontend/imprint/index.html.tmpl", map[string]interface{}{
		"Imprint":  imprint.Value,
		"Projects": projects,
	}, r, w)
}
