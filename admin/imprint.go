package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
	"github.com/mholt/binding"
	"go.jinya.de/ontheroad/database"
	httpUtils "go.jinya.de/ontheroad/utils/http"
	"net/http"
)

func EditImprintView(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	config, err := database.GetConfiguration("imprint")
	var tmplData map[string]interface{}
	if err != nil {
		tmplData = map[string]interface{}{
			"Error":    nil,
			"HasError": false,
			"Imprint":  "",
		}

		httpUtils.RenderAdmin("templates/admin/imprint/edit.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/imprint/edit.html.tmpl", map[string]interface{}{
		"Error":    nil,
		"HasError": false,
		"Imprint":  config.Value,
	}, w)
}

func EditImprintAction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.ContentLength == 0 {
		EditConfigView(w, r, params)
		return
	}

	var tmplData map[string]interface{}
	config := new(database.Configuration)
	validityErrors := binding.Bind(r, config)

	if validityErrors != nil {
		tmplData := map[string]interface{}{
			"Error":    validityErrors.Error(),
			"HasError": true,
			"Imprint":  "",
		}

		httpUtils.RenderAdmin("templates/admin/imprint/edit.html.tmpl", tmplData, w)
		return
	}

	err := database.SetConfiguration(config.Key, config.Value)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = map[string]interface{}{
				"Error":    pgErr.Detail,
				"HasError": true,
				"Imprint":  "",
			}
		} else {
			tmplData = map[string]interface{}{
				"Error":    pgErr.Detail,
				"HasError": true,
				"Imprint":  "",
			}
		}

		httpUtils.RenderAdmin("templates/admin/imprint/edit.html.tmpl", tmplData, w)
		return
	}

	processPostcssConfig()
	http.Redirect(w, r, "/admin/imprint", http.StatusSeeOther)
}
