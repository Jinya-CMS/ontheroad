package routing

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/database"
	httpUtils "go.jinya.de/ontheroad/utils/http"
	"net/http"
)

func AuthenticatedMiddleware(handler func(w http.ResponseWriter, r *http.Request, params httprouter.Params)) func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		authCookie, err := r.Cookie("auth")
		if err != nil || authCookie == nil {
			w.WriteHeader(http.StatusUnauthorized)
			httpUtils.RenderSingle("templates/admin/error/401.html.tmpl", nil, w)
			return
		}

		if database.ValidateAuthenticationToken(authCookie.Value) {
			handler(w, r, params)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		httpUtils.RenderSingle("templates/admin/error/401.html.tmpl", nil, w)
	}
}
