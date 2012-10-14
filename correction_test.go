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
		s1,
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

	// The distance parameter is an upper limit, so expected1 & expected2
	// should be returned
	expected2 = append(expected2, expected1...)
	deletions = trie.Deletions(s1, 2)
	if len(deletions) != len(expected2) {
		t.Errorf("Deletions has the wrong number of words %v", deletions)
	}
	assertAllIn(t, expected2, deletions)
}

func TestDeletionsUTF(t *testing.T) {
	s1 := "ab狐d犬"
	expected1 := []string{
		s1,
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

	expected3 = append(expected3, expected1...)
	deletions = trie.Deletions(s1, 3)
	if len(deletions) != len(expected3) {
		t.Errorf("Deletions has the wrong number of words %v", deletions)
	}
	assertAllIn(t, expected3, deletions)
}

func TestPermutations(t *testing.T) {
	s1 := "ab狐d犬"
	expected := []string{
		s1,
		"ba狐d犬",
		"abd狐犬",
		"犬d狐ba"}
	s2 := "abc"
	s3 := "abd"

	trie := NewTrie()
	trie.InsertString(s1)
	trie.InsertString(s2)
	trie.InsertString(s3)
	for _, s := range expected {
		trie.InsertString(s)
	}

	permutations := trie.Permutations(s1, 1)
	if len(permutations) != len(expected) {
		t.Errorf("Permutations has the wrong number of words %v", permutations)
	}
	assertAllIn(t, expected, permutations)
}

func TestAdditions(t *testing.T) {
	s1 := "ab狐d犬"
	expected1 := []string{
		s1,
		"abc狐d犬",
		"cab狐d犬"}
	expected2 := []string{
		"abc狐d犬e",
		"abc狐ed犬"}
	s2 := "abc"
	s3 := "abd"
	s4 := "犬d狐ba"

	trie := NewTrie()
	trie.InsertString(s1)
	trie.InsertString(s2)
	trie.InsertString(s3)
	trie.InsertString(s4)
	for _, s := range expected1 {
		trie.InsertString(s)
	}
	for _, s := range expected2 {
		trie.InsertString(s)
	}

	additions := trie.Additions(s1, 1)
	if len(additions) != len(expected1) {
		t.Errorf("Additions has the wrong number of words %v", additions)
	}
	assertAllIn(t, expected1, additions)

	expected2 = append(expected2, expected1...)
	additions = trie.Additions(s1, 2)
	if len(additions) != len(expected2) {
		t.Errorf("Additions has the wrong number of words %v", additions)
	}
	assertAllIn(t, expected2, additions)
}

func TestSubstitutions(t *testing.T) {
	s1 := "ab狐d犬"
	expected1 := []string{
		s1,
		"ac狐d犬",
		"犬b狐d犬",
		"ab狐de"}
	expected2 := []string{
		"bc狐d犬",
		"a犬狐犬犬"}
	s2 := "abc"
	s3 := "abd"
	s4 := "犬d狐ba"

	trie := NewTrie()
	trie.InsertString(s1)
	trie.InsertString(s2)
	trie.InsertString(s3)
	trie.InsertString(s4)
	for _, s := range expected1 {
		trie.InsertString(s)
	}
	for _, s := range expected2 {
		trie.InsertString(s)
	}

	substitutions := trie.Substitutions(s1, 1)
	if len(substitutions) != len(expected1) {
		t.Errorf("Substitutions has the wrong number of words %v", substitutions)
	}
	assertAllIn(t, expected1, substitutions)

	expected2 = append(expected2, expected1...)
	substitutions = trie.Substitutions(s1, 2)
	if len(substitutions) != len(expected2) {
		t.Errorf("Substitutions has the wrong number of words %v", substitutions)
	}
	assertAllIn(t, expected2, substitutions)
}

func TestSuggestions(t *testing.T) {
	s1 := "toad"
	expected := []string{
		s1,
		"todd",
		"load",
		"toads",
		"tod",
		"toda"}
	unexpected := []string{
		"robert",
		"today", // Note: today should probably be suggested in the future
		"teddy",
		"bad"}

	trie := NewTrie()
	trie.InsertString(s1)
	for _, s := range expected {
		trie.InsertString(s)
	}
	for _, s := range unexpected {
		trie.InsertString(s)
	}

	suggestions := trie.SuggestWords(s1, 2)
	if len(suggestions) != len(expected) {
		t.Errorf("Suggestions has the wrong number of words %v", suggestions)
	}
	assertAllIn(t, expected, suggestions)
}

func TestLoadDict(t *testing.T) {
	fname := "/usr/share/dict/words"
	trie, err := TrieFromFile(fname)
	if err != nil {
		t.Fatal(err)
	}
	if !trie.ContainsString("hello") {
		t.Error("'hello' not found in dictionary!")
	}
	suggestions := trie.SuggestWords("hyllo", 1)
	expected := []string{
		"hello",
		//"hell",  // This doesn't show up because it's a sub & deletion FIXME
		"holly",
		"hollo"}
	if len(expected) != len(suggestions) {
		t.Errorf("Suggestions has the wrong number of words %v", suggestions)
	}
	assertAllIn(t, expected, suggestions)
}

func TestDistance(t *testing.T) {
	s1 := "toad"
	// Ensure sorting is correct
	expectedOrdered := Matches{
		Match{runes(s1), 0, 0},
		Match{runes("load"), 1, 0},
		Match{runes("toads"), 1, 0},
		Match{runes("todd"), 1, 0},
		Match{runes("tod"), 2, 0},
		Match{runes("toda"), 2, 0},
	}
	expected := make([]string, len(expectedOrdered))
	for i, match := range expectedOrdered {
		expected[i] = string(match.Word)
	}
	unexpected := []string{
		"robert",
		"today", // Note: today should probably be suggested in the future
		"teddy",
		"bad"}

	trie := NewTrie()
	trie.InsertString(s1)
	for _, s := range expected {
		trie.InsertString(s)
	}
	for _, s := range unexpected {
		trie.InsertString(s)
	}

	suggestions := trie.suggestions(runes(s1), 2)
	if len(expected) != len(suggestions) {
		t.Errorf("Suggestions has the wrong number of matches %v", suggestions)
	}


	for i, match := range suggestions {
		if !expectedOrdered[i].Equal(match) {
			t.Errorf("%v != %v\n", expectedOrdered[i], match)
		}
	}
}

func BenchmarkAdditions1(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.additions(r, 1)})
}

func BenchmarkAdditions2(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.additions(r, 2)})
}

func BenchmarkDeletions1(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.deletions(r, 1)})
}

func BenchmarkDeletions2(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.deletions(r, 2)})
}

func BenchmarkSubstitutions1(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.substitutions(r, 1)})
}

func BenchmarkSubstitutions2(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.substitutions(r, 2)})
}

func BenchmarkPermutations1(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.permutations(r, 1)})
}

func BenchmarkPermutations2(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.permutations(r, 2)})
}

func BenchmarkSuggestions1(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.suggestions(r, 1)})
}

func BenchmarkSuggestions2(b *testing.B) {
	benchmarkOp(b, func(trie *Trie, r []rune) {trie.suggestions(r, 2)})
}

func benchmarkOp(b *testing.B, op func(*Trie, []rune)) {
	b.StopTimer()
	fname := "/usr/share/dict/words"
	trie, err := TrieFromFile(fname)
	if err != nil {
		b.Fatal(err)
	}

	children := trie.AllFullChildren()
	for i:= 0; i < b.N; i++ {
		for j, child := range children {
			if j % 1000 != 0 {
				continue
			}
			r := runes(child)
			b.StartTimer()
			op(trie, r)
			b.StopTimer()
		}
	}
}
