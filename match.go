package gospell

import (
	"fmt"
	"sort"
)

type Match struct {
	Word     []rune
	Distance int
	Weight   int
}

type Matches []Match

func (m1 Match) Equal(m2 Match) bool {
	return string(m1.Word) == string(m2.Word) &&
		m1.Distance == m2.Distance && m1.Weight == m2.Weight
}

func (m Match) String() string {
	return fmt.Sprintf("{%v %d %d}", string(m.Word), m.Distance, m.Weight)
}

func (s Matches) Len() int      { return len(s) }
func (s Matches) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Convert Matches to a list of strings
func (s Matches) Strings() []string {
	strings := make([]string, len(s))
	sort.Sort(ByDistance{s})
	for i, r := range s {
		strings[i] = string(r.Word)
	}

	return strings
}

type ByDistance struct {
	Matches
}

func (s ByDistance) Less(i, j int) bool {
	d1 := s.Matches[i].Distance
	d2 := s.Matches[j].Distance
	if d1 == d2 {
		return string(s.Matches[i].Word) < string(s.Matches[j].Word)
	}
	return d1 < d2
}

func (m *Match) update(r rune, distance, weight int) Match {
	m2 := Match{}
	if r != 0 {
		m2.Word = []rune{r}
		m2.Word = append(m2.Word, m.Word...)
	} else {
		m2.Word = m.Word[:]
	}
	m2.Distance = m.Distance + distance
	m2.Weight = m.Weight + weight
	return m2
}
