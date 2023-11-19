package article

import (
	"strings"

	"github.com/Ethanol48/medium-api-library/elements"
	"github.com/Ethanol48/medium-api-library/utilities"

	"github.com/gocolly/colly"
)

// whitelist of elements to be included in the article
var permElems []string = []string{
	"h1", "h2", "h3", "h4", "ul", "ol", "p", "img", "blockquote", "pre",
}

func GetArticle(link string) elements.Article {

	c := colly.NewCollector()
	// c.WithTransport(t)

	var art elements.Article

	c.OnHTML("section > div > div:nth-child(2) > div > div", func(h *colly.HTMLElement) {

		h.ForEach("*:not(:first-child)", func(_ int, e *colly.HTMLElement) {

			// don't include "Follow" button
			if utilities.StringInArray(e.Name, permElems) {

				var result elements.Element

				if utilities.TrimMoreThanOneSpace(e.Text) == "Follow" {
					result = &elements.PlaceHolder{
						Name: "",
						Elem: *e,
					}

				} else {
					result = elements.ExtractDataArticle(e)
				}

				if result.GetName() != "" {
					art.Content = append(art.Content, result)
				}
			}

			// fmt.Printf("\nPrinting number: %v\n", i)
			// fmt.Printf("e.Name: %v\n", e.Name)

		})

		var header *colly.HTMLElement
		headerSelection := h.DOM.Find("div").First()
		header = &colly.HTMLElement{}
		header.DOM = headerSelection

		art.Title = header.ChildText(`h1[data-testid="storyTitle"]`)
		art.ReadTime = header.ChildText(`span[data-testid="storyReadTime"]`)
		art.Published = header.ChildText(`span[data-testid="storyPublishDate"]`)

	})

	// take all anchor tags where href begins with ...
	c.OnHTML(`a[href^="https://medium.com/tag"]`, func(h *colly.HTMLElement) {

		// clean the link and eliminate the sufic of the link to get the tag
		str, _ := strings.CutPrefix(strings.Split(h.Attr("href"), "?source=")[0], "https://medium.com/tag/")

		art.Tags = append(art.Tags, str)

	})

	// eliminate this and create testfile
	if link == "test" {
		go utilities.SpinUp("testing")
		c.Visit("http://localhost:8080/article")
	} else {
		c.Visit(link)
	}
	c.Wait()

	return art

}
