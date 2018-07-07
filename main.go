package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	"time"
	"path/filepath"
	"os"
	"strings"
	"log"
	"flag"
	"bytes"
)

var siteUrl = "https://ipcheck.fun"
var siteName = "IPCheck.fun"
var host = ":8009"
var tmpl = ParseTemplates()
var starttime = time.Now()


func main() {
	flag.StringVar(&siteUrl, "siteurl", siteUrl, "The site url")
	flag.StringVar(&siteName, "sitename", siteName, "The name of the site")
	flag.StringVar(&host, "host", host, "The host and port (localhost:8009)")

	r := mux.NewRouter()
	r.HandleFunc("/favicon.ico", FileHandler).Methods("GET")
	r.HandleFunc("/favicon.png", FileHandler).Methods("GET")
	r.HandleFunc("/apple-touch-icon.png", FileHandler).Methods("GET")
	r.HandleFunc("/browserconfig.xml", FileHandler).Methods("GET")
	r.HandleFunc("/crossdomain.xml", FileHandler).Methods("GET")
	r.HandleFunc("/humans.txt", FileHandler).Methods("GET")
	r.HandleFunc("/robots.txt", FileHandler).Methods("GET")
	r.HandleFunc("/tile.png", FileHandler).Methods("GET")
	r.HandleFunc("/tile-wide.png", FileHandler).Methods("GET")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/json", GetIPJson).Methods("GET")
	r.HandleFunc("/xml", GetIPXML).Methods("GET")
	r.HandleFunc("/string", GetIPString).Methods("GET")
	r.HandleFunc("/", Index)
	log.Printf("Running on %s\n", host)
	http.ListenAndServe(host, r)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.WriteHeader(status)

	log.Printf("Error on: %s %v - %v\n", r.URL.Path, status, err)
	if status == http.StatusNotFound {
		tmpl.ExecuteTemplate(w, "404.html",nil)
		tmpl.Execute(w, nil)
	} else if status == http.StatusInternalServerError {
		tmpl.ExecuteTemplate(w, "500.html",nil)
	}
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Finding: %s\n", "assets" + r.URL.Path)
	filestring, err := Asset("assets" + r.URL.Path)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound, err)
		return
	}
	_, file := filepath.Split(r.URL.Path)
	log.Printf("File: %s\n", file)
	reader := bytes.NewReader([]byte(filestring))
	http.ServeContent(w, r, file, starttime, reader)
}

func ParseTemplates() *template.Template {
	log.Printf("Building template files\n")
	tmpl := template.New("")
	err := filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		log.Printf("Parsing: %s\n", path)
		if strings.Contains(path, ".html") {
			templateString, err := Asset(path)
			if err != nil {
				log.Fatal(err)
			} else {
				_, file := filepath.Split(path)
				_, err = tmpl.New(file).Parse(string(templateString))
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}
	log.Printf("Templates ready\n")

	return tmpl
}