package textkit

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	cases := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{"both empty", "", "", 0},
		{"one empty", "", "abc", 3},
		{"equal", "kitten", "kitten", 0},
		{"classic", "kitten", "sitting", 3},
		{"unicode", "café", "cafe", 1},
		{"emoji", "😀", "😃", 1},
		{"swap", "ab", "ba", 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := LevenshteinDistance(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("LevenshteinDistance(%q,%q)=%d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestLevenshteinDistance_Symmetric(t *testing.T) {
	a := "golang"
	b := "gopher"
	if LevenshteinDistance(a, b) != LevenshteinDistance(b, a) {
		t.Fatalf("distance should be symmetric")
	}
}
