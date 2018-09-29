package set

import (
	"bytes"
	"fmt"
)

// StrSet is an unordered collection of unique string elements
type StrSet struct {
	els map[string]bool
}

func (s *StrSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k := range s.els {
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
	return len(s.els)
}

// NewStr returns an empty StrSet
func NewStr() StrSet {
	s := StrSet{
		els: make(map[string]bool),
	}
	return s
}

// Elements returns an array of elements in set
func (s *StrSet) Elements() []string {
	v := make([]string, 0, s.Len())
	for k := range s.els {
		v = append(v, k)
	}
	return v
}

// NewStrFromArr returns a set filled with elements in arr
func NewStrFromArr(arr []string) StrSet {
	s := NewStr()
	s.AddAll(arr...)
	return s
}

// Add element to the StrSet
func (s *StrSet) Add(element string) {
	if !s.els[element] {
		s.els[element] = true
	}
}

// AddAll adds all elements to the set
func (s *StrSet) AddAll(elements ...string) {
	for _, el := range elements {
		s.Add(el)
	}
}

// Remove element from set
func (s *StrSet) Remove(element string) {
	delete(s.els, element)
}

// RemoveAll elements from set
func (s *StrSet) RemoveAll(elements ...string) {
	for _, el := range elements {
		s.Remove(el)
	}
}

// Difference returns all elements in s that aren't in o
func (s *StrSet) Difference(o *StrSet) StrSet {
	n := NewStr()
	for el := range s.els {
		if _, ok := o.els[el]; !ok {
			n.Add(el)
		}
	}
	return n
}

// Union returns a set of all elements present in both s and o
func (s *StrSet) Union(o *StrSet) StrSet {
	us := NewStr()
	for el := range s.els {
		us.Add(el)
	}
	for el := range o.els {
		us.Add(el)
	}
	return us
}

// Intersection returns a set with elements that are both in s and o
func (s *StrSet) Intersection(o *StrSet) StrSet {
	n := NewStr()
	for el := range s.els {
		if _, ok := o.els[el]; ok {
			n.Add(el)
		}
	}
	return n
}

// SymmetricDifference returns a set with elements that are in s and o but not both
func (s *StrSet) SymmetricDifference(o *StrSet) StrSet {
	d1 := s.Difference(o)
	d2 := o.Difference(s)
	return d1.Union(&d2)
}
