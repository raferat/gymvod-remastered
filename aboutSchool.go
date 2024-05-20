package main

import (
	"log"
	"net/http"
	"text/template"
  "errors"

	"github.com/gin-gonic/gin"
)

type AboutInputData struct {
  UciteleTable [][]string
}

// func aboutSchoolSite(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.ParseFiles("sites/about.tpl.html", "components/navbar.tpl.html", "components/footer.tpl.html", "components/logo.svg", "components/menu.svg", "components/settings.svg")

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	res, err := fetchCurrentInformation()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("<h1>Internal server errrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrror</h1>"))
// 	} else {
// 		w.WriteHeader(http.StatusOK)
// 		err = tmpl.Execute(w, res)
// 	}
// }

func aboutSchoolSiteGin(c *gin.Context) {
  //tmpl, err := template.ParseFiles("sites/about.tpl.html","sites/ucitele.tpl.html", "components/navbar.tpl.html", "components/footer.tpl.html", "components/logo.svg", "components/menu.svg", "components/settings.svg")
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
  }).ParseFiles("sites/about.tpl.html","sites/ucitele.tpl.html", "components/navbar.tpl.html", "components/footer.tpl.html", "components/logo.svg", "components/menu.svg", "components/settings.svg")


	if err != nil {
		log.Println(err)
    c.String(http.StatusInternalServerError, "<h1>Internal Server Error</h1>")
    return
  }

  dataframe := &AboutInputData{
    UciteleTable: getContacts(),
  }

  c.Writer.WriteHeader(http.StatusOK)
  err = tmpl.Execute(c.Writer, dataframe)
}
