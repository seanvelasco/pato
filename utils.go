package main

import (
	"html"
	"strings"
)

func formatString(result string) string {
	result = strings.ReplaceAll(result, "<b>", "**")
	result = strings.ReplaceAll(result, "</b>", "**")
	return html.UnescapeString(result)
}

func removeHTMLBTag(result string) string {
	result = strings.ReplaceAll(result, "<b>", "")
	result = strings.ReplaceAll(result, "</b>", "")
	return html.UnescapeString(result)
}
