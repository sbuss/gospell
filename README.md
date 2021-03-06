gospell
=======

Spelling suggestion in Go. Current version is
[gospell v0.1.0](https://github.com/sbuss/gospell/tarball/v0.1.0).

Installation
============
Your standard `go get` installation.

```sh
go get github.com/sbuss/gospell
```

Testing
-------
```sh
cd $GOPATH/src/github.com/sbuss/gospell
go test
go test -bench=".*"
```

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

Changelog
=========
* [v0.1.0](https://github.com/sbuss/gospell/tarball/v0.1.0) --
Initial release. Supports additions, deletions, and substitutions
for a given distance, and all valid permutations (support for distance is
planned). Can load a newline-delimited list of words to build a dict and can
suggest alternate spellings for strings within a given distance.

Warnings & Caveats
==================
The spelling suggestions are rudimentary and only rank by distance and
lexicographic order. Support for weighting words differently is TBD.

Memory seems ok, but I haven't done any tuning. Performance needs a lot of
work. `go test -bench=".*"` to see the current performance characteristics.

Common problems like a substitution followed by a deletion aren't currently
handled. e.g. "hyllo" -> "hell".

License
=======

[BSD 3-Clause](http://opensource.org/licenses/BSD-3-Clause), same as the
[Go license](http://golang.org/LICENSE).
