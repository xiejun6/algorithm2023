package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := []*TreeNode{root}
	for queue != nil {
		n := len(queue)
		var p []*TreeNode
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
			if i == n-1 {
				ans = append(ans, node.Val)
			}
		}
		queue = p
	}
	return ans
}
