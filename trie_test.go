package gospell

import (
	"bufio"
	"os"
	"strings"
	"testing"
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

func HasEveryElement(t *testing.T, s1, s2 []string) bool {
	s1Map := map[string]int{}
	for _, val := range s1 {
		s1Map[val] += 1
	}
	for _, val := range s2 {
		s1Map[val] -= 1
	}
	for _, val := range s1Map {
		if val != 0 {
			return false
		}
	}
	return true
}

func TestAllChildStrings(t *testing.T) {
	s1 := "Salmon"
	s2 := "Salmonella"

	trie := NewTrie()
	trie.InsertString(s1)
	trie.InsertString(s2)

	allFullStrings := trie.AllFullChildren()
	if !HasEveryElement(t, allFullStrings, []string{s1, s2}) {
		t.Error(allFullStrings)
	}

	t1 := trie.Get(strings.NewReader(s1))
	allFullStrings = t1.AllFullChildren()
	if !HasEveryElement(t, allFullStrings, []string{"ella"}) {
		t.Error(allFullStrings)
	}
}

func BenchmarkLoadDict(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		b.StopTimer()
		f, err := os.Open("/usr/share/dict/words")
		if err != nil {
			b.Fatal("Can't find words file")
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		word, err := reader.ReadString('\n')
		trie := NewTrie()
		for err == nil {
			// Don't insert the '\n'
			w := word[:len(word)-1]
			b.StartTimer()
			trie.InsertString(w)
			b.StopTimer()
			word, err = reader.ReadString('\n')
		}
	}
}

func BenchmarkLoadDictMap(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		b.StopTimer()
		f, err := os.Open("/usr/share/dict/words")
		if err != nil {
			b.Fatal("Can't find words file")
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		word, err := reader.ReadString('\n')
		m := make(map[string]int)
		for err == nil {
			// Don't insert the '\n'
			w := word[:len(word)-1]
			b.StartTimer()
			m[w] = 1
			b.StopTimer()
			word, err = reader.ReadString('\n')
		}
	}
}
