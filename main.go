package main

import (
	"go.jinya.de/ontheroad/routing"
	"log"
	"net/http"
)

func main() {
	router := routing.GetHttpRouter()
	router.ServeFiles("/public/*filepath", http.Dir("theme/build/"))

	log.Print("Starting ontheroad server on port 9000")
	err := http.ListenAndServe(":9000", router)
	if err != nil {
		log.Fatalf("Failed to start ontheroad server. %s", err.Error())
	}
}
