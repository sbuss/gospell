package gospell

// Convert a string into a slice of runes
func runes(s string) []rune {
	// Based on UTF-aware string reversal by Russ Cox.
	// See http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
	n := 0
	// First build a list of valid runes
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++
	}
	// Re-assign runes to get the correct length
	// Removing this causes invalid memory address errors
	runes = runes[0:n]
	return runes
}

func prependRune(runes []rune, newRune rune) []rune {
	return append([]rune{newRune}, runes...)
}

func (t *Trie) deletions(r []rune, distance int) [][]rune {
	runes := make([][]rune, 0)

	if len(r) == 0 {
		if t.leaf && distance == 0 {
			runes = append(runes, []rune{})
		}
		return runes
	}

	// Two cases:
	// 1. Pop the first rune from the list, recurse on the child with that rune
	//  as the key
	// 2. Pop the first rune from the list, recurse on the current node
	//  (effectively ignoring this rune)

	first := r[0]
	rest := r[1:]
	// Case 1
	child := t.children[first]
	if child != nil {
		childRunes := child.deletions(rest, distance)
		for _, c := range childRunes {
			runes = append(runes, prependRune(c, first))
		}
	}
	// Case 2
	if distance > 0 {
		runes = append(runes, t.deletions(rest, distance-1)...)
	}

	return runes
}

// Find all strings in the trie within a given deletion distance
// For example, for the Trie{"abcd", "abc", "ab", "cd"},
// Deletions("abcd", 2) would return ["ab", "cd"] and 
// Deletions("abcd", 1) would return ["abc"]
func (t *Trie) Deletions(s string, distance int) []string {
	strings := make([]string, 0)

	childRunes := t.deletions(runes(s), distance)
	for _, r := range childRunes {
		strings = append(strings, string(r))
	}

	return strings
}
