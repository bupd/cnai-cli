package utils

import "strings"

// RemoveHttps removes the "https://" prefix from a string if it exists.
func RemoveHttps(url string) string {
	// Check if the string starts with "https://"
	if strings.HasPrefix(url, "https://") {
		// Remove the "https://" part by slicing the string
		return url[len("https://"):]
	}
	if strings.HasPrefix(url, "http://") {
		// Remove the "https://" part by slicing the string
		return url[len("http://"):]
	}

	// If it doesn't start with "https://", return the original string
	return url
}
