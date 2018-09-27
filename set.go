package set

// IntSet is an unordered collection of unique int elements
type IntSet struct {
	keys map[int]bool
}

// StrSet is an unordered collection of unique string elements
type StrSet struct {
	keys map[string]bool
}

// Len returns the length of the set
func (s *IntSet) Len() int {
	return len(s.keys)
}

// NewInt returns an empty IntSet
func NewInt() IntSet {
	return IntSet{
		keys: make(map[int]bool),
	}
}

// NewIntFromArr returns a set filled with values in arr
func NewIntFromArr(arr []int) IntSet {
	s := NewInt()
	for _, v := range arr {
		s.Add(v)
	}
	return s
}

// Add value to the IntSet
func (s *IntSet) Add(value int) {
	if !s.keys[value] {
		s.keys[value] = true
	}
}

// Difference returns all values in s that aren't in o
func (s *IntSet) Difference(o *IntSet) IntSet {
	n := NewInt()
	for k := range s.keys {
		if _, ok := o.keys[k]; !ok {
			n.Add(k)
		}
	}
	return n
}

// Values returns an array of values in set
func (s *IntSet) Values() []int {
	v := make([]int, 0, len(s.keys))
	for k := range s.keys {
		v = append(v, k)
	}
	return v
}

// Len returns the length of the set
func (s *StrSet) Len() int {
	return len(s.keys)
}

// NewStr returns an empty StrSet
func NewStr() StrSet {
	return StrSet{
		keys: make(map[string]bool),
	}
}

// Values returns an array of values in set
func (s *StrSet) Values() []string {
	v := make([]string, 0, len(s.keys))
	for k := range s.keys {
		v = append(v, k)
	}
	return v
}

// NewFromStrArr returns a set filled with values in arr
func NewFromStrArr(arr []string) StrSet {
	s := NewStr()
	for _, v := range arr {
		s.Add(v)
	}
	return s
}

// Add value to the StrSet
func (s *StrSet) Add(value string) {
	if !s.keys[value] {
		s.keys[value] = true
	}
}

// Difference returns all values in s that aren't in o
func (s *StrSet) Difference(o *StrSet) StrSet {
	n := NewStr()
	for k := range s.keys {
		if _, ok := o.keys[k]; !ok {
			n.Add(k)
		}
	}
	return n
}
