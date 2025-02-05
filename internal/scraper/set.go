package scraper

type Set map[string]struct{}

type Broken map[string]int

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(val string) {
	s[val] = struct{}{}
}

func (s Set) Count() int {
	return len(s)
}

// Has checks if a string exists in the set.
func (s Set) Has(str string) bool {
	_, ok := s[str]
	return ok
}
