package traversal

import "fmt"

// preOrderDFS traverses to the left most child after solving the root.
//
//	root > left subtree > right subtree
func PreOrderDFS(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%v --> ", root.Val)
	PreOrderDFS(root.Left)
	PreOrderDFS(root.Right)
}

// inOrderDFS traverses to the lowest left leaf then solves up
//
//	left subtree > root > right subtree
func InOrderDFS(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderDFS(root.Left)
	fmt.Printf("%v --> ", root.Val)
	InOrderDFS(root.Right)
}

// postOrderDFS traverses the tree in a postOrder method:
//
//	left subTree > right subTree > root
func PostOrderDFS(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderDFS(root.Left)
	PostOrderDFS(root.Right)
	fmt.Printf("%v --> ", root.Val)
}
