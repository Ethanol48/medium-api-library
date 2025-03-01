package utilities

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// [x] get images from markdown
// [x] download images to files
// [ ] change markdown

func GetImagesFromMarkdown(path string) []string {

	var imageUrls []string

	// Read the contents of the Markdown file

	markdownContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return imageUrls
	}
	// Regular expression to match Markdown image links `![text](url)`
	imgRegex := regexp.MustCompile(`!\[.*\]\(.*\)`)

	imgMatches := imgRegex.FindAllStringSubmatch(string(markdownContent), -1)

	for _, match := range imgMatches {
		link := match[0]
		imageUrls = append(imageUrls, link)
	}

	return imageUrls

}

// Takes a markdown string and changes the images urls tags with new ones
func ChangeImageLinksInMarkdown(markdownContent string, oldImageTags []string, newLinks []string) (string, error) {
	// check arrays have same length

	same := len(oldImageTags) == len(newLinks)
	if !same {
		return "", errors.New("both arrays must be the same length")
	}

	// Iterate over the old image tags and their corresponding new links
	for i, oldImageTag := range oldImageTags {
		// Extract the image description from the old image tag
		re := regexp.MustCompile(`!\[(.*?)\]`)
		matches := re.FindStringSubmatch(oldImageTag)
		if len(matches) < 2 { // Make sure there is a match for the description
			fmt.Println("No match found for image description in:", oldImageTag)
			continue
		}
		imageDescription := matches[1]

		// Construct the new image tag with the old description and new link
		newImageTag := fmt.Sprintf("![%s](%s)", imageDescription, newLinks[i])

		// Replace the old image tag with the new one in the markdown content
		markdownContent = regexp.MustCompile(regexp.QuoteMeta(oldImageTag)).ReplaceAllString(markdownContent, newImageTag)
	}

	return markdownContent, nil

}

func DownloadAndSaveImageToFs(url string, filename string) error {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("There was an error getting the image: ", err)
		return err
	}

	defer response.Body.Close()

	// Check if the response status code indicates success (e.g., 200 OK)
	if response.StatusCode != http.StatusOK {
		err := fmt.Sprintf("HTTP code was different than 200: %d", response.StatusCode)
		return errors.New(err)
	}

	data := response.Body

	// Create a new file to save the downloaded image
	image, err := os.Create(filename)
	if err != nil {
		fmt.Println("There was an error creating the file: ", err)
		return err
	}
	defer image.Close()

	// Copy the image data from the HTTP response body to the file
	_, err = io.Copy(image, data)
	if err != nil {
		fmt.Println("There was an error writing the file: ", err)
		return err
	}

	return nil
}

// Takes the image tag string of markdown *![](link)* and extract the *link*
func ExtractUrlFromImageTag(imageTag string) string {

	urlRegex := regexp.MustCompile(`!\[.*?\]\(([^)]+)\)`)

	// Find all matches and iterate over them
	matches := urlRegex.FindAllStringSubmatch(imageTag, -1)

	if len(matches) > 0 {
		return matches[0][1]

	} else {
		return ""
	}
}

/*
path: path of the Markdown file

This function combines the functionability of multiple functions, it looks for every
image link in the markdown, it downloads the image in the same path as the markdown file
and modifies the file images tag to instead of point to the new downloaded images instead of an url
*/
func DownloadImagesAndModifyFile(path string) error {

	dir := filepath.Dir(path) // Get the directory of the markdown file

	// raw images links
	imageTags := GetImagesFromMarkdown(path)

	// link of image tag
	cleanUrls := make([]string, len(imageTags))

	for i := 0; i < len(imageTags); i++ {
		cleanUrls[i] = ExtractUrlFromImageTag(imageTags[i])
	}

	// download them with new names
	newUrls := make([]string, len(imageTags))
	for i := 0; i < len(imageTags); i++ {
		filename := fmt.Sprintf("image-%d.png", i)
		newUrls[i] = filename

		filenameWithPath := filepath.Join(dir, filename) // Use filepath.Join to construct file path
		fmt.Println("filename with path: ", filenameWithPath)
		err := DownloadAndSaveImageToFs(cleanUrls[i], filenameWithPath)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	// read markdown text
	markdownContent, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// call ChangeLinksInMarkdown to get new Markdown text
	newMarkdownContent, err := ChangeImageLinksInMarkdown(string(markdownContent), imageTags, newUrls)
	if err != nil {
		fmt.Println("Error changing markdownContent:", err)
		return err
	}

	// modify Markdown

	// Create or truncate the file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error while creating or truncating the file:", err)
		return err
	}
	defer file.Close()

	// Write new content to the file
	_, err = file.WriteString(newMarkdownContent)
	if err != nil {
		fmt.Println("Error while writing to the file:", err)
		return err
	}

	return nil
}
