package config

import (
	"log"
	"os"
)

var (
  PORT string
)

func init() {
  PORT = os.Getenv("PORT")
  if PORT == "" {
    log.Println("PORT not found, using default value")
    PORT = "8000"
  }
}

func GetPort() string {
  return PORT
}
