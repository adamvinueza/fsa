package fsa

type SymbolSet struct {
	symbols map[string]bool
}

func NewSymbolSet(ss []string) *SymbolSet {
	sSet := SymbolSet{symbols: make(map[string]bool)}
	for _, s := range ss {
		sSet.symbols[s] = true
	}
	return &sSet
}

func (sSet *SymbolSet) Iter() chan string {
	iter := make(chan string)
	go func() {
		for k, _ := range sSet.symbols {
			iter <- k
		}
		close(iter)
	}()
	return iter
}

func (sSet *SymbolSet) Add(s string) {
	sSet.symbols[s] = true
}

func (sSet *SymbolSet) Remove(s string) {
	delete(sSet.symbols, s)
}

func (sSet *SymbolSet) Copy() *SymbolSet {
	sSetCopy := NewSymbolSet([]string{})
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

func NewLanguage(ss []string) *Language {
	l := Language{
		Symbols: NewSymbolSet(ss),
	}
	return &l
}
