package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 返回{偷当前节点值，不偷当前节点值}
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	robCur := root.Val + left[1] + right[1]
	noRobCur := max(left[0], left[1]) + max(right[0], right[1])
	return []int{robCur, noRobCur}
}

func main() {

}
