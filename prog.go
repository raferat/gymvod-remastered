package main

import (
	"net/http"
  "io"
  "log"
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

func main() {
	http.HandleFunc("/", rootSite)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/soubor/{id}", staticSouborRoute)
  http.HandleFunc("/obrazek/{id}", staticObrazekRoute)
  http.HandleFunc("/nahled/{id}", staticObrazekRoute)

	http.ListenAndServe("0.0.0.0:4040", nil)
	//fetchCurrentInformation()
}
