package models

import (
	"Blood-Transmutation/config"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type gptRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type gptRequest struct {
	Model       string              `json:"model"`
	Messages    []gptRequestMessage `json:"messages"`
	Temperature int64               `json:"temperature"`
}

type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func getDataFromGpt(question, options string) (string, error) {
	var responce GPTResponse

	client := http.Client{}

	if question == "" {
		return "", errors.New(
			"No question found!",
		)
	}

	if options == "" {
		return "", errors.New(
			"No question found!",
		)
	}

	gptRequest := gptRequest{
		Model: "deepseek-chat",
		Messages: []gptRequestMessage{
			{
				Role:    "system",
				Content: options,
			},
			{
				Role:    "user",
				Content: question,
			},
		},
		Temperature: 0,
	}

	url, api := config.GetGPTTextData()

	jsonData, err := json.
		Marshal(gptRequest)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST", url,
		bytes.
			NewBuffer(jsonData),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+api,
	)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	err = json.
		NewDecoder(res.Body).
		Decode(&responce)
	if err != nil {
		panic(err)
	}

	return responce.
			Choices[0].
			Message.
			Content,
		nil
}
