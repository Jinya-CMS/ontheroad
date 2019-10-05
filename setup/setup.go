package setup

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"go.jinya.de/ontheroad/database"
	httpUtils "go.jinya.de/ontheroad/utils/http"
	"net/http"
	"os"
)

func checkIfSetup() bool {
	_, err := os.Stat("setup.lock")
	return os.IsNotExist(err)
}

func checkIfDatabaseSetup() bool {
	_, err := os.Stat("databaseSetup.lock")
	return os.IsNotExist(err)
}

func Welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !checkIfSetup() {
		http.Redirect(w, r, "/admin/login", http.StatusMovedPermanently)
		return
	}

	httpUtils.Render("templates/setup/index.html.tmpl", nil, w)
}

func DatabaseView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !checkIfSetup() {
		http.Redirect(w, r, "/admin/login", http.StatusMovedPermanently)
		return
	}

	if !checkIfDatabaseSetup() {
		http.Redirect(w, r, "/setup/admin", http.StatusTemporaryRedirect)
		return
	}

	connectionString := os.Getenv("connectionString")
	httpUtils.Render("templates/setup/database.html.tmpl", connectionString, w)
}

func DatabaseAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := database.CreateDatabase()
	if err != nil {
		httpUtils.Render("templates/setup/databaseError.html.tmpl", err, w)
		return
	}

	file, err := os.Create("databaseSetup.lock")
	if err != nil {
		httpUtils.Render("templates/setup/databaseError.html.tmpl", err, w)
		return
	}

	_ = file.Close()
	http.Redirect(w, r, "/setup/admin", http.StatusTemporaryRedirect)
}

func CreateAdminView(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !checkIfSetup() {
		http.Redirect(w, r, "/admin/login", http.StatusMovedPermanently)
		return
	}

	if checkIfDatabaseSetup() {
		http.Redirect(w, r, "/setup/database", http.StatusTemporaryRedirect)
		return
	}

	httpUtils.Render("templates/setup/admin.html.tmpl", nil, w)
}

func CreateAdminAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.ContentLength == 0 {
		CreateAdminView(w, r, nil)
		return
	}

	user := new(database.User)
	validityErrors := binding.Bind(r, user)
	type createAdminTmplData struct {
		Error string
	}

	if validityErrors != nil {
		tmplData := createAdminTmplData{
			Error: validityErrors.Error(),
		}

		httpUtils.Render("templates/setup/admin.html.tmpl", tmplData, w)
		return
	}

	err := database.CreateUser(user)
	if err != nil {
		tmplData := createAdminTmplData{
			Error: err.Error(),
		}

		httpUtils.Render("templates/setup/admin.html.tmpl", tmplData, w)
		return
	}

	file, err := os.Create("setup.lock")
	_ = file.Close()

	http.Redirect(w, r, "/admin/login", http.StatusTemporaryRedirect)
}
