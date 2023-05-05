package _57_binaryTreePaths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		if path == "" {
			path += strconv.Itoa(root.Val)
		} else {
			path += "->" + strconv.Itoa(root.Val)
		}
		if root.Left == nil && root.Right == nil {
			ans = append(ans, path)
			return
		}
		dfs(root.Left, path)
		dfs(root.Right, path)
	}
	dfs(root, "")
	return ans
}
