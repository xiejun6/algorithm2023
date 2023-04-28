package _13_pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var seqList []int
	var dfs func(root *TreeNode, targetSum int)
	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		seqList = append(seqList, root.Val)
		defer func() {
			seqList = seqList[:len(seqList)-1]
		}()
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			ans = append(ans, append([]int{}, seqList...))
			return
		}

		dfs(root.Left, targetSum-root.Val)
		dfs(root.Right, targetSum-root.Val)
	}
	dfs(root, targetSum)
	return ans
}
