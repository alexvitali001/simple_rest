package main

import (
	"go_simple_rest/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Route : The routes for the API
type Route struct {
	Router *mux.Router
	Path   string
	Func   func(http.ResponseWriter, *http.Request)
	Method string
}

var routes []Route

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	setupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server started and running at port", port)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router)))

}

// Setup REST routes
func setupRoutes(router *mux.Router) {
	v1Routes := router.PathPrefix("/v1").Subrouter()
	routes = append(routes, Route{Router: v1Routes, Path: "/version", Func: controllers.Version, Method: "GET"})
	routes = append(routes, Route{Router: v1Routes, Path: "/movie/list", Func: controllers.MovieList, Method: "GET"})

	for _, r := range routes {
		r.Router.HandleFunc(r.Path, r.Func).Methods(r.Method)
	}
}
