package linklist

import "testing"

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func Test_removeElements(t *testing.T) {
	data := []int{1, 2, 6, 3, 4, 5, 6}
	head := MakeList(data)
	newHead := removeElements(head, 6)
	PrintList(newHead)
}
