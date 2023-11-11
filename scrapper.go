package main

import (
	"encoding/json"
	"fmt"
	"medium-api/elements"
	"medium-api/utilities"
	"os"

	"github.com/gocolly/colly"
)

// var selector string = `#root > div:nth-child(1) > div:nth-child(5) > div:nth-child(2) > div:nth-child(3) > article:nth-child(2) > div:nth-child(1) > div:nth-child(1) > section:nth-child(2) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1)`

var art elements.Article

func main() {

	go utilities.SpinUp()

	c := colly.NewCollector()
	// c.WithTransport(t)

	c.OnHTML("section > div > div:nth-child(2) > div > div", func(h *colly.HTMLElement) {

		// for _, v := range h.ChildAttr("img", "src") {
		// 	fmt.Println("src: ", v)
		// }

		var header *colly.HTMLElement

		// nodes := h.DOM.Nodes
		h.ForEach("div", func(i int, e *colly.HTMLElement) {

			if i == 0 {
				header = e
			}

			// fmt.Printf("e with number %v: %v\n", i, e.Name)

			// fmt.Printf("\nPrinting number: %v\n", i)
			// fmt.Printf("e.Name: %v\n", e.Name)
		})

		var permElems []string

		// add permite elements
		permElems = append(permElems, "h1")
		permElems = append(permElems, "h2")
		permElems = append(permElems, "h3")
		permElems = append(permElems, "ul")
		permElems = append(permElems, "ol")
		permElems = append(permElems, "p")
		permElems = append(permElems, "li")
		permElems = append(permElems, "img")

		h.ForEach("*:not(:first-child)", func(i int, e *colly.HTMLElement) {

			if utilities.StringInArray(e.Name, permElems) {
				// fmt.Printf("e with number %v: %v\n", i, e.Name)
				// fmt.Printf("%s \n\n", trimMoreThanOneSpace(e.Text))

				result := elements.ExtractDataArticle(e, i)

				if result.GetName() != "" {
					art.Content = append(art.Content, result)
				}
			}

			// fmt.Printf("\nPrinting number: %v\n", i)
			// fmt.Printf("e.Name: %v\n", e.Name)

		})

		art.Title = header.ChildText(`h1[data-testid="storyTitle"]`)
		art.ReadTime = header.ChildText(`span[data-testid="storyReadTime"]`)
		art.Publisehd = header.ChildText(`span[data-testid="storyPublishDate"]`)

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", " ")

		fmt.Printf("\nPrinting article: ")
		enc.Encode(art)
		fmt.Printf("\n")
		// fmt.Printf("before: %#v \n\n", before)

		fmt.Printf("\n\n")

		err := os.WriteFile("Article.md", []byte(art.ToMarkdown()), 0644) // 0644 is a common permission setting allowing reading for everyone and full write access to the owner of the file.
		if err != nil {
			// Handle the error here
			fmt.Println("Error writing to file:", err)
		}

	})

	// c.Visit("https://medium.com/coinsbench/uniswap-whats-new-5e41307c97a2")

	c.Visit("http://localhost:8080/")
	c.Wait()

}
