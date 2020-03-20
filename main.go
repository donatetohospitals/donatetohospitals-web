package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var isProduction = os.Getenv("TARGET_ENV") == "release"
var keyFile = os.Getenv("KEY_FILE")
var certFile = os.Getenv("CERT_FILE")

var templateTitle = "DonateToHospitals - Donate your stockpiles to help covid19 corona virus relief"

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

var db *gorm.DB

func main() {
	fmt.Println("Starting server on port :9990")

	/*
		Application flow:
		1. Read config vars
		2. Instantiate database client
		3. Instantiate service and pass it a reference to the repository
		3.5 Parse templates
		4. Instantiate a server and pass a reference to the service
		5. Start accepting connections
	*/

	//config
	var configuration core.Configuration
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

	// TODO figure out how to assign to global variable without doing this

	defer db.Close()

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
