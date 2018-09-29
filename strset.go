package set

import (
	"bytes"
	"fmt"
)

// StrSet is an unordered collection of unique string elements
type StrSet map[string]bool

func (s *StrSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k := range *s {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "\"%s\"", k)
	}
	buf.WriteByte('}')
	return buf.String()
}

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

// NewStrFromArr returns a set filled with values in arr
func NewStrFromArr(arr []string) StrSet {
	s := NewStr()
	s.AddAll(arr...)
	return s
}

// Add value to the StrSet
func (s *StrSet) Add(value string) {
	if !(*s)[value] {
		(*s)[value] = true
	}
}

// AddAll adds all values to the set
func (s *StrSet) AddAll(values ...string) {
	for _, v := range values {
		s.Add(v)
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

// Union retuns a set of all values present in both s and o
func (s *StrSet) Union(o *StrSet) StrSet {
	v := make([]string, 0, s.Len()+o.Len())
	for k := range *s {
		v = append(v, k)
	}
	for k := range *o {
		v = append(v, k)
	}
	return NewStrFromArr(v)
}

// Intersection returns a set with values that are both in s and o
func (s *StrSet) Intersection(o *StrSet) StrSet {
	n := NewStr()
	for k := range *s {
		if _, ok := (*o)[k]; ok {
			n.Add(k)
		}
	}
	return n
}

// SymmetricDifference returns a set with values that are in s and o but not both
func (s *StrSet) SymmetricDifference(o *StrSet) StrSet {
	i := s.Intersection(o)
	d1 := s.Difference(&i)
	d2 := o.Difference(&i)
	return d1.Union(&d2)
}
