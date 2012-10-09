package gospell

import "log"

func (t *Trie) Deletions(s string) []string {
	// Go through every character and do t.Contains(s[0:n]+s[n+1:len(s)])
	// TODO: Is this uniode safe?
	n := 0
	deletions := []string{}

	// First build a list of valid runes
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++
	}

	// Re-assign runes to get the correct length
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
