package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func Convert(markdown string) string {
	var html strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(markdown))

	inList := false
	listType := ""

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			if inList {
				html.WriteString(fmt.Sprintf("</%s>\n", listType))
				inList = false
			}
			continue
		}

		// Headers
		if strings.HasPrefix(line, "# ") {
			line = fmt.Sprintf("<h1>%s</h1>", strings.TrimPrefix(line, "# "))
		} else if strings.HasPrefix(line, "## ") {
			line = fmt.Sprintf("<h2>%s</h2>", strings.TrimPrefix(line, "## "))
		} else if strings.HasPrefix(line, "### ") {
			line = fmt.Sprintf("<h3>%s</h3>", strings.TrimPrefix(line, "### "))
		} else if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
			// Unordered list
			if !inList || listType != "ul" {
				if inList {
					html.WriteString(fmt.Sprintf("</%s>\n", listType))
				}
				html.WriteString("<ul>\n")
				inList = true
				listType = "ul"
			}
			line = fmt.Sprintf("<li>%s</li>", strings.TrimPrefix(strings.TrimPrefix(line, "- "), "* "))
		} else if matched, _ := regexp.MatchString(`^\d+\. `, line); matched {
			// Ordered list
			if !inList || listType != "ol" {
				if inList {
					html.WriteString(fmt.Sprintf("</%s>\n", listType))
				}
				html.WriteString("<ol>\n")
				inList = true
				listType = "ol"
			}
			line = fmt.Sprintf("<li>%s</li>", regexp.MustCompile(`^\d+\. `).ReplaceAllString(line, ""))
		} else {
			// Paragraph
			if inList {
				html.WriteString(fmt.Sprintf("</%s>\n", listType))
				inList = false
			}
			line = fmt.Sprintf("<p>%s</p>", line)
		}

		// Inline elements
		// Bold
		line = regexp.MustCompile(`\*\*(.+?)\*\*`).ReplaceAllString(line, "<b>$1</b>")
		line = regexp.MustCompile(`__(.+?)__`).ReplaceAllString(line, "<b>$1</b>")

		// Italic
		line = regexp.MustCompile(`\*(.+?)\*`).ReplaceAllString(line, "<i>$1</i>")
		line = regexp.MustCompile(`_(.+?)_`).ReplaceAllString(line, "<i>$1</i>")

		// Images
		line = regexp.MustCompile(`!\[(.+?)\]\((.+?)\)`).ReplaceAllString(line, `<img src="$2" alt="$1">`)

		// Links
		line = regexp.MustCompile(`\[(.+?)\]\((.+?)\)`).ReplaceAllString(line, `<a href="$2">$1</a>`)

		html.WriteString(line + "\n")
	}

	if inList {
		html.WriteString(fmt.Sprintf("</%s>\n", listType))
	}

	return html.String()
}
