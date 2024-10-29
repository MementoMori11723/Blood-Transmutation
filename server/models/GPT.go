package models

import (
	"Blood-Transmutation/config"
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type GPTRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTRequest struct {
	Model    string              `json:"model"`
	Messages []GPTRequestMessage `json:"messages"`
}

type GPTResponse struct {}

func GetDataFromGpt(question, options string) GPTResponse {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	gptRequest := GPTRequest{
		Model: "gpt-4o-mini",
		Messages: []GPTRequestMessage{
			{
				Role:    "system",
				Content: options,
			},
			{
				Role:    "user",
				Content: question,
			},
		},
	}

	url := config.GetGPTTextData()
	jsonData, err := json.Marshal(gptRequest)
	if err != nil {
		panic(err)
  }

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	var body GPTResponse
  err = json.NewDecoder(res.Body).Decode(&body)
  if err != nil {
    panic(err)
  }

	return body
}
