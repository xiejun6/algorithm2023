package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(deep(root.Left)-deep(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(deep(root.Left)+1, deep(root.Right)+1)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
