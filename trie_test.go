package gospell

import (
	"testing"
	"strings"
)

func TestInsert(t *testing.T) {
	s1 := "Salmon"
	s2 := "Salmonella"

	trie := NewTrie()
	trie.Insert(strings.NewReader(s1))
	trie.InsertString(s2)
	if !trie.ContainsString(s1) {
		t.Errorf("%q is not in the Trie", s1)
	}
	if !trie.Contains(strings.NewReader(s2)) {
		t.Errorf("%q is not in the Trie", s2)
	}
	if trie.ContainsString(s1[:len(s1)-1]) {
		t.Error("Shouldn't contain part of the word.")
	}
	if trie.Get(strings.NewReader("abcdefg")) != nil {
		t.Error("This Trie shouldn't exist!")
	}
}
