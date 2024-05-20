package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AboutInputData struct {
  UciteleTable [][]string
}

func aboutSchoolSiteGin(c *gin.Context) {
  dataframe := &AboutInputData{
    UciteleTable: getContacts(),
  }

  c.Writer.WriteHeader(http.StatusOK)
  err := TEMPLATES.ExecuteTemplate(c.Writer, "about-site", dataframe)

  if err != nil {
    log.Printf("%+v\n", err)
  }
}
