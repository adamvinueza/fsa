package fsa

// ExpressionSet represents a set of expressions, the building blocks of a Language.
type ExpressionSet struct {
	expressions map[string]bool
}

func NewEmptyExpressionSet() *ExpressionSet {
	return &ExpressionSet{expressions: make(map[string]bool)}
}

// NewExpressionSet creates a ExpressionSet from a slice of strings. It is permissive in
// that it allows the slice to contain copies.
func NewExpressionSet(ss []string) (*ExpressionSet, error) {
	eSet := NewEmptyExpressionSet()
	for _, s := range ss {
		eSet.expressions[s] = true
	}
	return eSet, nil
}

func (eSet *ExpressionSet) Add(s string) {
	eSet.expressions[s] = true
}

func (eSet *ExpressionSet) Remove(s string) {
	delete(eSet.expressions, s)
}

func (eSet *ExpressionSet) Copy() *ExpressionSet {
	eSetCopy := NewEmptyExpressionSet()
	for k, _ := range eSet.expressions {
		eSetCopy.Add(k)
	}
	return eSetCopy
}

func (ss *ExpressionSet) Union(ee ExpressionSet) *ExpressionSet {
	ssCopy := ss.Copy()
	for k, _ := range ee.expressions {
		ssCopy.Add(k)
	}
	return ssCopy
}

type Language struct {
	expressions *ExpressionSet
}

func NewLanguage(ss []string) (*Language, error) {
	expressions, err := NewExpressionSet(ss)
	if err != nil {
		return nil, err
	}
	l := Language{
		expressions: expressions,
	}
	return &l, nil
}
