package hrp

import "testing"

func TestParse_AppendAddressBuiltin(t *testing.T) {
	p := NewParser()
	vars := map[string]interface{}{
		"request": map[string]interface{}{
			"body": map[string]interface{}{
				"address": "foo",
			},
		},
	}
	got, err := p.Parse("${append_address($request)}", vars)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	req, ok := got.(map[string]interface{})
	if !ok {
		t.Fatalf("unexpected type: %T", got)
	}
	body := req["body"].(map[string]interface{})
	if body["address"] != "foo123" {
		t.Fatalf("unexpected address: %v", body["address"])
	}
}
