package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	Title      string
	Suppliers  []Supplier
	WithFooter bool
}

type Supplier struct {
	Index int8
}

var templateTitle = "DonateToHospitals - Donate your stockpiles to help covid19 corona virus relief"

func handleRenderErr(err error) {
	if err != nil {
		fmt.Println("Template rendering error:", err)
	}
}

var indexTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/index.html",
	"templates/navigation.html",
	"templates/supplier.html")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO will be fetched from the db when there is one lul
	s := []Supplier{{Index: 1}, {Index: 2}, {Index: 3}, {Index: 4}, {Index: 5}}
	t := &Page{Title: templateTitle, Suppliers: s, WithFooter: true}
	err := indexTemplate.ExecuteTemplate(w, "layout", t)
	handleRenderErr(err)
}

// TODO find out how not to have to do this for each page in order to cache it?
var aboutTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/about.html",
	"templates/navigation.html",
	"templates/supplier.html")

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	t := &Page{Title: templateTitle, WithFooter: false}
	err := aboutTemplate.ExecuteTemplate(w, "layout", t)
	handleRenderErr(err)
}

var volunteerTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/volunteer.html",
	"templates/navigation.html",
	"templates/supplier.html")

func volunteerHandler(w http.ResponseWriter, r *http.Request) {
	t := &Page{Title: templateTitle, WithFooter: false}
	err := volunteerTemplate.ExecuteTemplate(w, "layout", t)
	handleRenderErr(err)
}

func main() {
	fmt.Println("Starting server on port :3000")
	fs := http.FileServer(http.Dir("./front"))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", indexHandler)
	r.Get("/about", aboutHandler)
	r.Get("/volunteer", volunteerHandler)
	r.Get("/front/vendor", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

	// fileserver
	basePath := "/assets"
	r.Route(basePath, func(root chi.Router) {
		workDir, _ := os.Getwd()
		filesDir := filepath.Join(workDir, "front")
		FileServer(root, basePath, "/front", http.Dir(filesDir))
	})

	// TODO: serve https on prod
	err := http.ListenAndServe(":3000", r)

	corsAssigned := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},           // TODO: make conditional to prod env to only accept correct origin
		AllowedMethods:   []string{"GET", "POST"}, // , "PUT", "DELETE", "PATCH", "OPTIONS"
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsAssigned.Handler)

	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}

}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
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
