package gospell

import "log"

// Go through every rune and do t.Contains(s[0:n]+s[n+1:len(s)])
// The distance parameter is currently ignored and is implicitly 1
func (t *Trie) Deletions(s string, distance int) []string {
	// Based on UTF-aware string reversal by Russ Cox.
	// See http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
	n := 0
	deletions := []string{}

	// First build a list of valid runes
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++
	}

	// Re-assign runes to get the correct length
	// Removing this causes invalid memory address errors
	runes = runes[0:n]

	c := make([]rune, n-1)

	// Then build substrings
	for i, _ := range runes {
		c = []rune{}
		log.Printf("c is %v", c)
		for j, v := range runes {
			if i == j {
				continue
			}
			c = append(c, v)
		}
		log.Printf("c is %v", c)
		if t.ContainsString(string(c)) {
			deletions = append(deletions, string(c))
		}
	}
	return deletions
}
