package main

import (
	"Blood-Transmutation/config"
	"Blood-Transmutation/server"
	"net/http"
	"log"
)

func main() {
  server := server.New()
  port := config.GetPort()
  log.Println("Application running on http://localhost:" + port)
  err := http.ListenAndServe(":" + port, server)
  if err != nil {
    log.Fatal(err)
  }
}
