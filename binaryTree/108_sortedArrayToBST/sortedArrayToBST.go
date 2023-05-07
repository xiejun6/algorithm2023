package _08_sortedArrayToBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	index := len(nums) / 2
	root := &TreeNode{Val: nums[index]}
	root.Left = sortedArrayToBST(nums[:index])
	root.Right = sortedArrayToBST(nums[index+1:])
	return root
}
