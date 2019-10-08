package main

import (
	"go.jinya.de/ontheroad/database/migrations"
	"go.jinya.de/ontheroad/routing"
	"log"
	"net/http"
	"os"
)

func contains(slice []string, search string) bool {
	for _, item := range slice {
		if item == search {
			return true
		}
	}

	return false
}

func main() {
	if contains(os.Args, "migrate") {
		err := migrations.Migrate()
		if err != nil {
			panic(err)
		}
	}

	if contains(os.Args, "start") {
		router := routing.GetHttpRouter()
		router.ServeFiles("/public/*filepath", http.Dir("theme/build/"))

		log.Print("Starting ontheroad server on port 9000")
		err := http.ListenAndServe(":9000", router)
		if err != nil {
			log.Fatalf("Failed to start ontheroad server. %s", err.Error())
		}
	}
}
