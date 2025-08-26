package util

import (
	"fmt"
	"reflect"
	"testing"
)

type testContains_Type[T comparable] struct {
	have   []T
	want   T
	expect bool
}

func TestContains_string(t *testing.T) {
	var battery = []testContains_Type[string]{
		{
			have:   []string{},
			want:   "a",
			expect: false,
		},
		{
			have:   []string{"a"},
			want:   "a",
			expect: true,
		},
		{
			have:   []string{"a", "a"},
			want:   "a",
			expect: true,
		},
		{
			have:   []string{"b"},
			want:   "a",
			expect: false,
		},
		{
			have:   []string{"a", "b", "c"},
			want:   "b",
			expect: true,
		},
	}

	for _, tt := range battery {
		if Contains(tt.have, tt.want) != tt.expect {
			t.Error(fmt.Errorf("have: %+v, want: %+v, expect: %t", tt.have, tt.want, tt.expect))
		}
	}
}

func TestContains_int(t *testing.T) {
	var battery = []testContains_Type[int]{
		{
			have:   []int{},
			want:   1,
			expect: false,
		},
		{
			have:   []int{1},
			want:   1,
			expect: true,
		},
		{
			have:   []int{1, 1},
			want:   1,
			expect: true,
		},
		{
			have:   []int{2},
			want:   1,
			expect: false,
		},
		{
			have:   []int{1, 2, 3},
			want:   2,
			expect: true,
		},
	}

	for _, tt := range battery {
		if Contains(tt.have, tt.want) != tt.expect {
			t.Error(fmt.Errorf("have: %+v, want: %+v, expect: %t", tt.have, tt.want, tt.expect))
		}
	}
}

type testMap_Type[T comparable] struct {
	input    []T
	expected []T
	fun      func(T) T
}

func Test_Map_Ints(t *testing.T) {
	var intCases = []testMap_Type[int]{
		{
			input:    []int{1, 2, 3},
			expected: []int{2, 4, 6},
			fun:      func(i int) int { return 2 * i },
		}, {
			input:    []int{1, 2, 3},
			expected: []int{1, 1, 1},
			fun:      func(i int) int { return 1 },
		}, {
			input:    []int{10, 20, 30},
			expected: []int{11, 21, 31},
			fun:      func(i int) int { return i + 1 },
		},
	}

	for _, intCase := range intCases {
		var actual = Map(intCase.input, intCase.fun)
		if !reflect.DeepEqual(actual, intCase.expected) {
			t.Errorf("int case: got %v, expected %v", actual, intCase.expected)
		}
	}
}

func Test_Map_Strs(t *testing.T) {
	var strCases = []testMap_Type[string]{
		{
			input:    []string{"1", "2", "3"},
			expected: []string{"1", "2", "3"},
			fun:      func(s string) string { return s },
		}, {
			input:    []string{"1", "2", "3"},
			expected: []string{"s", "s", "s"},
			fun:      func(s string) string { return "s" },
		}, {
			input:    []string{"10", "20", "30"},
			expected: []string{"101", "201", "301"},
			fun:      func(s string) string { return fmt.Sprintf("%s1", s) },
		},
	}

	for _, strCase := range strCases {
		var actual = Map(strCase.input, strCase.fun)
		if !reflect.DeepEqual(actual, strCase.expected) {
			t.Errorf("int case: got %v, expected %v", actual, strCase.expected)
		}
	}
}
