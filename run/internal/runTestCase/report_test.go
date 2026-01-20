package runTestCase

import (
	"strings"
	"testing"
)

func TestStringifyAttachments(t *testing.T) {
	if got := stringifyAttachments(nil); got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}

	if got := stringifyAttachments("boom"); got != "boom" {
		t.Fatalf("expected boom, got %q", got)
	}

	got := stringifyAttachments(map[string]interface{}{"a": 1})
	if !strings.Contains(got, `"a"`) || !strings.Contains(got, "1") {
		t.Fatalf("unexpected json string: %q", got)
	}
}
