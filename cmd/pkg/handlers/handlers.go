package handlers

import (
	"net/http"

	"github.com/vikashparashar/mywebsite/cmd/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// rander.Rander_Template(w, "home.page.tmpl")
	render.Render_Template(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	// rander.Rander_Template(w, "about.page.tmpl")
	render.Render_Template(w, "about.page.html")
}
