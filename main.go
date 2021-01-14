package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

// "io/ioutil"
// "log"

//Fact struct
type Fact struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func writeJSON(data []Fact, fileName string) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}
	_ = ioutil.WriteFile(fileName, file, 0644)
}

func main() {
	var allFacts []Fact

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)
	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		// strconv.Atoi convert string to int
		// strconv.Itoa convert int to string
		factID := element.Attr("id")
		factDesc := element.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}
		// fmt.Println(fact)
		allFacts = append(allFacts, fact)
	})
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL)
	})
	collector.Visit("https://www.factretriever.com/popcorn-facts")

	// fmt.Println(allFacts)
	writeJSON(allFacts, "allFacts.json")

}
