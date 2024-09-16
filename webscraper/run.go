package webscraper

import (
	"Blood-Transmutation/config"
	"fmt"
	"strings"
)

func RunScrape(urlArr *[]string) {
	for i := 1; i < 11; i++ {
		data := GetContent(config.Url + fmt.Sprint(strings.ReplaceAll(config.Questions[i], " ", "+")) + config.Pages + fmt.Sprint(i))
		if data.Err != nil {
			fmt.Println("Error:", data.Err)
			return
		}
		data.ParceURL(&urlArr)
	}
}
