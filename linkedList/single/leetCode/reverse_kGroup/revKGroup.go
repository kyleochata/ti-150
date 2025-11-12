package reversekgroup

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prevGroup := dummy

	for {
		kth := getK(prevGroup, k)
		if kth == nil {
			break
		}

		nextGroup := kth.Next
		// Reverse the current group
		prev, curr := kth.Next, prevGroup.Next
		firstNode := prevGroup.Next // Save the first node of current group

		for curr != nextGroup {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// Connect previous group to the new start of reversed group
		prevGroup.Next = kth
		// Move to next group
		prevGroup = firstNode
	}

	return dummy.Next
}

func getK(node *ListNode, k int) *ListNode {
	for node != nil && k > 0 {
		node = node.Next
		k--
	}
	return node
}
