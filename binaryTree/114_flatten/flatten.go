package _14_flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var nodeList []*TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			nodeList = append(nodeList, root)
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)
	for i := 0; i < len(nodeList)-1; i++ {
		nodeList[i].Left = nil
		nodeList[i].Right = nodeList[i+1]
	}
}
