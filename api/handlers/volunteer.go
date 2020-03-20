package handlers

import (
	"html/template"
	"net/http"

	"github.com/donatetohospitals/donatetohospitals-web/core"
)

var volunteerTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/volunteer.html",
	"templates/navigation.html",
	"templates/supplier.html")

func volunteerHandler(w http.ResponseWriter, r *http.Request) {
	t := &Page{Title: templateTitle, WithFooter: false}
	err := volunteerTemplate.ExecuteTemplate(w, "layout", t)
	handleErr(err, "render")
}

func supplyHandler(s core.DonationService, title string, errorHandler templateErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// TODO (daniel): fetch suppliers from service via s.GetAll() or some method alike

		t := &core.Page{Title: title, WithFooter: false}
		err := volunteerTemplate.ExecuteTemplate(w, "layout", t)
		errorHandler(err, "render")
	}
}
