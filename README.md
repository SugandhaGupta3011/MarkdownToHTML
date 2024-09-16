# Markdown to HTML Converter
This Go program converts Markdown text into HTML format. It reads input from a file or standard input, processes the Markdown to generate HTML, and either outputs the result to a file or prints it to standard output.

## Features
Supports conversion of headers (e.g., `# Header 1`, `## Header 2`) into appropriate HTML header tags (`<h1>, <h2>`, etc.).
Converts Markdown links in the format `[link text](URL)` into HTML anchor tags (`<a href="URL">link text</a>`).
Wraps other text in paragraph tags (`<p>`).

## Installation
1. Clone the repository or download the Go file.
2. Ensure that Go is installed on your system.
3. Run the program using the go run command, as shown in the usage examples.
4. Build using ```go build -o markdown_to_html```

# Usage
## Command-line Arguments
-input: Specifies the input Markdown file. If omitted, the program will read from standard input.

-output: Specifies the output HTML file. If omitted, the program will print the HTML to standard output.

Examples
1. Convert a Markdown file to HTML and print to console:
`go run main.go -input=example.md` or `./markdown_to_html -input = input.md`

2. Convert a Markdown file to HTML and save the result to a file:
`go run main.go -input=example.md -output=output.html` or `./markdown_to_html -input = input.md -output=output.html`

3. Convert Markdown from standard input and print the HTML to the console:
`go run main.go` or `./markdown_to_html`

Enter the Markdown text, then press Ctrl+D (Linux/Mac) or Ctrl+Z (Windows) to finish input.

## Input Markdown Example
```
# Header 1

This is a paragraph with a [link](http://example.com).


## Header 2

Another paragraph.
```

## Output HTML Example

```
<h1>Header 1</h1>
<p>This is a paragraph with a <a href="http://example.com">link</a>.</p>
<h2>Header 2</h2>
<p>Another paragraph.</p>
```

## Test Coverage
Run the following commands to get coverage reports -
1. `go test -cover ./...` to get test coverage percentage.
2. `go tool cover -html=coverage.out -o codecoverage.html` for a detailed html report on test coverage.


## Error Handling
1. If an input file cannot be read, the program will print an error message and exit.
2. If an output file cannot be written, the program will print an error message and exit.
