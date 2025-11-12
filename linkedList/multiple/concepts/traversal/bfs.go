package traversal

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFS(root *TreeNode, q []*TreeNode) {
	recurseBFS([]*TreeNode{root})
}
func recurseBFS(q []*TreeNode) {
	if len(q) == 0 {
		return
	}

	current := q[0]
	q = q[1:]
	fmt.Printf("%v --> ", current.Val)
	if current.Left != nil {
		q = append(q, current.Left)
	}
	if current.Right != nil {
		q = append(q, current.Right)
	}
	recurseBFS(q)
}
