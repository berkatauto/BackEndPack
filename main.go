package main

import (
	"net/http"

	controller "github.com/berkatauto/servs-kdrn/controller/authenticator"
	"github.com/berkatauto/servs-kdrn/login"
)

func main() {
	controller.Auth()
	login.RegisterLoginRoutes()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "pages/login.html")
		} else {
			http.Error(w, "Method not authorized.", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Handle the dashboard logic here and serve the dashboard HTML
			http.ServeFile(w, r, "pages/main.html")
		} else {
			http.Error(w, "Method not authorized.", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8000", nil)
}
