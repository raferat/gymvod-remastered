package main

import (
	"html/template"
	"net/http"
)

func rootSite(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles("sites/index.tmpl", "components/navbar.tmpl", "components/footer.tmpl", "components/logo.svg", "components/menu.svg", "components/settings.svg"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", rootSite)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe("0.0.0.0:4040", nil)
}
