package handlers

import (
	"html/template"
	"net/http"

	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/donatetohospitals/donatetohospitals-web/utils"
)

var supplyTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/postSupplies.html",
	"templates/navigation.html")

func GetSuppliersPage(s core.DonationService, title string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// TODO (daniel): fetch suppliers from service via s.GetAll() or some method alike

		//s := []core.Supplier{{}, {}, {}, {}, {}}
		t := &core.Page{Title: title, WithFooter: false}

		err := supplyTemplate.ExecuteTemplate(w, "layout", t)

		if err != nil {
			// TODO (daniel): Wrap and bubble error
			utils.TemplateError(err, "render")
		}
	}
}
