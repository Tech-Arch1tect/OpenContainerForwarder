package caddyManagement

import "strings"

// joinStrings is a template function to join a slice of strings with a separator
func joinStrings(s []string, sep string) string {
	return strings.Join(s, sep)
}
