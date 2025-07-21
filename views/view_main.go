package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Render the main view
	page := strings.TrimPrefix(r.URL.Path, "/")
	page = strings.TrimSuffix(page, "/")
	c := map[string]interface{}{}

	uadmin.Trail(uadmin.DEBUG, "page: %v", page)
	session := uadmin.IsAuthenticated(r)

	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// if page == "" {
	// 	page = "get"
	// 	http.Redirect(w, r, "/home", http.StatusSeeOther)
	// }
	switch page {
	case "index":
		c = HomeHandler(w, r)
		page = "get"
		fmt.Println("MainHandler called for index")
	case "officials":
		c = OfficialHandlers(w, r)
		page = "officials"
	case "contacts":
		c = ContactsHandler(w, r)
		page = "contacts"
	case "get":
		if session != nil {
			c = ApplicationHandler(w, r)
			page = "get"
		}
	default:
		uadmin.Trail(uadmin.DEBUG, "Unknown page: %s", page)
		// c = IndexHandler(w, r)
		c = HomeHandler(w, r)
		page = "get"
	}

	c["Page"] = page

	uadmin.Trail(uadmin.DEBUG, page)
	uadmin.RenderHTML(w, r, fmt.Sprintf("./templates/%v.html", page), c)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}

	return c
}
