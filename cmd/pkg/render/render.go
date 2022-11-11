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
func Render_Template(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//  check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		fmt.Println("creating the template and adding that into cache for future usees .")
		/*
			fmt.Println("temp not available but we are now creating a new template to use")
			parsedTemplate, err := template.ParseFiles("./templates/"+t, "./templates/base.layout.tmpl") // using .tmpl extension
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

		*/

		// need to create the template
		if err = createTemplateCache(t); err != nil {
			log.Println(err)
		}
	} else {

		// we have the template in the cache
		fmt.Println("using a cache template")
		/*
			err := tc[tmpl].Execute(w, nil)
			if err != nil {
				fmt.Println("error while parsing template : ", err)
				// fmt.Fprintln(w, fmt.Sprintf("error while parsing template : %s\n", err))
			}
		*/
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./template/%s", t),
		"./template/base.layout.tmpl",
	}
	parsedTemplate, err := template.ParseFiles(templates...)
	if err != nil {
		// log.Println("failed to parse template files")
		return err
	}
	tc[t] = parsedTemplate
	return nil
}
