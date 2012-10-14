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

// Prepend a rune to a slice of runes
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
	childRunes := t.deletions(runes(s), distance)
	strings := make([]string, len(childRunes))
	for i, r := range childRunes {
		strings[i] = string(r)
	}

	return strings
}

// Find all permutations of r that exist in the trie
// This does not currently respect the distance parameter
func (t *Trie) permutations(r []rune, distance int) [][]rune {
	runes := make([][]rune, 0)

	if len(r) == 0 {
		if t.leaf {
			runes = append(runes, []rune{})
		}
		return runes
	}

	for i, c := range r {
		// If we don't make a new slice things get overwritten... not sure why
		rest := make([]rune, 0)
		rest = append(rest, r[:i]...)
		rest = append(rest, r[i+1:]...)
		child := t.children[c]
		if child != nil {
			childRunes := child.permutations(rest, distance)
			for _, cr := range childRunes {
				runes = append(runes, prependRune(cr, c))
			}
		}
	}

	return runes
}

// Find all strings matching permutations of the given distance
func (t *Trie) Permutations(s string, distance int) []string {
	childRunes := t.permutations(runes(s), distance)
	strings := make([]string, len(childRunes))
	for i, r := range childRunes {
		strings[i] = string(r)
	}

	return strings
}

func (t *Trie) additions(r []rune, distance int) [][]rune {
	runes := make([][]rune, 0)

	if len(r) == 0 && distance == 0 {
		if t.leaf {
			runes = append(runes, []rune{})
		}
		return runes
	}

	// Two cases:
	// 1. Pop the first rune from the list, recurse on the child with that rune
	//   as the key
	// 2. Recurse on all children of the current node, effectively adding the
	//   key for the child to the word
	if len(r) > 0 {
		first := r[0]
		rest := r[1:]
		// Case 1
		child := t.children[first]
		if child != nil {
			childRunes := child.additions(rest, distance)
			for _, c := range childRunes {
				runes = append(runes, prependRune(c, first))
			}
		}
	}
	// Case 2
	if distance > 0 {
		for c, child := range t.children {
			if child == nil {
				continue
			}
			childRunes := child.additions(r, distance-1)
			for _, cr := range childRunes {
				runes = append(runes, prependRune(cr, c))
			}
		}
	}
	return runes
}

// Find all words in the Trie adding at most `distance` runes
func (t *Trie) Additions(s string, distance int) []string {
	childRunes := t.additions(runes(s), distance)
	strings := make([]string, len(childRunes))
	for i, r := range childRunes {
		strings[i] = string(r)
	}

	return strings
}
