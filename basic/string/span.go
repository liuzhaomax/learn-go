package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	html := `<div><span class="text-only" data-eleid="3"><span class="text-only">Please Generate a flappy bird game with </span><span class="text-only text-with-abbreviation text-with-abbreviation-bottomline windows-bottomline">flask</span><span class="text-only"> web framework in python</span></span></div>`

	// Compile the regular expression for extracting text within <span>
	re := regexp.MustCompile(`<span.*?>(.*?)<\/span>`)

	// Find all matches
	matches := re.FindAllStringSubmatch(html, -1)
	if matches != nil {
		var result strings.Builder
		for _, match := range matches {
			// match[1] contains the text inside the <span> tag
			result.WriteString(strings.TrimSpace(match[1]))
			result.WriteString(" ") // Add a space between parts
		}
		// Convert Builder to string, trim extra spaces, and print
		finalResult := strings.TrimSpace(result.String())
		fmt.Println(finalResult)
		// 有问题，还是带着一个span
	} else {
		fmt.Println("No matches found")
	}
}
