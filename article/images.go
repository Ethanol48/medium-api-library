package article

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func GetImagesFromMarkdown(path string) {

	// Read the contents of the Markdown file
	markdownContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// Regular expression to match Markdown image links `![text](url)`
	imgRegex := regexp.MustCompile(`!\[.*\]\(([^)]+)\)`)

	imgMatches := imgRegex.FindAllStringSubmatch(string(markdownContent), -1)
	// Iterate over the matches and print the image URLs
	for _, match := range imgMatches {
		if len(match) >= 2 {
			imageUrl := match[1]
			fmt.Println("Found image URL:", imageUrl)
		}
	}

}

// this function downloads an image from a URL and saves it to a file
func DownloadImageFromUrl(url string) {
	// Make an HTTP GET request to fetch the image
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code indicates success (e.g., 200 OK)
	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed with status code:", response.StatusCode)
		return
	}

	// Create a new file to save the downloaded image
	file, err := os.Create("downloaded_image.jpg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Copy the image data from the HTTP response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error copying image data:", err)
		return
	}

	fmt.Println("Image downloaded and saved as downloaded_image.jpg")
}
