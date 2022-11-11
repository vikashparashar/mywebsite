package handlers

import (
	"net/http"

	"github.com/vikashparashar/mywebsite/cmd/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl") // using .tmpl extension
	// render.Render_Template(w, "home.page.html") // using .html extension
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl") // using .tmpl extension
	// render.Render_Template(w, "about.page.html") // using .html extension
}
