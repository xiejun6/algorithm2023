package linklist

import (
	"testing"
)

func deleteDuplicates_82(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			v := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == v {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func Test_deleteDuplicates(t *testing.T) {
	data := []int{1, 2, 3, 3, 4, 4, 5}
	head := MakeList(data)
	PrintList(deleteDuplicates_82(head))

	data = []int{1, 1, 1, 2, 3}
	head = MakeList(data)
	PrintList(deleteDuplicates_82(head))
}
