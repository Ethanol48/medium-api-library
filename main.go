package main

import (

	// "os"

	"fmt"

	"github.com/Ethanol48/medium-api-library/lists"
	"github.com/Ethanol48/medium-api-library/utilities"
)

// var selector string = `#root > div:nth-child(1) > div:nth-child(5) > div:nth-child(2) > div:nth-child(3) > article:nth-child(2) > div:nth-child(1) > div:nth-child(1) > section:nth-child(2) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1)`

func main() {

	// go utilities.SpinUp("testing")

	// lists.GetArticlesInList("http://localhost:8080/list")
	lists := lists.GetArticlesInList("test")

	lists[0].Summary = utilities.TrimMoreThanOneSpace(lists[0].Summary)
	lists[0].Title = utilities.TrimMoreThanOneSpace(lists[0].Title)

	fmt.Printf("lists: %#v\n", lists[0])

	// time.Sleep(12)

	// art := article.GetArticle("https://medium.com/@champudelimon/testing-medium-markdown-f0eb5a7054bc")
	// fmt.Printf("art.ToMarkdown(): \n%v\n", art.ToMarkdown())

	// lists.GetArticlesInList("https://medium.com/@ethan-rouimi/list/solidity-content-24ad19a2c23d")

	// art := article.GetArticle("https://medium.com/@ethan-rouimi/paypal-usd-et-les-horreurs-des-cbdcs-861f2339304b")

	// wd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("art.Title: %v\n", art.Title)

	// // art.ToMarkdownFile(wd + "/output/article.md")
	// article.GetImagesFromMarkdown(wd + "/output/article.md")

	// article.DownloadImage("https://miro.medium.com/v2/resize:fit:640/0*YwKmtuNmHrw5FVVU")

}
