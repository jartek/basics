package merge_sort

import "testing"

type testpair struct {
	input  []float64
	output []float64
}

var tests = []testpair{
	{[]float64{11, 3, 42, 1, 8}, []float64{1, 3, 8, 11, 42}},
	{[]float64{-8, 0, -1}, []float64{-8, -1, 0}},
}

func TestSorting(t *testing.T) {
	for _, pair := range tests {
		sorted := MergeSort(pair.input)
		if !CompareSlices(sorted, pair.output) {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", sorted,
			)
		}
	}
}

func CompareSlices(a []float64, b []float64) bool {
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
