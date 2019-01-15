// Package acronym creates acronyms from strings.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes a string as input and returns a string as output.
func Abbreviate(s string) string {
	re := regexp.MustCompile("\\b\\w")
	return strings.ToUpper(strings.Join(re.FindAllString(strings.Replace(s, "'", "", -1), -1), ""))
}
