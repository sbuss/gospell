package gospell

import (
	"testing"
)

func assertAllIn(t *testing.T, expected, actual []string) {
	words := make(map[string]int)
	for _, word := range actual {
		words[word] += 1
	}
	for k, v := range words {
		if v != 1 {
			t.Errorf("%v occurs %d times!", k, v)
		}
		found := false
		for _, e := range expected {
			if k == e {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%v wasn't in the source set", k)
		}
	}
}

func TestDeletions(t *testing.T) {
	s1 := "abcd"
	expected := []string{
		"abc",
		"abd",
		"acd",
		"bcd"}
	s6 := "abcdef"
	s7 := "bcdef"

	trie := NewTrie()
	trie.InsertString(s1)
	for _, s := range expected {
		trie.InsertString(s)
	}
	trie.InsertString(s6)
	trie.InsertString(s7)

	deletions := trie.Deletions(s1)
	if len(deletions) != len(expected) {
		t.Errorf("Deletions is %v", deletions)
	}
	assertAllIn(t, expected, deletions)
}

func TestDeletionsUTF(t *testing.T) {
	s1 := "ab狐d犬"
	expected := []string{
		"b狐d犬",
		"a狐d犬",
		"abd犬",
		"ab狐犬",
		"ab狐d"}

	trie := NewTrie()
	trie.InsertString(s1)
	for _, s := range expected {
		trie.InsertString(s)
	}
	deletions := trie.Deletions(s1)
	if len(deletions) != len(expected) {
		t.Errorf("Deletions is %v", deletions)
	}
	assertAllIn(t, expected, deletions)
}
