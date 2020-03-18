package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var isProduction = os.Getenv("TARGET_ENV") == "release"
var keyFile = os.Getenv("KEY_FILE")
var certFile = os.Getenv("CERT_FILE")

var templateTitle = "DonateToHospitals - Donate your stockpiles to help covid19 corona virus relief"

type Page struct {
	Title      string
	Suppliers  []Supplier
	WithFooter bool
}

func handleErr(err error, context string) {
	if err != nil {
		fmt.Println("Template rendering error with ", context+": ", err)
	}
}

var indexTemplate, _ = template.ParseFiles(
	"templates/layout.html",
	"templates/index.html",
	"templates/navigation.html",
	"templates/supplier.html")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO will be fetched from the db when there is one lul
	s := []Supplier{{}, {}, {}, {}, {}}
	t := &Page{Title: templateTitle, Suppliers: s, WithFooter: true}
	err := indexTemplate.ExecuteTemplate(w, "layout", t)
	handleErr(err, "render")
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
	handleErr(err, "render")
}

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

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Render is the Renderer for ErrResponse struct
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest is used to indicate an error on user input (with wrapped error)
func ErrInvalidRequest(err error) render.Renderer {
	var errorText string
	if err != nil {
		errorText = err.Error()
	}
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      errorText,
	}
}

// ErrResponse is a generic struct for returning a standard error document
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func postSupplierHandler(w http.ResponseWriter, r *http.Request) {
	var supplier Supplier

	if err := render.DecodeJSON(r.Body, &supplier); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	create := db.Create(&supplier)

	// TODO also persist items

	if create.Error != nil {
		render.Render(w, r, ErrInvalidRequest(errors.New("unable to save record in db")))
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

type Configuration struct {
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	DbName   string
	Port     string
	User     string
	Password string
}

type Supplier struct {
	gorm.Model
	Email       string `gorm:"type:varchar(100);unique_index;not null"`
	Geo         string `gorm:"index:geo"` // state etc. create index with name `loc` for address
	ImageUrl    string `gorm:"size:255"`  // set field size to 255
	Items       []Item `gorm:"foreignkey:SupplierRefer"`
	IsAllocated bool
}

type Item struct {
	gorm.Model
	Name          string
	Count         int
	SupplierRefer uint
}

var db *gorm.DB

func main() {
	fmt.Println("Starting server on port :9990")

	//config
	var configuration Configuration
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	createdDb, err := gorm.Open("postgres", "host=localhost port="+configuration.Database.Port+" user="+configuration.Database.User+" dbname="+configuration.Database.DbName+" password="+configuration.Database.Password+" sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	// TODO figure out how to assign to global variable without doing this
	db = createdDb

	defer db.Close()

	db.AutoMigrate(&Supplier{})
	db.AutoMigrate(&Item{})

	fs := http.FileServer(http.Dir("./front"))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", indexHandler)
	r.Get("/about", aboutHandler)
	r.Get("/volunteer", volunteerHandler)

	r.Route("/suppliers", func(r chi.Router) {
		r.Post("/", postSupplierHandler)
	})
	r.Get("/front/vendor", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

	// file server
	basePath := "/assets"

	r.Route(basePath, func(root chi.Router) {
		workDir, _ := os.Getwd()
		filesDir := filepath.Join(workDir, "front")
		FileServer(root, basePath, "/front", http.Dir(filesDir))
	})

	allowedOrigins := []string{"*"}

	if isProduction {
		allowedOrigins = []string{"https://donatetohospitals.com", "https://www.donatetohospitals.com"}
		err := http.ListenAndServeTLS(":9990", certFile, keyFile, r)

		if err != nil {
			fmt.Println("ListenAndServeTLS:", err)
		}
	} else {
		err := http.ListenAndServe(":9990", r)
		if err != nil {
			fmt.Println("ListenAndServe:", err)
		}
	}

	corsAssigned := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST"}, // , "PUT", "DELETE", "PATCH", "OPTIONS"
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsAssigned.Handler)

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
