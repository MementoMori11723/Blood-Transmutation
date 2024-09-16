package main

import (
	"Blood-Transmutation/server"
	"Blood-Transmutation/webscraper"
	"fmt"
)

var (
	urlArr = []string{}
)

func main() {
  webscraper.RunScrape(&urlArr)
	go server.StartServer(&urlArr)
	fmt.Scanln()
}
