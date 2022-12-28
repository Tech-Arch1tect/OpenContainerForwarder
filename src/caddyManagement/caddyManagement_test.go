package caddyManagement

import "testing"

func TestJoinStrings(t *testing.T) {
	s := []string{"a", "b", "c"}
	sep := ","
	expected := "a,b,c"
	result := joinStrings(s, sep)
	if result != expected {
		t.Errorf("joinStrings(%v, %v) = %v; want %v", s, sep, result, expected)
	}
}
