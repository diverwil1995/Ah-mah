package main

import (
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Good struct {
	Url   string
	Image string
	Name  string
	Price string
}

var (
	url string = "https://www.tengoods.com.tw/"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("div.home-proitem-container", func(e *colly.HTMLElement) {

	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	c.Visit(url)

	file, err := os.Create("TenGoods_NewItems.csv")
	if err != nil {
		log.Fatal(`Create csv file failed.`)
	}
	defer file.Close()
}
