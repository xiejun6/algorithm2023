package _5_generateTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}

	var dfs func(start, end int) []*TreeNode
	dfs = func(start, end int) []*TreeNode {
		if start > end {
			//note:这个地方不能返回 []*TreeNode{}
			return []*TreeNode{nil}
		}
		var ans []*TreeNode
		for i := start; i <= end; i++ {
			left := dfs(start, i-1)
			right := dfs(i+1, end)
			for _, n1 := range left {
				for _, n2 := range right {
					root := &TreeNode{Val: i, Left: n1, Right: n2}
					ans = append(ans, root)
				}
			}
		}
		return ans
	}
	return dfs(1, n)
}
