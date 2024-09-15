package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
  })
  
  fmt.Println("Server is running on port 8080")
  http.ListenAndServe(":8080", nil)
}
