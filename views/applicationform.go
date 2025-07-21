package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func ApplicationHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}

	// Check if the user is authenticated
	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return c
	}

	// Render the application form view
	// uadmin.RenderHTML(w, r, "./templates/get.html", c)
	return c
}
