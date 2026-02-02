package textkit

// LevenshteinDistance returns the Levenshtein (edit) distance between two strings.
// It counts the minimum number of single-character insertions, deletions or substitutions
// required to transform a into b.
//
// This implementation works on runes (Unicode code points) for correct behavior with
// non-ASCII input.
func LevenshteinDistance(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)

	// Ensure rb is the shorter one to minimize memory.
	if len(ra) < len(rb) {
		ra, rb = rb, ra
	}

	prev := make([]int, len(rb)+1)
	cur := make([]int, len(rb)+1)

	for j := 0; j <= len(rb); j++ {
		prev[j] = j
	}

	for i := 1; i <= len(ra); i++ {
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 0
			if ra[i-1] != rb[j-1] {
				cost = 1
			}

			cur[j] = min3(
				prev[j]+1,    // delete
				cur[j-1]+1,   // insert
				prev[j-1]+cost, // substitute
			)
		}
		prev, cur = cur, prev
	}

	return prev[len(rb)]
}

func min3(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
