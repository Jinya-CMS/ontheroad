package routing

import (
	"github.com/julienschmidt/httprouter"
	"go.jinya.de/ontheroad/setup"
)

func GetHttpRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/setup", setup.Welcome)

	router.GET("/setup/database", setup.DatabaseView)
	router.POST("/setup/database", setup.DatabaseAction)

	router.GET("/setup/admin", setup.CreateAdminView)
	router.POST("/setup/admin", setup.CreateAdminAction)

	return router
}
