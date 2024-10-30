package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
  GPTTextUrl = "https://api.openai.com/v1/chat/completions"
  GPTApiKey string
)

func init() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }
  GPTApiKey = os.Getenv("API_KEY")
  if GPTApiKey == "" {
    log.Fatal("Can't find api key!")
  }
}

func GetGPTTextData() (string, string) {
  return GPTTextUrl, GPTApiKey
}
