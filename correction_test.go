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
	expected1 := []string{
		"abc",
		"abd",
		"acd",
		"bcd"}
	expected2 := []string{
		"ab",
		"cd"}
	s6 := "abcdef"
	s7 := "bcdef"

	trie := NewTrie()
	trie.InsertString(s1)
	for _, s := range expected1 {
		trie.InsertString(s)
	}
	for _, s := range expected2 {
		trie.InsertString(s)
	}
	trie.InsertString(s6)
	trie.InsertString(s7)

	deletions := trie.Deletions(s1, 0)
	if len(deletions) != 1 {
		t.Errorf("Deletions has too many words %v", deletions)
	}
	if deletions[0] != s1 {
		t.Errorf("Deletion of length 0 should be equivalent to Get")
	}

	deletions = trie.Deletions(s1, 1)
	if len(deletions) != len(expected1) {
		t.Errorf("Deletions has the wrong number of words %v", deletions)
	}
	assertAllIn(t, expected1, deletions)

	deletions = trie.Deletions(s1, 2)
	if len(deletions) != len(expected2) {
		t.Errorf("Deletions has the wrong number of words %v", deletions)
	}
	assertAllIn(t, expected2, deletions)
}

func TestDeletionsUTF(t *testing.T) {
	s1 := "ab狐d犬"
	expected1 := []string{
		"b狐d犬",
		"a狐d犬",
		"abd犬",
		"ab狐犬",
		"ab狐d"}
	expected3 := []string{
		"狐犬",
		"ab"}
	trie := NewTrie()
	trie.InsertString(s1)
	for _, s := range expected1 {
		trie.InsertString(s)
	}
	for _, s := range expected3 {
		trie.InsertString(s)
	}
	deletions := trie.Deletions(s1, 1)
	if len(deletions) != len(expected1) {
		t.Errorf("Deletions has the wrong number of words %v", deletions)
	}
	assertAllIn(t, expected1, deletions)

	deletions = trie.Deletions(s1, 3)
	if len(deletions) != len(expected3) {
		t.Errorf("Deletions has the wrong number of words %v", deletions)
	}
	assertAllIn(t, expected3, deletions)
}
