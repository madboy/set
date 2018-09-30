package set

import (
	"bytes"
	"fmt"
)

// IntSet is an unordered collection of unique int elements
type IntSet struct {
	els map[int]bool
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for el := range s.els {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", el)
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the length of the set
func (s *IntSet) Len() int {
	return len(s.els)
}

// NewInt returns an empty IntSet
func NewInt() IntSet {
	s := IntSet{
		els: make(map[int]bool),
	}
	return s
}

// NewIntFromArr returns a set filled with elements in arr
func NewIntFromArr(arr []int) IntSet {
	s := NewInt()
	s.AddAll(arr...)
	return s
}

// Add element to the IntSet
func (s *IntSet) Add(element int) {
	s.els[element] = true
}

// AddAll adds all elements to the set
func (s *IntSet) AddAll(elements ...int) {
	for _, el := range elements {
		s.Add(el)
	}
}

// Remove element from set
func (s *IntSet) Remove(element int) {
	delete(s.els, element)
}

// RemoveAll elements from set
func (s *IntSet) RemoveAll(elements ...int) {
	for _, v := range elements {
		s.Remove(v)
	}
}

// Difference returns all elements in s that aren't in o
func (s *IntSet) Difference(o *IntSet) IntSet {
	n := NewInt()
	for el := range s.els {
		if _, ok := o.els[el]; !ok {
			n.Add(el)
		}
	}
	return n
}

// Union returns a set of all elements present in both s and o
func (s *IntSet) Union(o *IntSet) IntSet {
	us := NewInt()
	for el := range s.els {
		us.Add(el)
	}
	for el := range o.els {
		us.Add(el)
	}
	return us
}

// Intersection returns a set with elements that are both in s and o
func (s *IntSet) Intersection(o *IntSet) IntSet {
	n := NewInt()
	for el := range s.els {
		if _, ok := o.els[el]; ok {
			n.Add(el)
		}
	}
	return n
}

// SymmetricDifference returns a set with elements that are in s and o but not both
func (s *IntSet) SymmetricDifference(o *IntSet) IntSet {
	d1 := s.Difference(o)
	d2 := o.Difference(s)
	return d1.Union(&d2)
}

// Elements returns an array of elements in set
func (s *IntSet) Elements() []int {
	els := make([]int, 0, s.Len())
	for el := range s.els {
		els = append(els, el)
	}
	return els
}
