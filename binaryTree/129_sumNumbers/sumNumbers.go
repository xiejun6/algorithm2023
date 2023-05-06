package _29_sumNumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(root *TreeNode, val int)
	dfs = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		val = val*10 + root.Val
		if root.Left == nil && root.Right == nil {
			sum += val
			return
		}
		dfs(root.Left, val)
		dfs(root.Right, val)
	}
	dfs(root, 0)
	return sum
}
