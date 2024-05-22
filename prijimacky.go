package main

import (
  "github.com/gin-gonic/gin"

  "net/http"
  "log"
)

func prijimackySiteGin (c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
  err := TEMPLATES.ExecuteTemplate(c.Writer, "prijimacky-site", nil)

  if err != nil {
    log.Println(err)
  }
}
