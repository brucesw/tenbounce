package util

import (
	"fmt"
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
