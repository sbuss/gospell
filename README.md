gospell
=======

Spelling suggestion in Go

Usage
=====

```go
package main

import (
	"fmt"
	"github.com/sbuss/gospell"
)

func main() {
	trie, err := gospell.TrieFromFile("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}
	suggestions := trie.SuggestWords("gospell", 2)
	fmt.Println(suggestions)
	//[gospelly bespell byspell enspell fostell gospel respell unspell spell]
}
```

Warnings & Caveats
==================
The spelling suggestions are rudimentary and only rank by distance and
lexicographic order. Support for weighting words differently is TBD.

Memory and performance both seem ok, but I haven't done any tuning.

Common problems like a substitution followed by a deletion aren't currently
handled. e.g. "hyllo" -> "hell".
