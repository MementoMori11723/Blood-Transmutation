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
    res := data.ParceHTML()
    urls := res.([]string)
    *urlArr = append(*urlArr, urls...)
	}
  for _, url := range *urlArr {
    fmt.Println(url)
    data := GetContent(url)
    if data.Err != nil {
      fmt.Println("Error:", data.Err)
      return
    }
    res := data.ParceHTML().(Scrape)
    res.Summarize()
  }
}
