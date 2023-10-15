package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PokemonProduct struct {
	Url   string
	Image string
	Name  string
	Price string
}

var pokemonProducts []PokemonProduct

func main() {
	url := "https://scrapeme.live/shop/"

	c := colly.NewCollector()

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.Url = e.ChildAttr("a", "href")
		pokemonProduct.Image = e.ChildAttr("img", "src")
		pokemonProduct.Name = e.ChildText("h2")
		pokemonProduct.Price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.Visit(url)

	// Create csv file.
	file, err := os.Create("pokemon_list.csv")
	if err != nil {
		log.Fatal(`Failed to create "pokemon_list.csv".`)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"url", "image", "name", "price",
	}
	writer.Write(headers)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.Url,
			pokemonProduct.Image,
			pokemonProduct.Name,
			pokemonProduct.Price,
		}
		writer.Write(record)
	}
	defer writer.Flush()
}
