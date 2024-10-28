package server

import (
  "net/http"
)

var (
  routes = map[string]func(http.ResponseWriter, *http.Request){
    "/": api,
  }
)

func New() *http.ServeMux {
  mux := http.NewServeMux()
  for route, handler := range routes {
    mux.HandleFunc(route, handler)
  }
  return mux
}
