package article

import (
	"github.com/Ethanol48/medium-api/elements"
	"github.com/Ethanol48/medium-api/utilities"

	"github.com/gocolly/colly"
)

func GetArticle(h *colly.HTMLElement) elements.Article {

	var header *colly.HTMLElement
	var art elements.Article

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

	// fmt.Printf("before: %#v \n\n", before)

	// err := os.WriteFile("Article.md", []byte(art.ToMarkdown()), 0644) // 0644 is a common permission setting allowing reading for everyone and full write access to the owner of the file.
	// if err != nil {
	// 	// Handle the error here
	// 	fmt.Println("Error writing to file:", err)
	// }

	return art

}
