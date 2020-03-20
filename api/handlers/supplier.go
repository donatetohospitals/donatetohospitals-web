package handlers

import (
	"errors"
	"net/http"

	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/go-chi/render"
)

func postSupplierHandler(w http.ResponseWriter, r *http.Request) {
	var supplier Supplier

	if err := render.DecodeJSON(r.Body, &supplier); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	create := db.Create(&supplier)

	if create.Error != nil {
		render.Render(w, r, ErrInvalidRequest(errors.New("unable to save record in db")))
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func supplyHandler(s core.DonationService, title string, errorHandler templateErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// TODO (daniel): fetch suppliers from service via s.GetAll() or some method alike

		s := []core.Supplier{{}, {}, {}, {}, {}}
		t := &core.Page{Title: title, WithFooter: false}

		err := supplyTemplate.ExecuteTemplate(w, "layout", t)
		errorHandler(err, "render")
	}
}
