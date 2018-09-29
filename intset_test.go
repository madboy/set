package set

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestNewInt(t *testing.T) {
	s := NewInt()
	empty := []int{}
	got := s.Values()
	if !equalInt(got, empty) {
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
		if !equalInt(got, tt.expected) {
			t.Errorf("Expected: %v, Got: %v", tt.expected, got)
		}

		if tt.input.Len() != len(tt.expected) {
			t.Errorf("Unexpectd number of elements, expected, %d, got %d", len(tt.expected), tt.input.Len())
		}
	}
}

func TestIntString(t *testing.T) {
	tests := []struct {
		s        IntSet
		expected int
	}{
		{
			s:        NewIntFromArr([]int{1, 2, 3}),
			expected: 3,
		},
		{
			s:        NewInt(),
			expected: 1,
		},
	}

	for _, tc := range tests {
		got := fmt.Sprintf("%s", &tc.s)
		if got[0] != '{' && got[len(got)-1] != '}' {
			t.Errorf("Expected string to start with { and end with }, got %s", got)
		}

		parts := strings.Split(got, " ")
		if len(parts) != tc.expected {
			t.Errorf("Wrong number of elements in output, expected %d, got %d (%s)", tc.expected, len(parts), got)
		}
	}
}
func TestIntAddAll(t *testing.T) {
	tests := []struct {
		s        IntSet
		input    []int
		expected []int
	}{
		{
			s:        NewInt(),
			input:    []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
			expected: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
		},
		{
			s:        NewInt(),
			input:    []int{1, 1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			s:        NewIntFromArr([]int{1, 2, 4}),
			input:    []int{1, 1, 1, 2, 3, 3},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		tt.s.AddAll(tt.input...)
		got := tt.s.Values()
		if !equalInt(got, tt.expected) {
			t.Errorf("Expected: %v, Got: %v", tt.expected, got)
		}

		if tt.s.Len() != len(tt.expected) {
			t.Errorf("Unexpectd number of elements, expected %d, got %d", len(tt.expected), tt.s.Len())
		}
	}
}

func TestIntAdd(t *testing.T) {
	tests := []struct {
		s        IntSet
		input    int
		expected []int
	}{
		{
			s:        NewIntFromArr([]int{1, 2, 3}),
			input:    5,
			expected: []int{1, 2, 3, 5},
		},
		{
			s:        NewInt(),
			input:    5,
			expected: []int{5},
		},
	}

	for _, tc := range tests {
		tc.s.Add(tc.input)
		got := tc.s.Values()
		if !equalInt(got, tc.expected) {
			t.Errorf("Expected: %v, Got: %v", tc.expected, got)
		}
	}
}

func TestIntRemove(t *testing.T) {
	tests := []struct {
		s        IntSet
		input    int
		expected []int
	}{
		{
			s:        NewIntFromArr([]int{1, 2, 3}),
			input:    3,
			expected: []int{1, 2},
		},
		{
			s:        NewInt(),
			input:    5,
			expected: []int{},
		},
	}

	for _, tc := range tests {
		tc.s.Remove(tc.input)
		got := tc.s.Values()
		if !equalInt(got, tc.expected) {
			t.Errorf("Expected: %v, Got: %v", tc.expected, got)
		}
	}
}

func TestIntRemoveAll(t *testing.T) {
	tests := []struct {
		s        IntSet
		input    []int
		expected []int
	}{
		{
			s:        NewIntFromArr([]int{1, 2, 3, 5, 6}),
			input:    []int{3, 1, 72},
			expected: []int{2, 5, 6},
		},
		{
			s:        NewIntFromArr([]int{1, 2, 3}),
			input:    []int{5},
			expected: []int{1, 2, 3},
		},
		{
			s:        NewIntFromArr([]int{1, 2, 3}),
			input:    []int{1, 2, 3},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		tc.s.RemoveAll(tc.input...)
		got := tc.s.Values()
		if !equalInt(got, tc.expected) {
			t.Errorf("Expected: %v, Got: %v", tc.expected, got)
		}
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
		if !equalInt(got, tt.expected) {
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
		if !equalInt(got, tt.expected) {
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
		if !equalInt(got, tt.expected) {
			t.Errorf("Values are not the same %v, got %v", tt.expected, got)
		}
	}
}

func TestIntSetSymmetricDifference(t *testing.T) {
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
			expected: []int{4, 5, 6},
		},
		{
			a:        NewIntFromArr([]int{1, 2, 3, 4, 5, 6}),
			b:        NewIntFromArr([]int{1, 2, 3}),
			expected: []int{4, 5, 6},
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
		{
			a:        NewIntFromArr([]int{1, 2, 3, 4}),
			b:        NewIntFromArr([]int{4, 5, 6, 7}),
			expected: []int{1, 2, 3, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		union := tt.a.SymmetricDifference(&tt.b)
		got := union.Values()
		if !equalInt(got, tt.expected) {
			t.Errorf("Values are not the same %v, got %v", tt.expected, got)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	s := NewInt()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkDifference(b *testing.B) {
	s1 := NewIntFromArr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	s2 := NewIntFromArr([]int{8, 9, 10, 11, 12, 13, 14, 15, 16})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s1.Difference(&s2)
	}
}

func BenchmarkUnion(b *testing.B) {
	s1 := NewIntFromArr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	s2 := NewIntFromArr([]int{8, 9, 10, 11, 12, 13, 14, 15, 16})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s1.Union(&s2)
	}
}

func BenchmarkIntersection(b *testing.B) {
	s1 := NewIntFromArr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	s2 := NewIntFromArr([]int{8, 9, 10, 11, 12, 13, 14, 15, 16})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s1.Intersection(&s2)
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	s1 := NewIntFromArr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	s2 := NewIntFromArr([]int{8, 9, 10, 11, 12, 13, 14, 15, 16})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s1.SymmetricDifference(&s2)
	}
}

// equalInt compares equality of two sorted arrays
func equalInt(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	sort.Ints(x)
	sort.Ints(y)

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
