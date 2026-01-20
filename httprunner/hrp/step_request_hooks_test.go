package hrp

import "testing"

func TestParseHeadersMap(t *testing.T) {
	got, err := parseHeadersMap(map[string]string{"A": "1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["A"] != "1" {
		t.Fatalf("expected A=1, got %v", got["A"])
	}

	got, err = parseHeadersMap(map[string]interface{}{"B": "2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["B"] != "2" {
		t.Fatalf("expected B=2, got %v", got["B"])
	}

	_, err = parseHeadersMap(map[string]interface{}{"C": 3})
	if err == nil {
		t.Fatalf("expected error")
	}

	_, err = parseHeadersMap([]string{"x"})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestMarshalRequestBody(t *testing.T) {
	b, err := marshalRequestBody("ok")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(b) != "ok" {
		t.Fatalf("expected ok, got %q", string(b))
	}

	b, err = marshalRequestBody(map[string]interface{}{"a": 1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(b) == 0 {
		t.Fatalf("expected non-empty json")
	}
}
