package _04_sumOfLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	v := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		v = root.Left.Val
	}
	return v + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
