package routing

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/admin"
	"go.jinya.de/ontheroad/setup"
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

	router.GET("/admin/user", AuthenticatedMiddleware(admin.ListUser))
	router.GET("/admin/user/add", AuthenticatedMiddleware(admin.AddUserView))
	router.POST("/admin/user/add", AuthenticatedMiddleware(admin.AddUserAction))
	router.GET("/admin/user/edit/:id", AuthenticatedMiddleware(admin.EditUserView))
	router.POST("/admin/user/edit/:id", AuthenticatedMiddleware(admin.EditUserAction))

	return router
}
