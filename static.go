package main

import (
	"log"
	"net/http"
)

func staticSouborRoute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	log.Println(id)
}
