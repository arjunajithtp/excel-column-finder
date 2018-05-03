package helpers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderPage(w http.ResponseWriter, pageName string, data interface{}) error {
	tmpl := template.Must(template.ParseFiles("./public/view/" + pageName + ".html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		return fmt.Errorf("error while trying to render the html page: %v", err)
	}
	return nil
}
