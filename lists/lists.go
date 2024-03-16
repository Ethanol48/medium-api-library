package lists

import (
	"strings"

	"github.com/Ethanol48/medium-api-library/utilities"
	"github.com/gocolly/colly"
)

// TODO: support for returning 'normal' link
type ArticleSimple struct {
	Title   string
	Summary string
	Url     string
}

type list struct {
	SimpleArticles []ArticleSimple
}

// TODO: Investigate network trafic to find other way to get more data
// Gets a maximum of 10 artciles in a list in medium.com
func GetArticlesInList(link string) list {
	var l list

	c := colly.NewCollector()
	var baseUrl string

	c.OnRequest(func(r *colly.Request) {

		baseUrl = r.URL.Host

	})

	c.OnHTML("article", func(h *colly.HTMLElement) {

		// The Url for the article is the second parent of the h2 element
		var artUrl *colly.HTMLElement
		var url string

		// find h2 element
		artUrlSelection := h.DOM.Find("h2").First()
		artUrl = &colly.HTMLElement{}
		artUrl.DOM = artUrlSelection
		artUrl.DOM = artUrl.DOM.Parent().Parent().Parent()

		href, exists := artUrl.DOM.Find("a").First().Attr("href")
		if !exists {
			url = "The link could not be found"
		} else {
			path := strings.Split(href, "?source")[0]
			url = "https://" + baseUrl + path

		}

		// get summary if it exists
		summ := artUrl.DOM.Find("p").First().Text()

		l.SimpleArticles = append(l.SimpleArticles, ArticleSimple{
			Title:   utilities.TrimMoreThanOneSpace(h.ChildText("h2")),
			Summary: utilities.TrimMoreThanOneSpace(summ),
			Url:     url,
		})
	})

	c.Visit(link)

	c.Wait()

	return l
}

func GetListOfUser(link string) /* []list */ {
	// TODO
}
