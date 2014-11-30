package insertion_sort

func Sort(input []float64) []float64 {
	for i := 1; i < len(input); i++ {
		j := i - 1
		k := input[i]
		for j >= 0 && input[j] > k {
			input[j+1] = input[j]
			j = j - 1
		}
		input[j+1] = k
	}
	return input
}
