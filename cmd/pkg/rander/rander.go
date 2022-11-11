package rander

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Rander_Template(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Fatalln(err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error while parsing template : ", err)
		fmt.Fprintln(w, fmt.Sprintf("error while parsing template : %s\n", err))
	}

}
