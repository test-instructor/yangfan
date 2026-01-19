package builtin

import "testing"

func TestAppendAddress_DefaultSuffix(t *testing.T) {
	req := map[string]interface{}{
		"body": map[string]interface{}{
			"address": "foo",
		},
	}
	got := appendAddress(req)
	body := got["body"].(map[string]interface{})
	if body["address"] != "foo123" {
		t.Fatalf("unexpected address: %v", body["address"])
	}
}

func TestAppendAddress_CustomSuffix(t *testing.T) {
	req := map[string]interface{}{
		"body": map[string]interface{}{
			"address": "foo",
		},
	}
	got := appendAddress(req, "bar")
	body := got["body"].(map[string]interface{})
	if body["address"] != "foobar" {
		t.Fatalf("unexpected address: %v", body["address"])
	}
}
