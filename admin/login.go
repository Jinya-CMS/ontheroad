package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"go.jinya.de/ontheroad/database"
	httpUtils "go.jinya.de/ontheroad/utils/http"
	"net/http"
	"time"
)

type authData struct {
	Email    string
	Password string
}

func (data *authData) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&data.Email: binding.Field{
			Form:         "email",
			Required:     true,
			ErrorMessage: "The email is required",
		},
		&data.Password: binding.Field{
			Form:         "password",
			Required:     true,
			ErrorMessage: "The password is required",
		},
	}
}

func LoginView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpUtils.RenderSingle("templates/admin/auth/login.html.tmpl", nil, w)
}

func LoginAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.ContentLength == 0 {
		LoginView(w, r, nil)
		return
	}

	data := new(authData)
	errors := binding.Bind(r, data)
	var templateData struct {
		Error    string
		Email    string
		HasError bool
	}
	templateData.HasError = false

	if errors != nil {
		templateData.Error = errors.Error()
		templateData.HasError = true
		httpUtils.RenderSingle("templates/admin/auth/login.html.tmpl", templateData, w)
		return
	}

	user, err := database.ValidateEmailAndPassword(data.Email, data.Password)
	if err != nil {
		templateData.Error = err.Error()
		templateData.HasError = true
		httpUtils.RenderSingle("templates/admin/auth/login.html.tmpl", templateData, w)
		return
	}

	if user == nil {
		templateData.Error = "Wrong email or wrong password"
		templateData.HasError = true
		httpUtils.RenderSingle("templates/admin/auth/login.html.tmpl", templateData, w)
		return
	}

	token, err := database.CreateAuthenticationToken(user, r.RemoteAddr)
	if err != nil {
		templateData.Error = "Wrong email or wrong password"
		templateData.HasError = true
		httpUtils.RenderSingle("templates/admin/auth/login.html.tmpl", templateData, w)
		return
	}

	authCookie := http.Cookie{
		Name:       "auth",
		Value:      token,
		Path:       "/",
		Expires:    time.Now().AddDate(0, 0, 1),
		RawExpires: "",
		HttpOnly:   true,
	}

	http.SetCookie(w, &authCookie)
	http.Redirect(w, r, "/admin/user", http.StatusTemporaryRedirect)
}
