package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new mux router
	r := mux.NewRouter()

	// Serve static files generated by Angular
	r.PathPrefix("/build/").Handler(http.StripPrefix("/build", http.FileServer(http.Dir("./dist"))))
	// Catch-all route to serve the Angular app for any request not matching a static file
	// r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./dist/index.html")
	// })

	// Catch-all route to serve the Angular app for any request not matching a static file

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("not found handler")
		http.ServeFile(w, r, "./dist/index.html")
	})

	log.Println("http://localhost:8080")
	// Start the server on port 8080
	http.ListenAndServe(":8080", r)
}