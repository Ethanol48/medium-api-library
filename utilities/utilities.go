package utilities

import (
	"log"
	"net/http"
	"strings"
)

func StringInArray(s string, array []string) bool {

	for i := 0; i < len(array); i++ {
		if s == array[i] {
			return true
		}
	}

	return false
}

func TrimMoreThanOneSpace(s string) string {
	str := strings.ReplaceAll(s, "  ", "")
	str = strings.ReplaceAll(str, "\n", " ")

	return str
}

// creates an http server for testing purposes
func SpinUp(path string) {
	// Set the directory to serve.
	fs := http.FileServer(http.Dir(path))

	// Handle all requests by serving a file of the same name.
	http.Handle("/", fs)

	// Start the server on port 8080 and handle any errors.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ExtractLinkImg(str string) (string, string) {

	str = strings.TrimSpace(str)

	strs := strings.Split(str, " ")

	s := TrimMoreThanOneSpace(strs[0])

	// Find the last index of "/"
	lastIndex := strings.LastIndex(s, "/")

	// Split the string around the last index of "/"
	base := s[:lastIndex]
	id := s[lastIndex+1:]

	return base, id
}
