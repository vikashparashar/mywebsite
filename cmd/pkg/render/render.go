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
func Render_Template(w http.ResponseWriter, tmpl string) {
	_, available := tc[tmpl]
	if !available {
		fmt.Println("temp not available but we are now creating a new template to use")
		parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl") // using .tmpl extension
		// parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")  // // using .html extension
		if err != nil {
			log.Fatalln(err)
		}
		err = parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error while parsing template : ", err)
			// fmt.Fprintln(w, fmt.Sprintf("error while parsing template : %s\n", err))
		}
		tc[tmpl] = parsedTemplate
	} else {
		fmt.Println("using a cache template")
		err := tc[tmpl].Execute(w, nil)
		if err != nil {
			fmt.Println("error while parsing template : ", err)
			// fmt.Fprintln(w, fmt.Sprintf("error while parsing template : %s\n", err))
		}
	}
}
