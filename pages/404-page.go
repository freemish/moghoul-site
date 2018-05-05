package pages

import (
	"html/template"
	"log"
	"net/http"
)

func NotFoundPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/data/404.html", "templates/layouts/main-layout.html")
	tmpl.ExecuteTemplate(w, "layout", nil)

	if err != nil {
		log.Println(err.Error() + " (in NotFoundPageHandler)")
	}
}
