package textkit

import "testing"

func TestSlugify(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"spaces only", "   ", ""},
		{"simple words", "Hello World", "hello-world"},
		{"punctuation", "Hello, World!!!", "hello-world"},
		{"multiple separators", "Go---is___fun", "go-is-fun"},
		{"trim dashes", "--Hello--World--", "hello-world"},
		{"numbers", "Issue 404: Not Found", "issue-404-not-found"},
		{"unicode kept out", "Привет мир", ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Slugify(tc.in)
			if got != tc.want {
				t.Fatalf("Slugify(%q)=%q; want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestSlugify_NoDoubleDashes(t *testing.T) {
	got := Slugify("a---b   c!!d")
	if got != "a-b-c-d" {
		t.Fatalf("unexpected slug: %q", got)
	}
}
