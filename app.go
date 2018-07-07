package main

import (
	"time"
	"net/http"
)

type IndexPageData struct {
	Date string
	Title string
	Count string
	Difficulty int64
	Strings string
	Remember string
	Action string
	MoreStyles []string
	MoreScripts []string
}

func Index(w http.ResponseWriter, r *http.Request) {

	data := new(IndexPageData)
	var err error
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound, err)
		return
	}
	data.Date = time.Now().Format("2006")
	data.Title = siteName

	err = tmpl.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}
}
