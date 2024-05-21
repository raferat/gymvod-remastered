package main

import (
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func makeRequest(url string, method string, body io.Reader) (resp *http.Response) {
  req, err := http.NewRequest(method, url, body)

  if err != nil {
    log.Println(err)
    return resp
  }

  req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.67")
  
  resp, err = http.DefaultClient.Do(req)

  if err != nil {
    log.Println(err)
    return resp
  }

  return resp
}

var TEMPLATES *template.Template

func loadAllTemplates() {
  files := make([]string, 0, 50)

  matched, err := filepath.Glob("sites/*")
  if err != nil {
    log.Fatalf("%+v\n", err)
  }
  for _,v := range matched {
    files = append(files, v)
  }
  matched, err = filepath.Glob("components/*")
  if err != nil {
    log.Fatalf("%+v\n", err)
  }
  for _,v := range matched {
    files = append(files, v)
  }

  log.Printf("Loaded: %+v\n", files)

  tmpl, err := template.New("").Funcs(template.FuncMap{
    "dict": func(values ...interface{}) (map[string]interface{}, error) {
      if len(values)%2 != 0 {
        return nil, errors.New("invalid dict call")
      }
      dict := make(map[string]interface{}, len(values)/2)
      for i := 0; i < len(values); i+=2 {
        key, ok := values[i].(string)
        if !ok {
          return nil, errors.New("dict keys must be strings")
        }
        dict[key] = values[i+1]
      }
      return dict, nil
    },
    "unesc": func(value string) (interface{}) {
      log.Println(value)
      a := template.HTML(value)
      return a
    },
  }).ParseFiles(files...)

  if err != nil {
    log.Fatalf("%+v\n", err)
  }

  TEMPLATES=tmpl;
}

func notFound(c *gin.Context) {
  c.Writer.WriteHeader(http.StatusNotFound)
  err := TEMPLATES.ExecuteTemplate(c.Writer, "404-site" ,"")
  if err != nil {
    log.Println(err)
  }
}

func main() {
  loadAllTemplates()


  ticker := time.NewTicker(time.Second)
  updateTemplatesDone := make(chan bool)

  go func() {
    for {
      select {
        case <-updateTemplatesDone:
          return
        case <-ticker.C:
          loadAllTemplates()
      }
    }
  }()

  defer func () {
    ticker.Stop()
    updateTemplatesDone <- true
  } ()

  r := gin.Default()
  r.Use(gzip.Gzip(gzip.DefaultCompression))
  r.Static("/static", "static/")
  r.StaticFile("/robots.txt", "./robots.txt")
  r.NoRoute(notFound)
  r.GET("/", rootSiteGin)
  r.GET("/ucitele", uciteleSiteGin)
  r.GET("/about", aboutSchoolSiteGin)
  r.GET("/soubor/:id", staticSouborRouteGin)
  r.GET("/obrazek/:id", staticObrazekRouteGin)
  r.GET("/nahled/:id", staticNahledRouteGin)

  if err := r.Run(":4040"); err != nil {
    log.Fatal(err)
  }
}
