package set

// IntSet is an unordered collection of unique int elements
type IntSet map[int]bool

// Len returns the length of the set
func (s *IntSet) Len() int {
	return len(*s)
}

// NewInt returns an empty IntSet
func NewInt() IntSet {
	return make(map[int]bool)
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
	if !(*s)[value] {
		(*s)[value] = true
	}
}

// Difference returns all values in s that aren't in o
func (s *IntSet) Difference(o *IntSet) IntSet {
	n := NewInt()
	for k := range *s {
		if _, ok := (*o)[k]; !ok {
			n.Add(k)
		}
	}
	return n
}

// Union retuns a set of all values present in both s and o
func (s *IntSet) Union(o *IntSet) IntSet {
	v := make([]int, 0, s.Len()+o.Len())
	for k := range *s {
		v = append(v, k)
	}
	for k := range *o {
		v = append(v, k)
	}
	return NewIntFromArr(v)
}

// Intersection returns a set with values that are both in s and o
func (s *IntSet) Intersection(o *IntSet) IntSet {
	n := NewInt()
	for k := range *s {
		if _, ok := (*o)[k]; ok {
			n.Add(k)
		}
	}
	return n
}

// SymmetricDifference returns a set with values that are in s and o but not both
func (s *IntSet) SymmetricDifference(o *IntSet) IntSet {
	i := s.Intersection(o)
	d1 := s.Difference(&i)
	d2 := o.Difference(&i)
	return d1.Union(&d2)
}

// Values returns an array of values in set
func (s *IntSet) Values() []int {
	v := make([]int, 0, s.Len())
	for k := range *s {
		v = append(v, k)
	}
	return v
}
