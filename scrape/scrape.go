package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
	url          string = "https://www.tengoods.com.tw"
	page         string = "https://www.tengoods.com.tw/Home/Product"
	goods        []Good
	imgDirectory = "./output/img"
)

func downloadImage(url, imgPath string) error {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("http.Get()錯誤：")
		return err
	}
	defer response.Body.Close()
	file, err := os.Create(imgPath)
	if err != nil {
		fmt.Printf("os.Create()錯誤：")
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("io.Copy()錯誤：")
		return err
	}
	return nil
}

func init() {
	if err := os.MkdirAll(imgDirectory, os.ModePerm); err != nil {
		log.Fatalf("創建圖片目錄失敗：%v", err)
	}
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

		// lock := sync.Mutex{}
		// lock.Lock()
		goods = append(goods, good)
		// lock.Unlock()

		imgPath := filepath.Join(imgDirectory, good.Name+".jpg")
		err := downloadImage(good.Image, imgPath)
		if err != nil {
			fmt.Printf("%v \n", err)
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	//TODO: use gorutine, channel to fix it.
	c.Visit(page)
	c.Visit(page + "/2")
	c.Visit(page + "/3")
	c.Visit(page + "/4")
	c.Visit(page + "/5")

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
