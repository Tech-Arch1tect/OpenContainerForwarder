package misc

import "strings"

func StripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(chr, r) {
			return r
		}
		return -1
	}, str)
}
