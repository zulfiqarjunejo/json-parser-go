package jp

import "testing"

// Test Parse() for "{}"
func TestParseEmptyObject(t *testing.T) {
	err := Parse("{}")
	if err != nil {
		t.Fatalf(`did not expect an error, received = %v`, err)
	}
}

// Test Parse() for ""
func TestParseBlankString(t *testing.T) {
	err := Parse("")
	if err == nil {
		t.Fatalf(`Expected an error, received nil.`)
	}
}
