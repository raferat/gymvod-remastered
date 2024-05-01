package main

import (
	"fmt"
	"log"
	"net/http"
)

func staticSouborRoute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

  url := fmt.Sprintf("https://gymvod.cz/soubor/%s", id)

  resp := makeRequest(url, "GET", nil)

  w.WriteHeader(resp.StatusCode)

  const bufferLength = 8192

  buffer := make([]byte, bufferLength)

  for {
    n, err := resp.Body.Read(buffer)
    if err != nil {
      log.Println(err)
      return
    }
    w.Write(buffer)

    if n != 8192 {
      break
    }
  }
}

func staticObrazekRoute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

  url := fmt.Sprintf("https://gymvod.cz/obrazek/%s", id)

  resp := makeRequest(url, "GET", nil)

  w.WriteHeader(resp.StatusCode)

  const bufferLength = 8192

  buffer := make([]byte, bufferLength)

  for {
    n, err := resp.Body.Read(buffer)
    if err != nil {
      log.Println(err)
      return
    }
    w.Write(buffer)

    if n != 8192 {
      break
    }
  }
}

func staticNahledRoute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

  url := fmt.Sprintf("https://gymvod.cz/nahled/%s", id)

  resp := makeRequest(url, "GET", nil)

  w.WriteHeader(resp.StatusCode)

  const bufferLength = 8192

  buffer := make([]byte, bufferLength)

  for {
    n, err := resp.Body.Read(buffer)
    if err != nil {
      log.Println(err)
      return
    }
    w.Write(buffer)

    if n != 8192 {
      break
    }
  }
}
