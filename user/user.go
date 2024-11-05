package user

import (
	"strings"

	"github.com/gocolly/colly"
)

// TODO: add support for obtaining top 10 articles of user

type UserMetadata struct {
  Name           string /* `json:Name` */
  About          string /* `json:About` */
  Desc           string /* `json:Desc` */
  Followers      string /* `json:Followers` */
  Following      string /* `json:Following` */
  twitterHandle  string /* `json:twitterHandle` */
  mastodonHandle string /* `json:mastodonHandle` */
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

	// Description
	c.OnHTML("body > div > div > div > div > div > div > div > div > div > div > div > div > p > span", func(h *colly.HTMLElement) {
		u.Desc = strings.TrimSpace(h.Text)
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

	// Mastodon
	c.OnHTML(`a[href^="https://me.dm"]`, func(h *colly.HTMLElement) {

		tmp := h.DOM.AttrOr("href", "Not Found")

		u.mastodonHandle = strings.Split(tmp, "?source=")[0]

	})

	// About
	c.OnHTML("h4", func(h *colly.HTMLElement) {

		var about *colly.HTMLElement
		aboutSelection := h.DOM.Parent().Parent().Find("div").First().Find("p")
		about = &colly.HTMLElement{}
		about.DOM = aboutSelection

		// TODO: review for returning html for including links
		// test, err := desc.DOM.Html()
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("test: %v\n", test)

		aboutText := about.DOM.Text()

		// this is for the case of scrapping your own profile
		aboutText, _ = strings.CutSuffix(aboutText, "Edit")

		u.About = strings.TrimSpace(aboutText)
	})

	c.Visit(link)

	return u

}
