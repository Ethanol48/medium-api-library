package user

import (
	"strings"

	"github.com/gocolly/colly"
)

// TODO: add support for obtaining top 10 articles of user

type UserMetadata struct {
	Name           string
	Desc           string
	Followers      string
	Following      string
	twitterHandle  string
	mastodonHandle string
}

func GetUserMetadata(link string) UserMetadata {

	var u UserMetadata

	c := colly.NewCollector()

	// Username
	c.OnHTML(".pw-author-name", func(h *colly.HTMLElement) {
		u.Name = h.Text
	})

	// Followers
	c.OnHTML("h4 > span > a", func(h *colly.HTMLElement) {

		u.Followers = strings.Split(h.Text, " ")[0]

	})

	// Following
	c.OnHTML("h4 > a", func(h *colly.HTMLElement) {

		u.Following = strings.Split(h.Text, " ")[0]

	})

	// Twitter
	c.OnHTML(`a[href^="https://twitter.com"]`, func(h *colly.HTMLElement) {

		tmp := h.DOM.AttrOr("href", "Not Found")

		u.twitterHandle = strings.Split(tmp, "?source=")[0]

	})

	// Twitter
	c.OnHTML(`a[href^="https://me.dm"]`, func(h *colly.HTMLElement) {

		tmp := h.DOM.AttrOr("href", "Not Found")

		u.mastodonHandle = strings.Split(tmp, "?source=")[0]

	})

	// Description
	c.OnHTML("h4", func(h *colly.HTMLElement) {

		var desc *colly.HTMLElement
		descSelection := h.DOM.Parent().Parent().Find("div").First().Find("p")
		desc = &colly.HTMLElement{}
		desc.DOM = descSelection

		// TODO: review for returning html for including links
		// test, err := desc.DOM.Html()
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("test: %v\n", test)

		descText := desc.DOM.Text()

		// this is for the case of scrapping your own profile
		descText, _ = strings.CutSuffix(descText, "Edit")

		u.Desc = descText
	})

	c.Visit(link)

	return u

}
