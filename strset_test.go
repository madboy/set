package set

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewStrSet(t *testing.T) {
	s := NewStr()
	empty := []string{}
	got := s.Values()
	if !cmp.Equal(got, empty) {
		t.Errorf("Expected: %v, Got: %v", empty, got)
	}
}

func TestNewStrFromArr(t *testing.T) {
	tests := []struct {
		input    StrSet
		expected []string
	}{
		{
			input:    NewStrFromArr([]string{"a", "b", "c", "d", "al", "bal"}),
			expected: []string{"a", "al", "b", "bal", "c", "d"},
		},
		{
			input:    NewStrFromArr([]string{"a", "a", "a", "b", "ce", "ce"}),
			expected: []string{"a", "b", "ce"},
		},
	}
	for _, tt := range tests {
		got := tt.input.Values()
		// To be able to compare we need to sort values as compare takes order into account
		sort.Strings(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Expected: %v, Got: %v", tt.expected, got)
		}

		if tt.input.Len() != len(tt.expected) {
			t.Errorf("Unexpectd number of elements, expected, %d, got %d", len(tt.expected), tt.input.Len())
		}
	}
}

func TestStrAddAll(t *testing.T) {
	tests := []struct {
		s        StrSet
		input    []string
		expected []string
	}{
		{
			s:        NewStr(),
			input:    []string{"a", "b", "c", "d", "al", "bal"},
			expected: []string{"a", "al", "b", "bal", "c", "d"},
		},
		{
			s:        NewStr(),
			input:    []string{"a", "a", "a", "b", "ce", "ce"},
			expected: []string{"a", "b", "ce"},
		},
		{
			s:        NewStrFromArr([]string{"a", "b"}),
			input:    []string{"a", "a", "a", "b", "ce", "ce"},
			expected: []string{"a", "b", "ce"},
		},
	}
	for _, tt := range tests {
		tt.s.AddAll(tt.input...)
		got := tt.s.Values()
		sort.Strings(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Expected: %v, Got: %v", tt.expected, got)
		}

		if tt.s.Len() != len(tt.expected) {
			t.Errorf("Unexpectd number of elements, expected, %d, got %d", len(tt.expected), tt.s.Len())
		}
	}
}

func TestStrAdd(t *testing.T) {
	tests := []struct {
		s        StrSet
		input    string
		expected []string
	}{
		{
			s:        NewStrFromArr([]string{"art", "paint", "sky"}),
			input:    "turner",
			expected: []string{"art", "paint", "sky", "turner"},
		},
		{
			s:        NewStr(),
			input:    "turner",
			expected: []string{"turner"},
		},
	}

	for _, tc := range tests {
		tc.s.Add(tc.input)
		got := tc.s.Values()
		sort.Strings(got)
		if !cmp.Equal(got, tc.expected) {
			t.Errorf("Expected: %v, Got: %v", tc.expected, got)
		}
	}
}

func TestStrSetDifference(t *testing.T) {
	tests := []struct {
		a        StrSet
		b        StrSet
		expected []string
	}{
		{
			a:        NewStrFromArr([]string{"a", "i"}),
			b:        NewStrFromArr([]string{"h", "a", "g", "e", "i"}),
			expected: []string{"e", "g", "h"},
		},
		{
			a:        NewStrFromArr([]string{"i", "a"}),
			b:        NewStrFromArr([]string{"h", "a", "g", "e", "i"}),
			expected: []string{"e", "g", "h"},
		},
		{
			a:        NewStrFromArr([]string{"h", "a", "g", "e", "i"}),
			b:        NewStrFromArr([]string{"i", "a"}),
			expected: []string{},
		},
	}

	for _, tt := range tests {
		diff := tt.b.Difference(&tt.a)

		got := diff.Values()
		sort.Strings(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same. Expected %v got %v", tt.expected, got)
		}
	}
}

func TestStrSetUnion(t *testing.T) {
	tests := []struct {
		a        StrSet
		b        StrSet
		expected []string
	}{
		{
			a:        NewStrFromArr([]string{"a", "b", "c"}),
			b:        NewStrFromArr([]string{"d", "e", "f"}),
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			a:        NewStrFromArr([]string{"a", "b", "c"}),
			b:        NewStrFromArr([]string{"a", "b", "c", "d", "e", "f"}),
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			a:        NewStrFromArr([]string{"a"}),
			b:        NewStrFromArr([]string{}),
			expected: []string{"a"},
		},
		{
			a:        NewStrFromArr([]string{}),
			b:        NewStrFromArr([]string{}),
			expected: []string{},
		},
	}

	for _, tt := range tests {
		union := tt.a.Union(&tt.b)
		got := union.Values()
		sort.Strings(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same %v, got %v", tt.expected, got)
		}
	}
}

func TestStrSetIntersection(t *testing.T) {
	tests := []struct {
		a        StrSet
		b        StrSet
		expected []string
	}{
		{
			a:        NewStrFromArr([]string{"a", "b", "c"}),
			b:        NewStrFromArr([]string{"d", "e", "f"}),
			expected: []string{},
		},
		{
			a:        NewStrFromArr([]string{"a", "b", "c"}),
			b:        NewStrFromArr([]string{"a", "b", "c", "d", "e", "f"}),
			expected: []string{"a", "b", "c"},
		},
		{
			a:        NewStrFromArr([]string{"a", "b", "c", "d", "e", "f"}),
			b:        NewStrFromArr([]string{"a", "b", "c"}),
			expected: []string{"a", "b", "c"},
		},
		{
			a:        NewStrFromArr([]string{"a"}),
			b:        NewStrFromArr([]string{}),
			expected: []string{},
		},
		{
			a:        NewStrFromArr([]string{}),
			b:        NewStrFromArr([]string{}),
			expected: []string{},
		},
	}

	for _, tt := range tests {
		union := tt.a.Intersection(&tt.b)
		got := union.Values()
		sort.Strings(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same %v, got %v", tt.expected, got)
		}
	}
}

func TestStrSetSymmetricDifference(t *testing.T) {
	tests := []struct {
		a        StrSet
		b        StrSet
		expected []string
	}{
		{
			a:        NewStrFromArr([]string{"a", "b", "c"}),
			b:        NewStrFromArr([]string{"d", "e", "f"}),
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			a:        NewStrFromArr([]string{"a", "b", "c"}),
			b:        NewStrFromArr([]string{"a", "b", "c", "d", "e", "f"}),
			expected: []string{"d", "e", "f"},
		},
		{
			a:        NewStrFromArr([]string{"a", "b", "c", "d", "e", "f"}),
			b:        NewStrFromArr([]string{"a", "b", "c"}),
			expected: []string{"d", "e", "f"},
		},
		{
			a:        NewStrFromArr([]string{"a"}),
			b:        NewStrFromArr([]string{}),
			expected: []string{"a"},
		},
		{
			a:        NewStrFromArr([]string{}),
			b:        NewStrFromArr([]string{}),
			expected: []string{},
		},
		{
			a:        NewStrFromArr([]string{"a", "b", "c", "d"}),
			b:        NewStrFromArr([]string{"d", "e", "f", "g"}),
			expected: []string{"a", "b", "c", "e", "f", "g"},
		},
	}

	for _, tt := range tests {
		union := tt.a.SymmetricDifference(&tt.b)
		got := union.Values()
		sort.Strings(got)
		if !cmp.Equal(got, tt.expected) {
			t.Errorf("Values are not the same %v, got %v", tt.expected, got)
		}
	}
}
