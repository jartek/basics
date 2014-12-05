package heap_sort

import "testing"

type testpair struct {
	input  []int
	output []int
}

var tests = []testpair{
	{[]int{11, 3, 42, 1, 8}, []int{1, 3, 8, 11, 42}},
	{[]int{-8, 0, -1}, []int{-8, -1, 0}},
}

func TestSorting(t *testing.T) {
	for _, pair := range tests {
		sorted := HeapSort(pair.input)
		if !CompareSlices(sorted, pair.output) {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", sorted,
			)
		}
	}
}

func CompareSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
