package handlers

import (
	"html/template"
	"net/http"

	"github.com/donatetohospitals/donatetohospitals-web/core"
)

var indexTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/index.html",
	"templates/navigation.html",
	"templates/supplier.html")

func indexHandler(s core.DonationService, title string, errorHandler templateErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// TODO (daniel): fetch suppliers from service via s.GetAll() or some method alike

		s := []core.Supplier{{}, {}, {}, {}, {}}
		t := &core.Page{Title: title, Suppliers: s, WithFooter: true}

		err := indexTemplate.ExecuteTemplate(w, "layout", t)
		errorHandler(err, "render")
	}
}
