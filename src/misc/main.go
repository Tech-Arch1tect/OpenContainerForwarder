package misc

import "strings"

// StripChars removes all characters from a string that are in the second string
func StripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(chr, r) {
			return r
		}
		return -1
	}, str)
}
