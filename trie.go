/*Gospell suggests alternate spellings for words.

Based on https://github.com/AlanQuatermain/go-trie and
https://github.com/sbuss/pyspellsug
*/
package gospell

import (
	"fmt"
	"strings"
)

type children map[rune]*Trie

type Trie struct {
	children children
	leaf     bool
}

// Create a new Trie with no children and leaf=false
func NewTrie() *Trie {
	t := new(Trie)
	t.children = make(children)
	t.leaf = false
	return t
}

// Insert a strings.Reader into the Trie
func (t *Trie) Insert(s *strings.Reader) {
	rune, _, err := s.ReadRune()
	if err != nil {
		// We have reached EOF
		t.leaf = true
		return
	}

	child := t.children[rune]
	if child == nil {
		child = NewTrie()
		t.children[rune] = child
	}
	child.Insert(s)
}

// Insert a string into the Trie
func (t *Trie) InsertString(s string) {
	t.Insert(strings.NewReader(s))
}

// Get the Trie at the end of a strings.Reader
func (t *Trie) Get(s *strings.Reader) *Trie {
	rune, _, err := s.ReadRune()
	if err != nil {
		// We have reached EOF
		return t
	}
	t = t.children[rune]
	if t == nil {
		return nil
	}
	return t.Get(s)
}

// Return true if the Trie contains the word in a strings.Reader
// This should only be true if the word is a leaf. For example, say you
// Insert the strings Salmon and Salmonella. Sal will not be Contained in the
// Trie, but Salmon and Salmonella both are.
func (t *Trie) Contains(s *strings.Reader) bool {
	return t.Get(s).leaf
}

// Check is a String is Contained in the Trie. See Trie.Contains.
func (t *Trie) ContainsString(s string) bool {
	if s == "" {
		return false
	}

	return t.Contains(strings.NewReader(s))
}

// Get all of the complete child words under this Trie node
func (t *Trie) AllFullChildren() []string {
	childStrings := []string{}

	for r, child := range t.children {
		if child != nil {
			if child.leaf {
				childStrings = append(childStrings, string(r))
			}
			for _, ccs := range child.AllFullChildren() {
				childStrings = append(childStrings, string(r)+ccs)
			}
		}
	}
	return childStrings
}

// Convert a Trie to a String
func (t *Trie) String() string {
	c := ""
	for r, child := range t.children {
		c += fmt.Sprintf("%q: %v", r, child)
	}
	s := fmt.Sprintf("{leaf: %t, c: %v}", t.leaf, c)
	return s
}
