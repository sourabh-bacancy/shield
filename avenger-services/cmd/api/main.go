package main

import (
	"avenger/cmd/api/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var port string

func main() {
	port = ":5000"
	// router := httprouter.New()
	router := mux.NewRouter()
	// router.RedirectTrailingSlash = true
	handler.AddRouteHandlers(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "Authorization"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	fmt.Println("listening on port " + port)
	log.Fatal(http.ListenAndServe(port, c.Handler(router)))
}
