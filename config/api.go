package config

var (
  GPTTextUrl = "https://api.openai.com/v1/chat/completions"
)

func GetGPTTextData() string {
  return GPTTextUrl
}
