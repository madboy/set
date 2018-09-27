package set

// StrSet is an unordered collection of unique string elements
type StrSet map[string]bool

// Len returns the length of the set
func (s *StrSet) Len() int {
	return len(*s)
}

// NewStr returns an empty StrSet
func NewStr() StrSet {
	return make(map[string]bool)
}

// Values returns an array of values in set
func (s *StrSet) Values() []string {
	v := make([]string, 0, s.Len())
	for k := range *s {
		v = append(v, k)
	}
	return v
}

// NewFromStrArr returns a set filled with values in arr
func NewStrFromArr(arr []string) StrSet {
	s := NewStr()
	for _, v := range arr {
		s.Add(v)
	}
	return s
}

// Add value to the StrSet
func (s *StrSet) Add(value string) {
	if !(*s)[value] {
		(*s)[value] = true
	}
}

// Difference returns all values in s that aren't in o
func (s *StrSet) Difference(o *StrSet) StrSet {
	n := NewStr()
	for k := range *s {
		if _, ok := (*o)[k]; !ok {
			n.Add(k)
		}
	}
	return n
}
