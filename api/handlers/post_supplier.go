package handlers

import (
	"errors"
	"net/http"

	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/donatetohospitals/donatetohospitals-web/utils"
	"github.com/go-chi/render"
)

func PostSupplier(s *core.DonationService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var supplier core.Supplier

		if err := render.DecodeJSON(r.Body, &supplier); err != nil {
			render.Render(w, r, utils.ErrInvalidRequest(err))
			return
		}

		// TODO(daniel) Replace this with call to suppliers.Create()
		db := s.GetDatabase()
		create := db.Create(&supplier)

		if create.Error != nil {
			render.Render(w, r, utils.ErrInvalidRequest(errors.New("unable to save record in db")))
			return
		}

		utils.ResponseJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})

	}
}
