package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates using html templates
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template", err)
// 		return
// 	}
// }

var template_cache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, page string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in cache
	_, inMap := template_cache[page]
	if !inMap {
		// need to create template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(page)
		if err != nil {
			log.Println(err)
		}

	} else {
		// we have the template in cache
		log.Println("using cached template")
	}

	tmpl = template_cache[page]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(page string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", page),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to the cache(map)
	template_cache[page] = tmpl
	return nil
}
