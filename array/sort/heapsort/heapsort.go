package heapsort

type MaxHeap struct {
	heap     []int
	heapSize int
}

func (h *MaxHeap) Heapify(length, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && h.heap[left] > h.heap[largest] {
		largest = left
	}
	if right < length && h.heap[right] > h.heap[largest] {
		largest = right
	}
	if largest != i {
		h.heap[i], h.heap[largest] = h.heap[largest], h.heap[i]
		h.Heapify(length, largest)
	}
}
func (h *MaxHeap) Sort() {
	length := len(h.heap)
	for i := length/2 - 1; i >= 0; i-- {
		h.Heapify(length, i)
	}

	for i := length - 1; i >= 0; i-- {
		h.heap[0], h.heap[i] = h.heap[i], h.heap[0]
		h.Heapify(i, 0)
	}
}
