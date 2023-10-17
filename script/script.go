package main

import (
	f "fmt"
	"log"

	s "github.com/bitfield/script"
)

func main() {
	data, err := s.File("../pokemon_list.csv").String()
	if err != nil {
		log.Fatal("Read file failed.")
	}
	f.Println(data)
}
