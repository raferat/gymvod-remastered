package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

type CurrentInformation []CurrentInformationAction

type CurrentInformationAction struct {
	StartTag    bool
	EndTag      bool
	TextTag     bool
	TextContent string // for text
	TagName     string
	Attrs       map[string]string
}

func normalizeText(input string) string {
	return strings.ReplaceAll(input, " \xc2\xa0", "")
}

func parseCurrentInformationHtmlContent(htmlContent string) []CurrentInformationAction {
	z := html.NewTokenizer(strings.NewReader(htmlContent))
	actions := make([]CurrentInformationAction, 0, 50)
	for {
		next := z.Next()
		if next == html.ErrorToken {
			break
		} else if next == html.StartTagToken || next == html.SelfClosingTagToken {
			tname, hasAttrs := z.TagName()
      if string(tname) == "script" {
        log.Fatalln("Panic because of XSS: " + htmlContent)
      }
			if string(tname) == "hr" {
				continue
			}
			attrs := make(map[string]string)
			for hasAttrs {
				name, value, next := z.TagAttr()
				hasAttrs = next
				attrs[string(name)] = string(value)
			}
			actions = append(actions, CurrentInformationAction{
				StartTag: true,
				EndTag:   false,
				TextTag:  false,
				TagName:  string(tname),
				Attrs:    attrs,
			})
		} else if next == html.EndTagToken {
			tname, _ := z.TagName()
			actions = append(actions, CurrentInformationAction{
				StartTag: false,
				EndTag:   true,
				TextTag:  false,
				TagName:  string(tname),
			})
		} else if next == html.TextToken {
			text := string(z.Text())
			text = strings.TrimSpace(text)
			text = normalizeText(text)
			if len(text) == 0 {
				continue
			}
			actions = append(actions, CurrentInformationAction{
				StartTag:    false,
				EndTag:      false,
				TextTag:     true,
				TextContent: text,
			})
		}
	}

	return actions
}

func fetchCurrentInformation() (result []CurrentInformation, err error) {
	res, err := http.Get("https://gymvod.cz/aktualni-informace/")
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return result, err
	}

	doc.Find(".flexaret").Children().Each(func(i int, s *goquery.Selection) {
		content, err := s.Html()
		if err != nil || len(content) == 0 {
			return
		}
		result = append(result, parseCurrentInformationHtmlContent(content))
	})

	return result, err
}

/*
func rootSite(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("sites/index.tpl.html", "components/navbar.tpl.html", "components/footer.tpl.html", "components/logo.svg", "components/menu.svg", "components/settings.svg")

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
*/

func rootSiteGin(c *gin.Context) {
	tmpl, err := template.ParseFiles("sites/index.tpl.html", "components/navbar.tpl.html", "components/footer.tpl.html", "components/logo.svg", "components/menu.svg", "components/settings.svg")

	if err != nil {
		log.Println(err)
    c.String(http.StatusInternalServerError, "<h1>Error</h1>")
    return
	}

	res, err := fetchCurrentInformation()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.String(http.StatusInternalServerError,"<h1>Internal server errrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrror</h1>")
	} else {
		c.Writer.WriteHeader(http.StatusOK)
		err = tmpl.Execute(c.Writer, res)
	}
}
