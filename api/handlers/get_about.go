package handlers

import (
	"html/template"
	"net/http"

	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/donatetohospitals/donatetohospitals-web/utils"
)

// TODO find out how not to have to do this for each page in order to cache it?
var aboutTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/about.html",
	"templates/navigation.html",
	"templates/supplier.html")

func GetAboutPage(
	s core.DonationService,
	title string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := &core.Page{Title: title, WithFooter: false}
		err := aboutTemplate.ExecuteTemplate(w, "layout", t)

		if err != nil {
			utils.TemplateError(err, "render")
		}

	}
}
