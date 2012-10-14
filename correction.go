package gospell

import "sort"

// Find all strings in the trie within a given deletion distance
// For example, for the Trie{"abcd", "abc", "ab", "cd"},
// Deletions("abcd", 2) would return ["ab", "cd"] and
// Deletions("abcd", 1) would return ["abc"]
func (t *Trie) Deletions(s string, distance int) []string {
	return t.deletions(runes(s), distance).Strings()
}

// Find all words in the Trie adding at most `distance` runes
func (t *Trie) Additions(s string, distance int) []string {
	return t.additions(runes(s), distance).Strings()
}

// Find all words in the Trie adding at most `distance` runes
func (t *Trie) Substitutions(s string, distance int) []string {
	return t.substitutions(runes(s), distance).Strings()
}

// Find all strings matching permutations of the given distance
func (t *Trie) Permutations(s string, distance int) []string {
	return t.permutations(runes(s), distance).Strings()
}

// Return spelling suggestions, ranked by Distance then lexicographically
func (t *Trie) SuggestWords(s string, distance int) []string {
	return t.suggestions(runes(s), distance).Strings()
}

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

func (t *Trie) deletions(r []rune, distance int) Matches {
	matches := Matches{}

	if len(r) == 0 {
		if t.leaf {
			matches = append(matches, Match{})
		}
		return matches
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
		childMatches := child.deletions(rest, distance)
		for _, m := range childMatches {
			matches = append(matches, m.update(first, 0, 0))
		}
	}
	// Case 2
	if distance > 0 {
		childMatches := t.deletions(rest, distance-1)
		for _, m := range childMatches {
			matches = append(matches, m.update(0, 2, 0))
		}
	}

	return matches
}

// Find all permutations of r that exist in the trie
// This does not currently respect the distance parameter
func (t *Trie) permutations(r []rune, distance int) Matches {
	matches := Matches{}

	if len(r) == 0 {
		if t.leaf {
			matches = append(matches, Match{})
		}
		return matches
	}

	for i, c := range r {
		// If we don't make a new slice things get overwritten... not sure why
		rest := make([]rune, 0)
		rest = append(rest, r[:i]...)
		rest = append(rest, r[i+1:]...)
		child := t.children[c]
		if child != nil {
			childMatches := child.permutations(rest, distance)
			for _, cr := range childMatches {
				matches = append(matches, cr.update(c, 2, 0))
			}
		}
	}

	return matches
}

func (t *Trie) additions(r []rune, distance int) Matches {
	matches := Matches{}

	// Three cases:
	// 0. All runes have been seen, but we have distance to spare, so add runes
	// 1. Pop the first rune from the list, recurse on the child with that rune
	//   as the key
	// 2. Recurse on all children of the current node, effectively adding the
	//   key for the child to the word
	if len(r) == 0 {
		// Case 0: no more runes but more to add
		if t.leaf {
			matches = append(matches, Match{})
		}
	} else {
		// Case 1
		first := r[0]
		rest := r[1:]
		child, ok := t.children[first]
		if ok {
			childMatches := child.additions(rest, distance)
			for _, cr := range childMatches {
				matches = append(matches, cr.update(first, 0, 0))
			}
		}
	}

	// Case 2
	if distance > 0 {
		for c, child := range t.children {
			childMatches := child.additions(r, distance-1)
			for _, cr := range childMatches {
				matches = append(matches, cr.update(c, 1, 0))
			}
		}
	}

	return matches
}

func (t *Trie) substitutions(r []rune, distance int) Matches {
	matches := Matches{}

	if len(r) == 0 {
		if t.leaf {
			matches = append(matches, Match{})
		}
		return matches
	}

	// Two cases:
	// 1. Pop the first rune from the list, recurse on the child with that rune
	//  as the key
	// 2. Pop the first rune from the list, recurse on the current node
	//  (effectively ignoring this rune)
	first := r[0]
	rest := r[1:]
	for c, child := range t.children {
		if child == nil {
			continue
		}
		d := 0
		childMatches := Matches{}
		if c == first {
			// Case 1
			childMatches = child.substitutions(rest, distance)
		} else if distance > 0 {
			// Case 2
			childMatches = child.substitutions(rest, distance-1)
			d = 1
		} else {
			continue
		}

		for _, cr := range childMatches {
			matches = append(matches, cr.update(c, d, 0))
		}
	}
	return matches
}

func (t *Trie) suggestions(r []rune, distance int) Matches {
	suggestions := Matches{}
	additions := t.additions(r, distance)
	deletions := t.deletions(r, distance)
	permutations := t.permutations(r, distance)
	substitutions := t.substitutions(r, distance)

	unique := func(matches Matches) Matches {
		dupes := make(map[string]int)
		ret := make(Matches, len(matches))
		i := 0
		for _, match := range matches {
			word := string(match.Word)
			if _, ok := dupes[word]; !ok {
				ret[i] = match
				dupes[word] = 1
				i += 1
			}
		}
		return ret[:i]
	}

	// Combine and remove duplicates
	suggestions = append(suggestions, additions...)
	suggestions = append(suggestions, deletions...)
	suggestions = append(suggestions, permutations...)
	suggestions = append(suggestions, substitutions...)
	// Sort by distance so we only keep the lowest scoring match
	sort.Sort(ByDistance{suggestions})
	return unique(suggestions)
}
