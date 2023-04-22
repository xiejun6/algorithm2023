package linklist

import "testing"

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//if fast != nil {
	//	slow = slow.Next
	//}
	return slow
}

func Test_middleNode(t *testing.T) {
	head := MakeList([]int{1, 2, 3, 4, 5})
	h2 := middleNode(head)
	PrintList(h2)
}
