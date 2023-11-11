package main

import (
	"os"

	"github.com/Ethanol48/medium-api/article"
	"github.com/Ethanol48/medium-api/utilities"

	"github.com/gocolly/colly"
)

// var selector string = `#root > div:nth-child(1) > div:nth-child(5) > div:nth-child(2) > div:nth-child(3) > article:nth-child(2) > div:nth-child(1) > div:nth-child(1) > section:nth-child(2) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1)`

func main() {

	go utilities.SpinUp("testing")

	c := colly.NewCollector()
	// c.WithTransport(t)

	c.OnHTML("section > div > div:nth-child(2) > div > div", func(h *colly.HTMLElement) {

		art := article.GetArticle(h)

		// fmt.Println("hello??")
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		art.ToMarkdownFile(wd + "/Example.md")

	})

	// c.Visit("https://medium.com/coinsbench/uniswap-whats-new-5e41307c97a2")

	c.Visit("http://localhost:8080/")
	c.Wait()

}
