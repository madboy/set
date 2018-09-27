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
		input    IntSet
		expected []int
	}{
		{
			input:    NewIntFromArr([]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}),
			expected: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
		},
		{
			input:    NewIntFromArr([]int{1, 1, 1, 2, 3, 3}),
			expected: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		got := tt.input.Values()
		// To be able to compare we need to sort values as compare takes order into account
		sort.Ints(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Expected: %v, Got: %v", tt.expected, got)
		}

		if tt.input.Len() != len(tt.expected) {
			t.Errorf("Unexpectd number of elements, expected, %d, got %d", len(tt.expected), tt.input.Len())
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
		a        IntSet
		b        IntSet
		expected []int
	}{
		{
			a:        NewIntFromArr([]int{9, 11}),
			b:        NewIntFromArr([]int{7, 8, 9, 10, 11}),
			expected: []int{7, 8, 10},
		},
		{
			a:        NewIntFromArr([]int{11, 9}),
			b:        NewIntFromArr([]int{7, 8, 9, 10, 11}),
			expected: []int{7, 8, 10},
		},
		{
			a:        NewIntFromArr([]int{7, 8, 9, 10, 11}),
			b:        NewIntFromArr([]int{9, 11}),
			expected: []int{},
		},
	}

	for _, tt := range tests {
		diff := tt.b.Difference(&tt.a)

		got := diff.Values()
		sort.Ints(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same. Expected %v got %v", tt.expected, got)
		}
	}

}

func TestIntSetUnion(t *testing.T) {
	tests := []struct {
		a        IntSet
		b        IntSet
		expected []int
	}{
		{
			a:        NewIntFromArr([]int{1, 2, 3}),
			b:        NewIntFromArr([]int{4, 5, 6}),
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			a:        NewIntFromArr([]int{1, 2, 3}),
			b:        NewIntFromArr([]int{1, 2, 3, 4, 5, 6}),
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			a:        NewIntFromArr([]int{1}),
			b:        NewIntFromArr([]int{}),
			expected: []int{1},
		},
		{
			a:        NewIntFromArr([]int{}),
			b:        NewIntFromArr([]int{}),
			expected: []int{},
		},
	}

	for _, tt := range tests {
		union := tt.a.Union(&tt.b)
		got := union.Values()
		sort.Ints(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same %v, got %v", tt.expected, got)
		}
	}
}

func TestIntSetIntersection(t *testing.T) {
	tests := []struct {
		a        IntSet
		b        IntSet
		expected []int
	}{
		{
			a:        NewIntFromArr([]int{1, 2, 3}),
			b:        NewIntFromArr([]int{4, 5, 6}),
			expected: []int{},
		},
		{
			a:        NewIntFromArr([]int{1, 2, 3}),
			b:        NewIntFromArr([]int{1, 2, 3, 4, 5, 6}),
			expected: []int{1, 2, 3},
		},
		{
			a:        NewIntFromArr([]int{1, 2, 3, 4, 5, 6}),
			b:        NewIntFromArr([]int{1, 2, 3}),
			expected: []int{1, 2, 3},
		},
		{
			a:        NewIntFromArr([]int{1}),
			b:        NewIntFromArr([]int{}),
			expected: []int{},
		},
		{
			a:        NewIntFromArr([]int{}),
			b:        NewIntFromArr([]int{}),
			expected: []int{},
		},
	}

	for _, tt := range tests {
		union := tt.a.Intersection(&tt.b)
		got := union.Values()
		sort.Ints(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same %v, got %v", tt.expected, got)
		}
	}
}
