package merge_sort

func Merge(array []float64, left int, mid int, right int) {
	m1 := mid - left + 1
	m2 := right - mid
	tmp1 := make([]float64, m1)
	tmp2 := make([]float64, m2)
	var i, j, k int

	for i := 0; i < m1; i++ {
		tmp1[i] = array[left+i]
	}

	for j := 0; j < m2; j++ {
		tmp2[j] = array[mid+j+1]
	}

	k = left
	i, j = 0, 0

	for i < m1 && j < m2 {
		if tmp1[i] <= tmp2[j] {
			array[k] = tmp1[i]
			i = i + 1
		} else {
			array[k] = tmp2[j]
			j = j + 1
		}
		k = k + 1
	}

	for i < m1 {
		array[k] = tmp1[i]
		i = i + 1
		k = k + 1
	}

	for j < m2 {
		array[k] = tmp2[j]
		j = j + 1
		k = k + 1
	}
}

func Split(input []float64, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		Split(input, left, mid)
		Split(input, mid+1, right)
		Merge(input, left, mid, right)
	}
}

func MergeSort(input []float64) []float64 {
	Split(input, 0, len(input)-1)
	return input
}
