package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Good struct {
	Url   string
	Image string
	Name  string
	Price string
	Store string
	Left  string
}

var (
	url   string = "https://www.tengoods.com.tw/"
	goods []Good
)

func main() {
	c := colly.NewCollector()
	c.OnHTML(".home-proitem-container", func(e *colly.HTMLElement) {
		good := Good{}
		good.Url = e.ChildAttr("a", "href")
		good.Image = e.ChildAttr("img", "data-src")
		good.Name = e.ChildText(".home-proitem-item.home-proitem-pro")
		good.Store = e.ChildText(".home-proitem-item.home-proitem-shop")
		good.Price = e.ChildText(".home-proitem-item.home-proitem-price")
		good.Left = e.ChildText(".home-proitem-item.home-proitem-qty")
		goods = append(goods, good)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	c.Visit(url)

	file, err := os.Create("TenGoods_HomePageItems.csv")
	if err != nil {
		log.Fatal(`Create csv file failed.`)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"url", "image", "name", "store", "price", "left",
	}
	writer.Write(headers)

	for _, good := range goods {
		record := []string{
			good.Url,
			good.Image,
			good.Name,
			good.Store,
			good.Price,
			good.Left,
		}
		writer.Write(record)
	}
	defer writer.Flush()
}
