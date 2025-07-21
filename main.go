package main

import (
	"net/http"

	"github.com/bimjx/barangay/models"
	"github.com/bimjx/barangay/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Applications{},
	)
	uadmin.RootURL = "/admin/"

	// Routings
	http.HandleFunc("/home", uadmin.Handler(views.MainHandler))
	http.HandleFunc("/", uadmin.Handler(views.IndexHandler))
	uadmin.StartServer()
}
