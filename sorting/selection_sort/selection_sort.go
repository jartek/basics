package selection_sort

func Swap(min *float64, current *float64) {
	tmp := *min
	*min = *current
	*current = tmp
}

func Sort(input []float64) []float64 {
	for i := 0; i < len(input)-1; i++ {
		min := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		Swap(&input[min], &input[i])
	}
	return input
}
