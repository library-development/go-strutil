package strutil

import (
	"strings"
)

// Filter returns a new string containing only the characters from s that are in validChars.
func Filter(validChars, s string) string {
	var b strings.Builder
	for _, r := range s {
		if strings.ContainsRune(validChars, r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
