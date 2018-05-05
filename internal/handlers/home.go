package handlers

import (
	"encoding/json"
	"github.com/arjunajithtp/excel-column-finder/internal/helpers"
	"github.com/arjunajithtp/excel-column-finder/internal/services"
	"log"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		r.ParseForm()
		startingColumn := r.Form.Get("startingColumn")
		rows, err := strconv.Atoi(r.Form.Get("rows"))
		if err != nil {
			log.Println("error while trying to convert row value to integer: ", err)
		}
		columns, err := strconv.Atoi(r.Form.Get("columns"))
		if err != nil {
			log.Println("error while trying to convert column value to integer: ", err)
		}
		words := services.LetterIncrementation(startingColumn, rows*columns)
		var wordSlices [][]string

		start := 0
		for i := 0; i < rows; i++ {
			end := (i + 1) * columns
			rowSlice := words[start:end]
			wordSlices = append(wordSlices, rowSlice)
			start = end
		}

		wordsByte, err := json.Marshal(wordSlices)
		if err != nil {
			log.Println("error while trying to marshal the map data: ", err)
		}
		_, err = w.Write(wordsByte)
		if err != nil {
			log.Println("error while trying to write data to UI: ", err)
		}
		return
	}
	err := helpers.RenderPage(w, "home", nil)
	if err != nil {
		log.Println("error while trying to render page on home handler: ", err)
	}
}
