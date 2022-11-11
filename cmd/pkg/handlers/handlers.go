package handlers

import (
	"net/http"

	"github.com/vikashparashar/mywebsite/cmd/pkg/rander"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// rander.Rander_Template(w, "home.page.tmpl")
	rander.Rander_Template(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	// rander.Rander_Template(w, "about.page.tmpl")
	rander.Rander_Template(w, "about.page.html")
}
