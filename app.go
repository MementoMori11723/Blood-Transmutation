package main

import (
	"Blood-Transmutation/server"
	"Blood-Transmutation/webscraper"
	"fmt"
)

var (
  url = "https://www.google.com/search?q=blood&start="
  urlArr = []string{}
)

func main() {
	for i := 1; i < 11; i++ {
		data := webscraper.GetContent(url + fmt.Sprint(i))
		if data.Err != nil {
			fmt.Println("Error:", data.Err)
			return
		}
		data.ParceURL(&urlArr)
	}
  go server.StartServer(&urlArr)
  fmt.Scanln()
}
