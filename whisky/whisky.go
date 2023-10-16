package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

//	type Whisky struct {
//		Name  string
//		Price string
//		Img   string
//		Abv   string
//	}
type Catgory struct {
	Country string
	Name    string
	Img     string
	Url     string
}

var catgorys []Catgory

func main() {
	url := "https://www.609.com.tw/Country"
	c := colly.NewCollector()
	c.OnHTML(".card", func(e *colly.HTMLElement) {
		catgory := Catgory{}

		catgory.Country = e.ChildText("h5") //TODO: 帶入國家分類，應該是<h5>
		catgory.Name = e.ChildText("h6")
		catgory.Img = e.ChildAttr("img", "src")
		catgory.Url = e.ChildAttr("a", "href")
		catgorys = append(catgorys, catgory)

	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visisting", r.URL)
	})
	c.OnScraped(func(r *colly.Response) {
		if len(catgorys) == 0 {
			log.Println("No data found on the page.")
		} else {
			fmt.Println(catgorys)
		}
		c.OnError(func(r *colly.Response, err error) {
			log.Println("Request error:", err)
		})
	})
	c.Visit(url)

	fmt.Println(catgorys)
}
