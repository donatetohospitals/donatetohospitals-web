package handlers

import (
	"html/template"
	"net/http"

	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/donatetohospitals/donatetohospitals-web/utils"
)

var volunteerTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/volunteer.html",
	"templates/navigation.html",
	"templates/supplier.html",
)

func GetVolunteersPage(s core.DonationService, title string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		t := &core.Page{Title: title, WithFooter: false}

		err := volunteerTemplate.ExecuteTemplate(w, "layout", t)
		if err != nil {
			// TODO (daniel): Wrap and bubble error up
			utils.TemplateError(err, "render")
		}

	}

}
