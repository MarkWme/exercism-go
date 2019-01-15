// Package bob implements a function that returns a response to a remark
package bob

import (
	"regexp"
	"strings"
)

// Hey function takes a remark as a string input and returns a response as a string
func Hey(remark string) string {
	// Remove whitespace from remark
	remark = strings.TrimSpace(remark)
	// Set default response
	response := "Whatever."
	switch {
	// No remark
	case remark == "":
		response = "Fine. Be that way!"
	// Shouting is detected by matching the remark against a capitalised version of the remark.
	// However, if the remark only contains non-alphabetic characters, it will also match.
	// So, first check if the remark contains any alphabetic characters.
	case func() bool { matched, _ := regexp.MatchString("[[:alpha:]]", remark); return matched }():
		// If remark matches capitalised remark, then it's being shouted
		if remark == strings.ToUpper(remark) {
			// Is remark shouting, or a shouting question?
			if strings.HasSuffix(remark, "?") {
				// Shouted question
				response = "Calm down, I know what I'm doing!"
			} else {
				// Just shouting
				response = "Whoa, chill out!"
			}
		} else {
			if strings.HasSuffix(remark, "?") {
				// Just a question
				response = "Sure."
			}
		}
	// Match case where remark contains no alphabetic characters
	case strings.HasSuffix(remark, "?"):
		response = "Sure."
	}
	return response
}
