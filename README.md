gospell
=======

Spelling suggestion in Go

Usage
=====

```go
import (
	"bufio"
	"github.com/sbuss/gospell"
	"os"
)

func LoadDict(path string) gospell.Trie {
	f, err := os.Open(path)
	if err != nil {
		t.Fatal("Can't find words file")
	}
	trie := gospell.NewTrie()
	reader := bufio.NewReader(f)
	word, err := reader.ReadString('\n')
	for err == nil {
		// Don't insert the '\n'
		trie.InsertString(word[:len(word)-1])
		word, err = reader.ReadString('\n')
	}
	return trie
}

// ...

func main() {
	trie := LoadDict("/usr/share/dict/words")
	suggestions := trie.SuggestWords("gospell", 2)
}
```

Warnings & Caveats
==================
The spelling suggestions are rudimentary and don't rank results in a good order
(which is to say the ranking is arbitrary).

Memory and performance both seem ok, but I haven't done any tuning.

Common problems like a substitutions followed by a deletion aren't currently
handled. e.g. "hyllo" -> "hell".
