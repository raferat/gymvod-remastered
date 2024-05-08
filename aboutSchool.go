package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)



func aboutSchoolSite(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("sites/about.tmpl", "components/navbar.tmpl", "components/footer.tmpl", "components/logo.svg", "components/menu.svg", "components/settings.svg")

	if err != nil {
		log.Println(err)
	}

	res, err := fetchCurrentInformation()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("<h1>Internal server errrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrror</h1>"))
	} else {
		w.WriteHeader(http.StatusOK)
		err = tmpl.Execute(w, res)
	}
}

func aboutSchoolSiteGin(c *gin.Context) {
  tmpl, err := template.ParseFiles("sites/about.tmpl", "components/navbar.tmpl", "components/footer.tmpl", "components/logo.svg", "components/menu.svg", "components/settings.svg")

	if err != nil {
		log.Println(err)
    c.String(http.StatusInternalServerError, "<h1>Internal Server Error</h1>")
    return
	}

	res, err := fetchCurrentInformation()
	if err != nil {
    c.String(http.StatusInternalServerError, "<h1>Internal server error, the school website is probably down.</h1>")
    return
	} else {
    //c.Status(http.StatusOK)
    c.Writer.WriteHeader(http.StatusOK)
		err = tmpl.Execute(c.Writer, res)
	}
}
