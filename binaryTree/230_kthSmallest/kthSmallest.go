package _30_kthSmallest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	ans := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		count++
		if count == k {
			ans = root.Val
			return
		}
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
