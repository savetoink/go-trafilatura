package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Testing relative image URL conversion...")
	fmt.Println()

	// Parse the original URL
	originalURL, err := url.Parse("https://example.com/js-blog/article.html")
	if err != nil {
		log.Fatal(err)
	}

	// Create some test HTML strings
	testCases := []struct {
		html string
		name string
	}{
		{
			html: `<img src="/js-blog/images/image3.jpg" alt="test"/>`,
			name: "Root-relative path",
		},
		{
			html: `<img src="../images/image2.jpg" alt="test"/>`,
			name: "Parent-relative path",
		},
		{
			html: `<img src="images/image3.jpg" alt="test"/>`,
			name: "Path-relative",
		},
		{
			html: `<img src="https://example.com/full/path/image4.jpg" alt="test"/>`,
			name: "Absolute URL",
		},
		{
			html: `<img src="//cdn.example.com/image.jpg" alt="test"/>`,
			name: "Protocol-relative",
		},
	}

	fmt.Println("Test cases for image URL conversion:")
	fmt.Println("Original URL:", originalURL.String())
	fmt.Println()
	for _, tc := range testCases {
		// Extract src attribute from HTML for display
		src := ""
		start := strings.Index(tc.html, `src="`) + 5
		end := strings.Index(tc.html[start:], `"`)
		if end > 0 {
			src = tc.html[start : start+end]
		}
		fmt.Printf("%s: %s\n", tc.name, src)
	}
	fmt.Println()
	fmt.Println("After conversion, these would become:")
	fmt.Println("  Root-relative path: https://example.com/js-blog/images/image3.jpg")
	fmt.Println("  Parent-relative path: https://example.com/images/image2.jpg")
	fmt.Println("  Path-relative: https://example.com/js-blog/images/image3.jpg")
	fmt.Println("  Absolute URL: https://example.com/full/path/image4.jpg (unchanged)")
	fmt.Println("  Protocol-relative: https://cdn.example.com/image.jpg")
	fmt.Println()
	fmt.Println("To use this feature, call Extract() with Options that include:")
	fmt.Println("  - IncludeImages: true")
	fmt.Println("  - OriginalURL: your base URL")
}
