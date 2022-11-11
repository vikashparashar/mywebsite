/*

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vikashparashar/mywebsite/cmd/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("")
	fmt.Println("  >>------>  Starting The Application  <------<<  ")
	fmt.Println("")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

*/

package main

import (
	"fmt"
	"github.com/vikashparashar/mywebsite/cmd/pkg/config"
	"github.com/vikashparashar/mywebsite/cmd/pkg/handlers"
	"github.com/vikashparashar/mywebsite/cmd/pkg/render"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreatingTemplateCache()

	if err != nil {
		log.Println("can not create template cache")

	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("  >>------>  Starting The Application  <------<<  ")
	fmt.Println("_____________      Port : 8081      ______________")
	fmt.Println("_____________    Host : LocalHost   ______________")
	fmt.Println("____________        Path : /         _____________")
	fmt.Println("___________      Path : /about       ____________")

	http.ListenAndServe(":8081", nil)
}
