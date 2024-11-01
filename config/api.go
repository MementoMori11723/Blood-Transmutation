package config

import (
	"log"
	"os"
)

var (
  GPTTextUrl = "https://api.openai.com/v1/chat/completions"
  GPTApiKey string
)

func init() {
  GPTApiKey = os.Getenv("API_KEY")
  if GPTApiKey == "" {
    log.Fatal("Can't find api key!")
  }
}

func GetGPTTextData() (string, string) {
  return GPTTextUrl, GPTApiKey
}
