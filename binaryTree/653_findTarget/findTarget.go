package _53_findTarget

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var data []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		data = append(data, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	i, j := 0, len(data)-1
	for i < j {
		sum := data[i] + data[j]
		if sum == k {
			return true
		} else if sum < k {
			i++
		} else {
			j--
		}
	}
	return false
}
