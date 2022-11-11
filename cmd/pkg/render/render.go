package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tc = make(map[string]*template.Template)

/*
	func Render_Template(w http.ResponseWriter, tmpl string) {
		parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl") // using .tmpl extension
		// parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")  // // using .html extension
		if err != nil {
			log.Fatalln(err)
		}
		err = parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error while parsing template : ", err)
			fmt.Fprintln(w, fmt.Sprintf("error while parsing template : %s\n", err))
		}

}
*/
func Render_Template_Test(w http.ResponseWriter, tmpl string) {
	_, ok := tc[tmpl]
	if !ok {
		parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl") // using .tmpl extension
		// parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")  // // using .html extension
		if err != nil {
			log.Fatalln(err)
		}
		err = parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error while parsing template : ", err)
			fmt.Fprintln(w, fmt.Sprintf("error while parsing template : %s\n", err))
		}
		tc[tmpl] = parsedTemplate
	} else {
		fmt.Println("using cache template")
	}
}
