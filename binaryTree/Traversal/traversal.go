package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
144. 二叉树的前序遍历
94. 二叉树的中序遍历
145. 二叉树的后序遍历
*/
//前序
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			ans = append(ans, root.Val)
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)
	return ans
}

// 中序
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			ans = append(ans, root.Val)
			dfs(root.Right)
		}
	}
	dfs(root)
	return ans
}

// 后序
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			dfs(root.Right)
			ans = append(ans, root.Val)
		}
	}
	dfs(root)
	return ans
}
func main() {

}
