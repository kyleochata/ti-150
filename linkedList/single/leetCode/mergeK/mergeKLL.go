package mergek

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	minQ := &minH{
		h:     []*ListNode{},
		hSize: 0,
	}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			minQ.insert(lists[i])
		}
	}
	head := &ListNode{}
	current := head
	for minQ.hSize > 0 {
		temp := minQ.remove()
		current.Next = temp
		current = current.Next

		if temp.Next != nil {
			minQ.insert(temp.Next)
		}
	}
	return head.Next
}

type minH struct {
	h     []*ListNode
	hSize int
}

func (m *minH) insert(n *ListNode) {
	m.h = append(m.h, n)
	m.hSize += 1
	m.bubble_up(m.hSize - 1)
}
func (m *minH) bubble_up(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if parent >= 0 && m.h[parent].Val > m.h[index].Val {
		m.swap(parent, index)
		m.bubble_up(parent)
	}
}

func (m *minH) swap(a, b int) {
	m.h[a], m.h[b] = m.h[b], m.h[a]
}

func (m *minH) remove() *ListNode {
	minNode := m.h[0]
	m.swap(0, m.hSize-1)
	m.h = m.h[:m.hSize-1]
	m.hSize -= 1
	m.bubble_down(0)
	return minNode
}

func (m *minH) bubble_down(index int) {
	if index >= m.hSize {
		return
	}
	smallest := index
	left := index*2 + 1
	right := index*2 + 2

	if left < m.hSize && m.h[left].Val < m.h[smallest].Val {
		smallest = left
	}
	if right < m.hSize && m.h[right].Val < m.h[smallest].Val {
		smallest = right
	}

	if smallest != index {
		m.swap(smallest, index)
		m.bubble_down(smallest)
	}
}
