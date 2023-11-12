package article

import (
	"github.com/Ethanol48/medium-api/elements"
	"github.com/Ethanol48/medium-api/utilities"

	"github.com/gocolly/colly"
)

func GetArticle(link string) elements.Article {

	c := colly.NewCollector()
	// c.WithTransport(t)

	var art elements.Article

	c.OnHTML("section > div > div:nth-child(2) > div > div", func(h *colly.HTMLElement) {

		var header *colly.HTMLElement

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

		headerSelection := h.DOM.Find("div").First()
		header = &colly.HTMLElement{}
		header.DOM = headerSelection

		art.Title = header.ChildText(`h1[data-testid="storyTitle"]`)
		art.ReadTime = header.ChildText(`span[data-testid="storyReadTime"]`)
		art.Publisehd = header.ChildText(`span[data-testid="storyPublishDate"]`)

	})

	if link == "test" {
		go utilities.SpinUp("testing")
		c.Visit("http://localhost:8080/")
	} else {
		c.Visit(link)
	}
	c.Wait()

	return art

}
