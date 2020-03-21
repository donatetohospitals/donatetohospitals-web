package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/donatetohospitals/donatetohospitals-web/api/handlers"
	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type APIServer struct {
	router *chi.Mux
}

func (as *APIServer) setupRoutes(service *core.DonationService) {

	fs := http.FileServer(http.Dir("./front"))

	as.router.Use(middleware.Logger)

	as.router.Get(
		"/",
		handlers.GetIndexPage(service, "lorem ipsum dolor sit amet")
	)

	as.router.Get(
		"/about",
		handlers.GetAboutPage(service, "lorem ipsum dolor sit amet")
	)

	as.router.Get(
		"/volunteer",
		handlers.GetVolunteersPage(service, "lorem ipsum")
	)

	as.router.Get(
		"/supply",
		handlers.GetSuppliersPage(service, "lorem ipsum dolor sit amte")
	)

	as.router.Route(
		"/suppliers", func(r chi.Router) {
		r.Post("/", handlers.PostSupplier(service))
	})

	as.router.Get(
		"/front/vendor",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fs.ServeHTTP(w, r)
		})
	)

	// File server
	basePath := "/assets"

	as.router.Route(basePath, func(root chi.Router) {
		workDir, _ := os.Getwd()
		filesDir := filepath.Join(workDir, "front")

		// TODO (daniel): Refactor this server
		fileServer(root, basePath, "/front", http.Dir(filesDir))
	})

	allowedOrigins := []string{"*"}

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

type APIConfiguration struct {
	service *core.DonationService
	// Cors, Logger, certificates and so on
}

func NewAPI(ac APIConfiguration) (APIServer, error) {
	r := chi.NewRouter()
	apiServer := APIServer{
		router: r,
	}

	apiServer.setupRoutes(ac.service)

	return apiServer, nil
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(basePath+path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"
	fmt.Println(basePath, path, root, basePath+path)
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
