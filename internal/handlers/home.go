package handlers

import (
	"github.com/arjunajithtp/excel-column-finder/internal/helpers"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		r.ParseForm()

		return
	}
	err := helpers.RenderPage(w, "home", nil)
	if err != nil {
		log.Println("error while trying to render page on home handler: ", err)
	}
}
