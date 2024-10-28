package client

import (
  "net/http"
)

var (
  routes = map[string]func(http.ResponseWriter, *http.Request){
    "/": home,
    "/about": about,
    "/dashboard": dashboard,
    "/error": error,
  }
)

func New() *http.ServeMux {
  mux := http.NewServeMux()
  for route, handler := range routes { 
    mux.HandleFunc(route, handler)
  }
  return mux
}
