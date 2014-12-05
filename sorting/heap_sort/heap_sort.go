package heap_sort

type Heap struct {
	size     int
	elements []int
}

func Swap(min *int, current *int) {
	tmp := *min
	*min = *current
	*current = tmp
}

func Heapify(heap *Heap, rootIdx int) {
	largestIdx := rootIdx
	lNodeIdx := rootIdx*2 + 1
	rNodeIdx := rootIdx*2 + 2

	if lNodeIdx < heap.size && heap.elements[lNodeIdx] > heap.elements[rootIdx] {
		largestIdx = lNodeIdx
	}

	if rNodeIdx < heap.size && heap.elements[rNodeIdx] > heap.elements[largestIdx] {
		largestIdx = rNodeIdx
	}

	if largestIdx != rootIdx {
		Swap(&heap.elements[largestIdx], &heap.elements[rootIdx])
		Heapify(heap, largestIdx)
	}
}

func BuildHeap(input []int) *Heap {
	size := len(input)
	var heap Heap
	heap.size = size
	heap.elements = input
	for i := (heap.size - 2) / 2; i >= 0; i-- {
		Heapify(&heap, i)
	}
	return &heap
}

func HeapSort(input []int) []int {
	heap := BuildHeap(input)
	for heap.size > 1 {
		Swap(&heap.elements[0], &heap.elements[heap.size-1])
		heap.size--
		Heapify(heap, 0)
	}
	return input
}
