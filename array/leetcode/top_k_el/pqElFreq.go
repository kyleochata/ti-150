package topkel

type element struct {
	value     int
	frequency int
}

type pq struct {
	heap     []*element
	elMap    map[int]int
	heapSize int
}

func (p *pq) insert(val int) {
	index, ok := p.elMap[val]
	if ok {
		p.heap[index].frequency += 1
		// check if the new frequency gets onto the heap
		p.bubble_up(index)
		return
	}
	el := &element{
		value:     val,
		frequency: 1,
	}
	p.elMap[val] = p.heapSize
	p.heapSize += 1
	p.heap = append(p.heap, el)
	p.bubble_up(p.elMap[val])
}

func (p *pq) bubble_up(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2

	if parent >= 0 && p.heap[parent].frequency < p.heap[index].frequency {
		p.swapAll(parent, index)
		p.bubble_up(parent)
	}
}

func (p *pq) swapAll(parent, index int) {
	p.swapMap(parent, index)
	p.swapHeap(parent, index)
}

func (p *pq) swapHeap(parent, index int) {
	p.heap[parent], p.heap[index] = p.heap[index], p.heap[parent]
}

func (p *pq) swapMap(parent, index int) {
	p.elMap[p.heap[parent].value], p.elMap[p.heap[index].value] = index, parent
}

func (p *pq) removeVal(val int) {
	index, ok := p.elMap[val]
	if !ok {
		return
	}
	//swap with end
	p.swapAll(index, p.heapSize-1)
	p.removeEnd(p.heap[p.heapSize-1].value)
	parent := (index) / 2
	lChild, rChild := (index)*2+1, (index-1)*2+2
	if parent >= 0 && p.heap[parent].value < p.heap[index].value {
		p.bubble_up(index)
		return
	}
	lGreaterThanR := p.heap[lChild].value > p.heap[rChild].value
	if lGreaterThanR {
		if lChild < p.heapSize && p.heap[index].value < p.heap[lChild].value {
			p.swapAll(index, lChild)
			p.bubble_down(lChild)
		}
	} else {
		if rChild < p.heapSize && p.heap[index].value < p.heap[rChild].value {
			p.swapAll(index, rChild)
			p.bubble_down(rChild)
		}
	}
	return
}

func (p *pq) removeEnd(val int) {
	delete(p.elMap, val)
	p.heap = p.heap[:p.heapSize-1]
	p.heapSize -= 1
}

func (p *pq) bubble_down(index int) {
	largest := index
	lChild := index*2 + 1
	rChild := index*2 + 2

	if lChild < p.heapSize && rChild < p.heapSize {
		if p.heap[lChild].value > p.heap[rChild].value {
			largest = lChild
		} else {
			largest = rChild
		}
	}
	if lChild < p.heapSize && rChild > p.heapSize {
		largest = lChild
	}
	if lChild > p.heapSize && rChild < p.heapSize {
		largest = rChild
	}
	if largest != index {
		p.swapAll(largest, index)
		p.bubble_down(largest)
	}
}
