package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type APIServer struct {
	router *chi.Mux
}

func (as *APIServer) setupRoutes(service *core.DonationService) {

	fs := http.FileServer(http.Dir("./front"))

	r.Use(middleware.Logger)
	r.Get("/", indexHandler(service, "lorem ipsum dolor sit amet", handleErr))
//	r.Get("/about", aboutHandler)
//	r.Get("/volunteer", volunteerHandler)
//	r.Get("/supply", supplyHandler)
//
//	r.Route("/suppliers", func(r chi.Router) {
//		r.Post("/", postSupplierHandler)
//	})
//
//	r.Get("/front/vendor", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fs.ServeHTTP(w, r)
//	}))
//
//	// file server
//	basePath := "/assets"
//
//	r.Route(basePath, func(root chi.Router) {
//		workDir, _ := os.Getwd()
//		filesDir := filepath.Join(workDir, "front")
//		FileServer(root, basePath, "/front", http.Dir(filesDir))
//	})
//
//	allowedOrigins := []string{"*"}
//
//	corsAssigned := cors.New(cors.Options{
//		AllowedOrigins:   allowedOrigins,
//		AllowedMethods:   []string{"GET", "POST"}, // , "PUT", "DELETE", "PATCH", "OPTIONS"
//		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
//		AllowCredentials: true,
//		MaxAge:           300, // Maximum value not ignored by any of major browsers
//	})
//
//	r.Use(corsAssigned.Handler)

}

func (as *APIServer) Listen() {}

type APIConfiguration {
	service *core.DonationService
	// Cors, Logger, certificates and so on
}

func NewAPI(ac APIConfiguration) (APIServer, error) {
	r := chi.NewRouter()
	apiServer := &APIServer{
		router: r,
	}

	apiServer.setupRoutes(ac.service)

	return apiService, nil
}
