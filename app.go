package main

import (
	"Blood-Transmutation/webscraper"
	"fmt"
)

func main() {
	url := "https://www.google.com/search?q=blood&start="
	for i := 1; i < 11; i++ {
		data := webscraper.GetContent(url + fmt.Sprint(i))
		if data.Err != nil {
			fmt.Println("Error:", data.Err)
			return
		}
		fmt.Println("URL: ", data.Url)
		data.ParceURL()
	}
}
