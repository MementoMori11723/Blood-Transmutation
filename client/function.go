package client

import (
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("home"))
}

func about(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("about"))
}

func dashboard(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("dashboard"))
}

func error(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("error"))
}
