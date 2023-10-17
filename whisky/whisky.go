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
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
	c.OnHTML(".card a", func(e *colly.HTMLElement) {
		catgory := Catgory{}

		catgory.Country = "" //TODO: 帶入國家分類，應該是<h5>
		catgory.Name = e.ChildText("h6")
		catgory.Img = e.ChildAttr("img", "src")
		catgory.Url = e.ChildAttr("a", "href")
		catgorys = append(catgorys, catgory)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
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
