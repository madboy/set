package set

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewInt(t *testing.T) {
	s := NewInt()
	empty := []int{}
	got := s.Values()
	if !cmp.Equal(got, empty) {
		t.Errorf("Expected: %v, Got: %v", empty, got)
	}
}

func TestNewIntFromArr(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
			expected: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
		},
		{
			input:    []int{1, 1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		s := NewIntFromArr(tt.input)
		got := s.Values()
		// To be able to compare we need to sort values as compare takes order into account
		sort.Ints(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Expected: %v, Got: %v", tt.expected, got)
		}

		if s.Len() != len(tt.expected) {
			t.Errorf("Unexpectd number of elements, expected, %d, got %d", len(tt.expected), s.Len())
		}
	}
}

func TestIntAdd(t *testing.T) {
	s := NewIntFromArr([]int{1, 2, 3})
	s.Add(5)
	expected := []int{1, 2, 3, 5}

	got := s.Values()
	sort.Ints(got)
	if !cmp.Equal(got, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}

func TestIntSetDifference(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected []int
	}{
		{
			a:        []int{9, 11},
			b:        []int{7, 8, 9, 10, 11},
			expected: []int{7, 8, 10},
		},
		{
			a:        []int{11, 9},
			b:        []int{7, 8, 9, 10, 11},
			expected: []int{7, 8, 10},
		},
		{
			a:        []int{7, 8, 9, 10, 11},
			b:        []int{9, 11},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		a := NewIntFromArr(tt.a)
		b := NewIntFromArr(tt.b)
		diff := b.Difference(&a)

		got := diff.Values()
		sort.Ints(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same. Expected %v got %v", tt.expected, got)
		}
	}

}

func TestNewStrSet(t *testing.T) {
	s := NewStr()
	empty := []string{}
	got := s.Values()
	if !cmp.Equal(got, empty) {
		t.Errorf("Expected: %v, Got: %v", empty, got)
	}
}

func TestNewStrSetFromArr(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c", "d", "al", "bal"},
			expected: []string{"a", "al", "b", "bal", "c", "d"},
		},
		{
			input:    []string{"a", "a", "a", "b", "ce", "ce"},
			expected: []string{"a", "b", "ce"},
		},
	}
	for _, tt := range tests {
		s := NewFromStrArr(tt.input)
		got := s.Values()
		// To be able to compare we need to sort values as compare takes order into account
		sort.Strings(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Expected: %v, Got: %v", tt.expected, got)
		}

		if s.Len() != len(tt.expected) {
			t.Errorf("Unexpectd number of elements, expected, %d, got %d", len(tt.expected), s.Len())
		}
	}
}

func TestStrAdd(t *testing.T) {
	s := NewFromStrArr([]string{"art", "paint", "sky"})
	s.Add("turner")
	expected := []string{"art", "paint", "sky", "turner"}

	got := s.Values()
	sort.Strings(got)
	if !cmp.Equal(got, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}

func TestStrSetDifference(t *testing.T) {
	tests := []struct {
		a        []string
		b        []string
		expected []string
	}{
		{
			a:        []string{"a", "i"},
			b:        []string{"h", "a", "g", "e", "i"},
			expected: []string{"e", "g", "h"},
		},
		{
			a:        []string{"i", "a"},
			b:        []string{"h", "a", "g", "e", "i"},
			expected: []string{"e", "g", "h"},
		},
		{
			a:        []string{"h", "a", "g", "e", "i"},
			b:        []string{"i", "a"},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		a := NewFromStrArr(tt.a)
		b := NewFromStrArr(tt.b)
		diff := b.Difference(&a)

		got := diff.Values()
		sort.Strings(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same. Expected %v got %v", tt.expected, got)
		}
	}
}
