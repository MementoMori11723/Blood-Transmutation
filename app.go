package main

import (
	"Blood-Transmutation/client"
	"Blood-Transmutation/config"
	"Blood-Transmutation/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
  go func ()  {
    client := client.New()
    server := server.New()
    port := config.GetPort()
    http.Handle("/", client)
    http.Handle("/api", server)
    log.Println("Application running on http://localhost:" + port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
      log.Fatal(err)
    }
  }()
  fmt.Scanln()
}
