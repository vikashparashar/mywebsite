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
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vikashparashar/mywebsite/cmd/pkg/handlers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.Home)
	r.Get("/about", handlers.About)
	http.ListenAndServe(":8080", r)
}
