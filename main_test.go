package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

// Helper function to simplify testing
func assertEqual(t *testing.T, got, expected string) {
	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

// Test case for heading conversion
func TestConvertHeading(t *testing.T) {
	markdown := "# Heading 1"
	expectedHTML := "<h1>Heading 1</h1>"
	got := convertMarkdownToHTML(markdown)
	assertEqual(t, got, expectedHTML)
}

// Test case for multiple headers
func TestMultipleHeaders(t *testing.T) {
	markdown := "## Heading 2\n### Heading 3"
	expectedHTML := "<h2>Heading 2</h2>\n<h3>Heading 3</h3>"
	got := convertMarkdownToHTML(markdown)
	assertEqual(t, got, expectedHTML)
}

// Test case for paragraph conversion
func TestConvertParagraph(t *testing.T) {
	markdown := "This is a paragraph."
	expectedHTML := "<p>This is a paragraph.</p>"
	got := convertMarkdownToHTML(markdown)
	assertEqual(t, got, expectedHTML)
}

// Test case for link conversion
func TestConvertLink(t *testing.T) {
	markdown := "[Link](https://example.com)"
	expectedHTML := `<p><a href="https://example.com">Link</a></p>`
	got := convertMarkdownToHTML(markdown)
	assertEqual(t, got, expectedHTML)
}

// Test case for complex input (header, paragraph, and link)
func TestComplexInput(t *testing.T) {
	markdown := "# Heading 1\n\nThis is a [link](https://example.com)."
	expectedHTML := "<h1>Heading 1</h1>\n<p>This is a <a href=\"https://example.com\">link</a>.</p>"
	got := convertMarkdownToHTML(markdown)
	assertEqual(t, got, expectedHTML)
}

// Test case for blank lines (which should be ignored)
func TestBlankLines(t *testing.T) {
	markdown := "\n# Heading 1\n\n\nThis is a paragraph.\n\n"
	expectedHTML := "<h1>Heading 1</h1>\n<p>This is a paragraph.</p>"
	got := convertMarkdownToHTML(markdown)
	assertEqual(t, got, expectedHTML)
}

// Test case for fallback in header conversion (no match or too many '#')
func TestConvertHeaderFallback(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No match case",
			input:    "No header",
			expected: "<p>No header</p>",
		},
		{
			name:     "Too many # case",
			input:    "####### Too many hashes",
			expected: "<p>####### Too many hashes</p>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := convertMarkdownToHTML(tt.input)
			if output != tt.expected {
				t.Errorf("got %s, want %s", output, tt.expected)
			}
		})
	}
}

// Test case for stdin input (no input file provided)
func TestMain_NoInputFile(t *testing.T) {
	// Backup and restore os.Stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// Create a simulated stdin input
	input := "### Heading 3"
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write to stdin
	go func() {
		defer w.Close()
		io.WriteString(w, input)
	}()

	// Reinitialize flag set to avoid conflicts
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Simulate no command-line flags (stdin input)
	os.Args = []string{"cmd"}

	main()

	// Check the output file
	outputData, err := os.ReadFile("output.html")
	if err != nil {
		t.Fatalf("error reading output file: %v", err)
	}

	expectedOutput := "<h3>Heading 3</h3>"
	if !strings.Contains(string(outputData), expectedOutput) {
		t.Errorf("expected output to contain %s, got %s", expectedOutput, string(outputData))
	}
}

// Test case for input file
func TestMain_WithInputFile(t *testing.T) {
	// Create a temporary markdown file
	inputMarkdown := "## Heading 2"
	inputFile, err := os.CreateTemp("", "test_input.md")
	if err != nil {
		t.Fatalf("error creating temp input file: %v", err)
	}
	defer os.Remove(inputFile.Name()) // Clean up

	if _, err := inputFile.WriteString(inputMarkdown); err != nil {
		t.Fatalf("error writing to temp input file: %v", err)
	}
	inputFile.Close()

	// Reinitialize flag set to avoid conflicts
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Simulate command-line argument for input file
	os.Args = []string{"cmd", "-input", inputFile.Name()}

	main()

	// Check the output file
	outputData, err := os.ReadFile("output.html")
	if err != nil {
		t.Fatalf("error reading output file: %v", err)
	}

	expectedOutput := "<h2>Heading 2</h2>"
	if string(outputData) != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, string(outputData))
	}
}
