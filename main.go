package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
	adv "github.com/rcacunar/air-scrapper/utils"
)

func main() {
	//dataFile
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	//arrays

	city := []string{}
	state := []string{}

	//fileWrite
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Instantiate default collector
	c := colly.NewCollector()

	//csv header
	writer.Write([]string{
		"ciudad",
		"estado del aire",
	})

	// Before making a request put the URL with
	// the key of "url" into the context of the request
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	// After making a request get "url" from
	// the context of the request
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Ctx.Get("url"))
	})

	println(" ## se genera csv ##")
	//obtain city from the main container (city)
	c.OnHTML("a.container-city", func(e *colly.HTMLElement) {
		//iterate for every li element
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {

			if el.Text != " Medidas y recomendaciones" {

				city = append(city, el.Text)
				adv.SaveC = append(adv.SaveC, el.Text)

			}

		})
		//iterates for every span.label (state)
		e.ForEach("span.label", func(_ int, sp *colly.HTMLElement) {

			if sp.Text != " " {

				state = append(state, sp.Text)
			}

		})
		//fmt.Println(" ")
		//fmt.Println(city[len(city)-1], ",", state[len(state)-1])
		//write the result
		writer.Write([]string{
			city[len(city)-1],
			state[len(state)-1],
		})

	})

	c.Visit("https://airechile.mma.gob.cl/")

	fmt.Println(adv.SaveC)
	adv.Save()

}
