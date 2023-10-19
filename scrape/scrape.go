package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gocolly/colly"
)

type Good struct {
	Url   string
	Image string //TODO: Download image.
	Name  string
	Price string
	Store string
	Left  string
}

var (
	url          string = "https://www.tengoods.com.tw"
	goods        []Good
	imgDirectory = "./output/img"
)

func downloadImage(url, imgPath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	file, err := os.Create(imgPath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	c := colly.NewCollector()
	c.OnHTML(".home-proitem-container", func(e *colly.HTMLElement) {
		good := Good{}
		good.Url = e.ChildAttr("a", "href")
		good.Image = url + e.ChildAttr("img", "data-src")
		good.Name = e.ChildText(".home-proitem-item.home-proitem-pro")
		good.Store = e.ChildText(".home-proitem-item.home-proitem-shop")
		good.Price = e.ChildText(".home-proitem-item.home-proitem-price")
		good.Left = e.ChildText(".home-proitem-item.home-proitem-qty")
		goods = append(goods, good)

		imgPath := filepath.Join(imgDirectory, good.Name+".jpg")
		err := downloadImage(good.Image, imgPath)
		if err != nil {
			log.Fatal(err)
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	c.Visit(url)

	file, err := os.Create("./output/TenGoods_HomePageItems.csv")
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
