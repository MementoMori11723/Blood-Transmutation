package config

import (
	"log"
	"os"
)

var (
	PORT      string
	GPTApiKey string

	GPTTextUrl = "https://api.openai.com/v1/chat/completions"
)

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Println("PORT not found, using default value")
		PORT = "8000"
	}
	GPTApiKey = os.Getenv("API_KEY")
	if GPTApiKey == "" {
		log.Fatal("Can't find api key!")
	}
}

func GetPort() string {
	return PORT
}

func GetGPTTextData() (string, string) {
	return GPTTextUrl, GPTApiKey
}
