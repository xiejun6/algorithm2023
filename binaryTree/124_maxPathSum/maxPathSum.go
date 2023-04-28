package maxpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftVal := max(dfs(root.Left), 0)
		rightVal := max(dfs(root.Right), 0)
		val := root.Val + leftVal + rightVal
		if val > ans {
			ans = val
		}
		return root.Val + max(leftVal, rightVal)
	}
	dfs(root)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
