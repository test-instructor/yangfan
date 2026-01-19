package runTestCase

import "testing"

func TestSanitizeRequestURL(t *testing.T) {
	raw := " `http://httpbin.org/post\\` "
	got := sanitizeRequestURL(raw)
	if got != "http://httpbin.org/post" {
		t.Fatalf("unexpected url: %q", got)
	}
}
