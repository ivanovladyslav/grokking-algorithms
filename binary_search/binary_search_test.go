package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		array []int
		value int
		want  bool
	}{
		{[]int{0, 1, 2, 3, 4, 10, 20, 30, 50, 70, 90, 100}, 10, true},
		{[]int{0, 1, 2, 3, 4, 10, 20, 30, 50, 70, 90, 100}, 75, false},
		{[]int{0, 1, 2, 3, 4, 10, 20, 30, 50, 70, 90, 100}, 100, true},
	}

	for _, c := range cases {
		got := BinarySearch(c.array, c.value)

		if got != c.want {
			t.Errorf("TestBinarySearch(%v, %v) == %v, want %v", c.array, c.value, got, c.want)
		}
	}
}
