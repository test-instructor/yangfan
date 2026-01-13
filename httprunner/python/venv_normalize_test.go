package python

import "testing"

func TestNormalizeVenvDir(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"/Users/taylor/.yf/venv", "/Users/taylor/.yf/venv"},
		{"/Users/taylor/.yf/venv/bin/python3", "/Users/taylor/.yf/venv"},
		{"/Users/taylor/.yf/venv/bin/python", "/Users/taylor/.yf/venv"},
		{"/Users/taylor/.yf/venv/Scripts/python.exe", "/Users/taylor/.yf/venv"},
		{"/Users/taylor/.yf/venv/bin/python3.12", "/Users/taylor/.yf/venv"},
	}

	for _, tc := range cases {
		got := normalizeVenvDir(tc.in)
		if got != tc.want {
			t.Fatalf("normalizeVenvDir(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
