package bubble_sort

func Swap(min *float64, current *float64) {
	tmp := *min
	*min = *current
	*current = tmp
}

func Sort(input []float64) []float64 {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				Swap(&input[j], &input[j+1])
			}
		}
	}

	return input
}
