package pages

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

type SubmittedContactDetails struct {
	Success bool
	Contact ContactDetails
}

func ContactFormPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/templates/contactform.html"))

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	submission := SubmittedContactDetails{
		Success: false,
		Contact: details,
	}

	tmpl.Execute(w, submission)
}
