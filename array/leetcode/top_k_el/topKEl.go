package topkel

func topKEl(nums []int, k int) int {
	ax := &heap{
		heap: []int{},
		size: 0,
	}
	for _, num := range nums {
		ax.insert(num)
	}
	var kEl int
	for i := 0; i < k; i++ {
		cur := ax.removeMaxEl()
		if i == k-1 {
			kEl = cur
		}
	}
	return kEl
}

type heap struct {
	heap []int
	size int
}

func (h *heap) insert(val int) {
	h.heap = append(h.heap, val)
	h.size += 1
	h.bubble_up(h.size)
}

func (h *heap) bubble_up(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if parent >= 0 && index != 0 {
		if h.heap[index] > h.heap[parent] {
			h.heap[index], h.heap[parent] = h.heap[parent], h.heap[index]
			h.bubble_up(parent)
		}
	}
	return
}

func (h *heap) removeMaxEl() int {
	maxEl := h.heap[0]
	h.heap[0], h.heap[h.size-1] = h.heap[h.size-1], h.heap[0]
	h.heap = h.heap[:h.size-1]
	h.size -= 1
	h.bubble_down(0)
	return maxEl
}

func (h *heap) bubble_down(parent int) {
	leftChild := parent*2 + 1
	rightChild := parent*2 + 2
	largest := 0
	if leftChild < h.size && rightChild < h.size {
		largest = max(h.heap[leftChild], h.heap[rightChild])
	}
	if h.heap[parent] < largest {
		if largest == h.heap[leftChild] {
			h.swap(parent, leftChild)
			h.bubble_down(leftChild)
		} else {
			h.swap(parent, rightChild)
			h.bubble_down(rightChild)
		}
	}
}

func (h *heap) swap(parent, child int) {
	h.heap[parent], h.heap[child] = h.heap[child], h.heap[parent]
}
