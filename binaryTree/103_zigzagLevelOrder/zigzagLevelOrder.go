package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	level := 1
	queue := []*TreeNode{root}
	for queue != nil {
		var q []*TreeNode
		var valList []int
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			valList = append(valList, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if level%2 == 0 {
			for i, j := 0, len(valList)-1; i < j; i, j = i+1, j-1 {
				valList[i], valList[j] = valList[j], valList[i]
			}
		}
		queue = q
		ans = append(ans, valList)
		level++
	}
	return ans
}
