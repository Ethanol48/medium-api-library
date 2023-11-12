package main

import (
	"os"

	"github.com/Ethanol48/medium-api/article"
)

// var selector string = `#root > div:nth-child(1) > div:nth-child(5) > div:nth-child(2) > div:nth-child(3) > article:nth-child(2) > div:nth-child(1) > div:nth-child(1) > section:nth-child(2) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1)`

func main() {

	art := article.GetArticle("test")

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	art.ToHtmlFile(wd + "/Test.html")

}
