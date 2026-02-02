package textkit

import (
	"strings"
	"unicode"
)

// Slugify converts an arbitrary string into a URL-friendly slug.
//
// Rules:
//   - Lowercases letters
//   - Keeps only latin letters and digits (A-Z, a-z, 0-9)
//   - Converts any run of non-alphanumerics into a single '-'
//   - Trims leading/trailing '-'
func Slugify(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	var b strings.Builder
	b.Grow(len(s))

	prevDash := false
	for _, r := range s {
		if isASCIILetterOrDigit(r) {
			b.WriteRune(unicode.ToLower(r))
			prevDash = false
			continue
		}

		if !prevDash && b.Len() > 0 {
			b.WriteByte('-')
			prevDash = true
		}
	}

	out := b.String()
	out = strings.Trim(out, "-")
	return out
}

func isASCIILetterOrDigit(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}
