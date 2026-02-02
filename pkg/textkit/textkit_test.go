package textkit

import "testing"

func TestSlugify(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   string
		want string
	}{
		{"simple", "Hello World", "hello-world"},
		{"trim", "  Hello   World  ", "hello-world"},
		{"punctuation", "Go, Go, Go!!!", "go-go-go"},
		{"mixed", "CI/CD on GitHub Actions", "ci-cd-on-github-actions"},
		{"already_slug", "hello-world", "hello-world"},
		{"non_latin_removed", "Привет мир", ""},
		{"digits", "Version 2.0.1", "version-2-0-1"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := Slugify(tc.in); got != tc.want {
				t.Fatalf("Slugify(%q)=%q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestLevenshtein(t *testing.T) {
	t.Parallel()

	cases := []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"a", "", 1},
		{"", "abc", 3},
		{"kitten", "sitting", 3},
		{"flaw", "lawn", 2},
		{"go", "go", 0},
		{"gumbo", "gambol", 2},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.a+"->"+tc.b, func(t *testing.T) {
			t.Parallel()
			if got := Levenshtein(tc.a, tc.b); got != tc.want {
				t.Fatalf("Levenshtein(%q,%q)=%d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
