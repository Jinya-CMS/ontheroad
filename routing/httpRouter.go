package routing

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/admin"
	"go.jinya.de/ontheroad/api"
	"go.jinya.de/ontheroad/frontend"
	"go.jinya.de/ontheroad/setup"
	"net/http"
)

func GetHttpRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/setup", setup.Welcome)

	router.GET("/setup/database", setup.DatabaseView)
	router.POST("/setup/database", setup.DatabaseAction)

	router.GET("/setup/admin", setup.CreateAdminView)
	router.POST("/setup/admin", setup.CreateAdminAction)

	router.GET("/admin/login", admin.LoginView)
	router.POST("/admin/login", admin.LoginAction)
	router.POST("/admin/logout", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		authCookie, err := r.Cookie("auth")
		if err != nil {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		authCookie.MaxAge = 0
		http.SetCookie(w, authCookie)
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	})

	router.GET("/admin/user", AuthenticatedMiddleware(admin.ListUser))
	router.GET("/admin/user/add", AuthenticatedMiddleware(admin.AddUserView))
	router.POST("/admin/user/add", AuthenticatedMiddleware(admin.AddUserAction))
	router.GET("/admin/user/edit/:id", AuthenticatedMiddleware(admin.EditUserView))
	router.POST("/admin/user/edit/:id", AuthenticatedMiddleware(admin.EditUserAction))
	router.GET("/admin/user/delete/:id", AuthenticatedMiddleware(admin.DeleteUserView))
	router.POST("/admin/user/delete/:id", AuthenticatedMiddleware(admin.DeleteUserAction))

	router.GET("/admin/project", AuthenticatedMiddleware(admin.ListProject))
	router.GET("/admin/project/add", AuthenticatedMiddleware(admin.AddProjectView))
	router.POST("/admin/project/add", AuthenticatedMiddleware(admin.AddProjectAction))
	router.GET("/admin/project/details/:id", AuthenticatedMiddleware(admin.DetailsProjectView))
	router.GET("/admin/project/edit/:id", AuthenticatedMiddleware(admin.EditProjectView))
	router.POST("/admin/project/edit/:id", AuthenticatedMiddleware(admin.EditProjectAction))
	router.GET("/admin/project/delete/:id", AuthenticatedMiddleware(admin.DeleteProjectView))
	router.POST("/admin/project/delete/:id", AuthenticatedMiddleware(admin.DeleteProjectAction))

	router.GET("/api/:id/version", api.GetAllVersionsAction)
	router.GET("/api/:id/subsystem", api.GetAllSubsystemsAction)
	router.GET("/api/:id/types", api.GetAllTypesAction)
	router.GET("/api/:id/issues", api.GetIssuesAction)

	router.GET("/", frontend.RoadmapViewWithoutKey)
	router.GET("/roadmap", frontend.RoadmapViewWithoutKey)
	router.GET("/roadmap/:key", frontend.RoadmapView)

	return router
}
