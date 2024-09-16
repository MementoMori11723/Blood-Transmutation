package server

import (
	"fmt"
	"net/http"
)

func StartServer(arr *[]string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URLs: %v", *arr)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
