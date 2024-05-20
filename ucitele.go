package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func getContacts() [][]string {
  resp := makeRequest("https://gymvod.cz/ucitele", "GET", nil)
  doc, err := goquery.NewDocumentFromReader(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  table := make([][]string, 0, 20) 

  doc.Find("div.flexaret table tr").Each(func(i int, s *goquery.Selection) {
    if s.Children().Length() != 5 {
      return
    }

    if s.Children().First().Children().Text() == "Jméno" {
      return
    }

    record := make([]string, 0, 5)

    s.Children().Each(func(in int, se *goquery.Selection) {
      record = append(record,se.Children().Text())
    })

    table = append(table, record)
  })

  return table
}

func uciteleSiteGin (c *gin.Context) {
  tmpl, err := template.ParseFiles("sites/ucitele.tpl.html", "components/navbar.tpl.html", "components/footer.tpl.html", "components/logo.svg", "components/menu.svg", "components/settings.svg")

	if err != nil {
		log.Println(err)
    c.String(http.StatusInternalServerError, "<h1>Error</h1>")
    return
	}

  dataFrame := getContacts()

	c.Writer.WriteHeader(http.StatusOK)
	err = tmpl.Execute(c.Writer, dataFrame)

  if err != nil {
    log.Println(err)
  }
}
