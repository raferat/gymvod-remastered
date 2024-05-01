package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", rootSite)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/soubor/{id}", staticSouborRoute)

	http.ListenAndServe("0.0.0.0:4040", nil)
	//fetchCurrentInformation()
}
