package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Converts markdown to HTML
func convertMarkdownToHTML(markdown string) string {
	lines := strings.Split(markdown, "\n")
	var htmlOutput []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Handle headers
		if isHeader(line) {
			htmlOutput = append(htmlOutput, convertHeader(line))
			continue
		}

		// Handle links
		line = convertLinks(line)

		htmlOutput = append(htmlOutput, fmt.Sprintf("<p>%s</p>", line))
	}

	return strings.Join(htmlOutput, "\n")
}

// Check if the line is a header
func isHeader(line string) bool {
	return regexp.MustCompile(`^#{1,6}\s`).MatchString(line) // Ensures max 6 '#'
}

// Convert a markdown header to HTML header
func convertHeader(line string) string {
	headerRegex := regexp.MustCompile(`^(#{1,6})\s*(.*)`)
	match := headerRegex.FindStringSubmatch(line)
	if len(match) == 0 {
		return fmt.Sprintf("<p>%s</p>", line) // Fallback if no match or more than 6 '#'
	}
	headerLevel := len(match[1])
	headerContent := match[2]

	headerContent = convertLinks(headerContent)
	return fmt.Sprintf("<h%d>%s</h%d>", headerLevel, headerContent, headerLevel)
}

// Convert markdown links to HTML links
func convertLinks(line string) string {
	linkRegex := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	return linkRegex.ReplaceAllString(line, `<a href="$2">$1</a>`)
}

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Markdown input file")
	outputFile := flag.String("output", "", "HTML output file (optional)")

	flag.Parse()

	var markdown string

	// Read from input file if specified
	if *inputFile != "" {
		content, err := os.ReadFile(*inputFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		markdown = string(content)
	} else {
		// Read from stdin if no input file
		fmt.Println("Enter markdown text (Ctrl+D to finish):")
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
		markdown = string(input)
	}

	// Convert markdown to HTML
	htmlOutput := convertMarkdownToHTML(markdown)

	// Write to output file if specified
	if *outputFile != "" {
		err := os.WriteFile(*outputFile, []byte(htmlOutput), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			os.Exit(1)
		}
		fmt.Printf("HTML output written to %s\n", *outputFile)
	} else {
		// Write to 'output.html' and print its contents
		outputFile = new(string)
		*outputFile = "output.html"
		err := os.WriteFile(*outputFile, []byte(htmlOutput), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			os.Exit(1)
		}
		fmt.Printf("HTML output written to %s\n", *outputFile)

		// Read and print the output file contents
		content, err := os.ReadFile(*outputFile)
		if err != nil {
			fmt.Println("Error reading output file:", err)
			os.Exit(1)
		}
		fmt.Println(string(content))
	}
}
