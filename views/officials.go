package views

import (
	"net/http"
)

func OfficialHandlers(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}
	// Render the official view
	// uadmin.RenderHTML(w, r, "./templates/officials.html", c)
	return c
}
