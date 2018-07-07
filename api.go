package main

import (
	"net/url"
	"net/http"
	"encoding/json"
	"log"
	"encoding/xml"
)


type Hash struct {
	Hash string ` json:"hash" xml:"hash" `
}

type Params struct {
	url.Values
}

func GetIP(r *http.Request) string {
	data := r.Header.Get("X-FORWARDED-FOR")

	if len(data) == 0 {
		data = r.RemoteAddr
	}
	return data
}

func GetIPJson(w http.ResponseWriter, r *http.Request) {

	data := GetIP(r)

	payload, err := json.Marshal(data)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Request received on: %s\n", r.URL.Path)
	log.Printf("IP: %s\n", data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func GetIPString(w http.ResponseWriter, r *http.Request) {

	data := GetIP(r)

	log.Printf("Request received on: %s\n", r.URL.Path)
	log.Printf("IP: %s\n", data)

	w.Write([]byte(data))
}

func GetIPXML(w http.ResponseWriter, r *http.Request) {

	data := GetIP(r)

	payload, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Request received on: %s\n", r.URL.Path)
	log.Printf("IP: %s\n", data)

	w.Header().Set("Content-Type", "application/xml")
	w.Write(payload)
}

