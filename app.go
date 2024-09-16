package main

import (
	"Blood-Transmutation/config"
	"Blood-Transmutation/server"
	"Blood-Transmutation/webscraper"
	"fmt"
	"strings"
)

var (
	urlArr = []string{}
)

func main() {
	for i := 1; i < 11; i++ {
		data := webscraper.GetContent(config.Url + fmt.Sprint(strings.ReplaceAll(config.Questions[i], " ", "+")) + config.Pages + fmt.Sprint(i))
		if data.Err != nil {
			fmt.Println("Error:", data.Err)
			return
		}
		data.ParceURL(&urlArr)
	}
	go server.StartServer(&urlArr)
	fmt.Scanln()
}
