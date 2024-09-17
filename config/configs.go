package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Url       = "https://www.google.com/search?q="
	Pages     = "&start="
	Questions = []string{
		"what is blood",
		"what is blood made of",
		"how is blood made",
		"what are the types of blood",
		"what is blood used for",
		"what is the difference between blood types",
		"what is safe blood transfusion",
		"how many different chemicals are in blood",
		"how many different chemical reactions are in blood",
		"how many different chemical reactions are there in chemistry",
		"what type of chemicals are used in medical field",
		"what chemicals are legal in the medical field",
	}
)

func init() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }
}

func ApiKey() string {
  return os.Getenv("ApiKey")
}
