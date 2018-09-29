package set

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestNewStrSet(t *testing.T) {
	s := NewStr()
	empty := []string{}
	got := s.Elements()
	if !equalStr(got, empty) {
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
		got := tt.input.Elements()
		if !equalStr(got, tt.expected) {
			t.Errorf("Expected: %v, Got: %v", tt.expected, got)
		}

		if tt.input.Len() != len(tt.expected) {
			t.Errorf("Unexpectd number of elements, expected, %d, got %d", len(tt.expected), tt.input.Len())
		}
	}
}

func TestStrString(t *testing.T) {
	tests := []struct {
		s        StrSet
		expected int
	}{
		{
			s:        NewStrFromArr([]string{"a", "b", "c"}),
			expected: 3,
		},
		{
			s:        NewStr(),
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
		got := tt.s.Elements()
		if !equalStr(got, tt.expected) {
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
		got := tc.s.Elements()
		if !equalStr(got, tc.expected) {
			t.Errorf("Expected: %v, Got: %v", tc.expected, got)
		}
	}
}

func TestStrRemove(t *testing.T) {
	tests := []struct {
		s        StrSet
		input    string
		expected []string
	}{
		{
			s:        NewStrFromArr([]string{"a", "b", "c"}),
			input:    "a",
			expected: []string{"b", "c"},
		},
		{
			s:        NewStr(),
			input:    "p",
			expected: []string{},
		},
	}

	for _, tc := range tests {
		tc.s.Remove(tc.input)
		got := tc.s.Elements()
		if !equalStr(got, tc.expected) {
			t.Errorf("Expected: %v, Got: %v", tc.expected, got)
		}
	}
}

func TestStrRemoveAll(t *testing.T) {
	tests := []struct {
		s        StrSet
		input    []string
		expected []string
	}{
		{
			s:        NewStrFromArr([]string{"a", "b", "c", "d", "e"}),
			input:    []string{"a", "d", "po"},
			expected: []string{"b", "c", "e"},
		},
		{
			s:        NewStrFromArr([]string{"a", "b", "c"}),
			input:    []string{"f"},
			expected: []string{"a", "b", "c"},
		},
		{
			s:        NewStrFromArr([]string{"a", "b", "c"}),
			input:    []string{"a", "b", "c"},
			expected: []string{},
		},
	}

	for _, tc := range tests {
		tc.s.RemoveAll(tc.input...)
		got := tc.s.Elements()
		if !equalStr(got, tc.expected) {
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

		got := diff.Elements()
		if !equalStr(got, tt.expected) {
			t.Errorf("Elements are not the same. Expected %v got %v", tt.expected, got)
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
		got := union.Elements()
		if !equalStr(got, tt.expected) {
			t.Errorf("Elements are not the same %v, got %v", tt.expected, got)
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
		got := union.Elements()
		if !equalStr(got, tt.expected) {
			t.Errorf("Elements are not the same %v, got %v", tt.expected, got)
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
		got := union.Elements()
		if !equalStr(got, tt.expected) {
			t.Errorf("Elements are not the same %v, got %v", tt.expected, got)
		}
	}
}

// equalStr compares equality of two sorted arrays
func equalStr(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	sort.Strings(x)
	sort.Strings(y)

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
