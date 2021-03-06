package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
	"github.com/mholt/binding"
	"go.jinya.de/ontheroad/database"
	httpUtils "go.jinya.de/ontheroad/utils/http"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

type configActionTmplData struct {
	Error    string
	Config   *database.Configuration
	HasError bool
}

func processPostcssConfig() {
	tmpl, err := template.ParseFiles("templates/theme/tailwind.frontend.config.jstmpl")
	if err != nil {
		return
	}

	primaryColor, err := database.GetConfiguration("PrimaryColor")
	if err != nil {
		primaryColor = &database.Configuration{
			Id:    "",
			Key:   "",
			Value: "#da9c8a",
		}
	}

	secondaryColor, err := database.GetConfiguration("SecondaryColor")
	if err != nil {
		secondaryColor = &database.Configuration{
			Id:    "",
			Key:   "",
			Value: "#8ac8da",
		}
	}

	grayColor, err := database.GetConfiguration("GrayColor")
	if err != nil {
		grayColor = &database.Configuration{
			Id:    "",
			Key:   "",
			Value: "#adb5bd",
		}
	}

	data := map[string]string{
		"PrimaryColor":   primaryColor.Value,
		"SecondaryColor": secondaryColor.Value,
		"GrayColor":      grayColor.Value,
	}

	targetFile, err := os.OpenFile("theme/tailwind.frontend.config.js", os.O_RDWR|os.O_CREATE, 0755)
	err = tmpl.Execute(targetFile, data)
	if err != nil {
		return
	}

	cmd := exec.Command("yarn", "gulp")
	cmd.Dir = "theme/"
	_ = cmd.Run()
}

func ListConfig(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type listData struct {
		Error    string
		HasError bool
		Configs  []database.Configuration
	}

	configs, err := database.GetConfigurations()
	if err != nil {
		tmplData := listData{
			Error:    err.Error(),
			Configs:  nil,
			HasError: true,
		}

		httpUtils.RenderAdmin("templates/admin/config/list.html.tmpl", tmplData, w)
		return
	}

	tmplData := listData{
		Error:    "",
		Configs:  configs,
		HasError: false,
	}

	httpUtils.RenderAdmin("templates/admin/config/list.html.tmpl", tmplData, w)
}

func AddConfigView(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	httpUtils.RenderAdmin("templates/admin/config/add.html.tmpl", nil, w)
}

func AddConfigAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.ContentLength == 0 {
		AddConfigView(w, r, nil)
		return
	}

	config := new(database.Configuration)
	validityErrors := binding.Bind(r, config)

	if validityErrors != nil {
		tmplData := configActionTmplData{
			Error:    validityErrors.Error(),
			HasError: true,
			Config:   config,
		}

		httpUtils.RenderAdmin("templates/admin/config/add.html.tmpl", tmplData, w)
		return
	}

	err := database.SetConfiguration(config.Key, config.Value)
	if err != nil {
		var tmplData configActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = configActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Config:   config,
			}
		} else {
			tmplData = configActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Config:   config,
			}
		}

		httpUtils.RenderAdmin("templates/admin/config/add.html.tmpl", tmplData, w)
		return
	}

	processPostcssConfig()
	http.Redirect(w, r, "/admin/config", http.StatusSeeOther)
}

func EditConfigView(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	config, err := database.GetConfiguration(params.ByName("key"))
	var tmplData configActionTmplData
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = configActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Config:   config,
			}
		} else {
			tmplData = configActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Config:   config,
			}
		}

		httpUtils.RenderAdmin("templates/admin/config/edit.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/config/edit.html.tmpl", configActionTmplData{
		Error:    "",
		Config:   config,
		HasError: false,
	}, w)
}

func EditConfigAction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.ContentLength == 0 {
		EditConfigView(w, r, params)
		return
	}

	config, err := database.GetConfiguration(params.ByName("key"))
	if err != nil {
		var tmplData configActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = configActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Config:   config,
			}
		} else {
			tmplData = configActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Config:   config,
			}
		}

		httpUtils.RenderAdmin("templates/admin/config/edit.html.tmpl", tmplData, w)
		return
	}

	validityErrors := binding.Bind(r, config)

	if validityErrors != nil {
		tmplData := configActionTmplData{
			Error:    validityErrors.Error(),
			HasError: true,
			Config:   config,
		}

		httpUtils.RenderAdmin("templates/admin/config/edit.html.tmpl", tmplData, w)
		return
	}

	err = database.SetConfiguration(config.Key, config.Value)
	if err != nil {
		var tmplData configActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = configActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Config:   config,
			}
		} else {
			tmplData = configActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Config:   config,
			}
		}

		httpUtils.RenderAdmin("templates/admin/config/edit.html.tmpl", tmplData, w)
		return
	}

	processPostcssConfig()
	http.Redirect(w, r, "/admin/config", http.StatusSeeOther)
}

func DeleteConfigView(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	config, err := database.GetConfiguration(params.ByName("key"))
	var tmplData configActionTmplData
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = configActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Config:   config,
			}
		} else {
			tmplData = configActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Config:   config,
			}
		}

		httpUtils.RenderAdmin("templates/admin/config/delete.html.tmpl", tmplData, w)
		return
	}

	httpUtils.RenderAdmin("templates/admin/config/delete.html.tmpl", configActionTmplData{
		Error:    "",
		Config:   config,
		HasError: false,
	}, w)
}

func DeleteConfigAction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := database.DeleteConfiguration(params.ByName("key"))
	if err != nil {
		config, _ := database.GetConfiguration(params.ByName("key"))
		var tmplData configActionTmplData
		if pgErr, ok := err.(*pq.Error); ok {
			tmplData = configActionTmplData{
				Error:    pgErr.Detail,
				HasError: true,
				Config:   config,
			}
		} else {
			tmplData = configActionTmplData{
				Error:    err.Error(),
				HasError: true,
				Config:   config,
			}
		}

		httpUtils.RenderAdmin("templates/admin/config/delete.html.tmpl", tmplData, w)
		return
	}

	processPostcssConfig()
	http.Redirect(w, r, "/admin/config", http.StatusSeeOther)
}
