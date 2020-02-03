package fsa

// SymbolSet represents a set of symbols, the building blocks of a Language.
type SymbolSet struct {
	symbols map[string]bool
}

func NewEmptySymbolSet() *SymbolSet {
	return &SymbolSet{symbols: make(map[string]bool)}
}

// NewSymbolSet creates a SymbolSet from a slice of strings. It is permissive in
// that it allows the slice to contain copies.
func NewSymbolSet(ss []string) (*SymbolSet, error) {
	sSet := NewEmptySymbolSet()
	for _, s := range ss {
		sSet.symbols[s] = true
	}
	return sSet, nil
}

func (sSet *SymbolSet) Add(s string) {
	sSet.symbols[s] = true
}

func (sSet *SymbolSet) Remove(s string) {
	delete(sSet.symbols, s)
}

func (sSet *SymbolSet) Copy() *SymbolSet {
	sSetCopy := NewEmptySymbolSet()
	for k, _ := range sSet.symbols {
		sSetCopy.Add(k)
	}
	return sSetCopy
}

func (ss *SymbolSet) Union(tt SymbolSet) *SymbolSet {
	ssCopy := ss.Copy()
	for k, _ := range tt.symbols {
		ssCopy.Add(k)
	}
	return ssCopy
}

type Language struct {
	Symbols *SymbolSet
}

func NewLanguage(ss []string) (*Language, error) {
	symbols, err := NewSymbolSet(ss)
	if err != nil {
		return nil, err
	}
	l := Language{
		Symbols: symbols,
	}
	return &l, nil
}
