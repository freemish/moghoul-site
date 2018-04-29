package pages

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func TodoPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/todo.html"))

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Aerobics", Done: true},
			{Title: "Happiness", Done: true},
			{Title: "Diddly-squat", Done: false},
		},
	}

	tmpl.Execute(w, data)
}
