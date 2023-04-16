package linklist

import "testing"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	fast, slow := head, head
	for i := 1; i <= n; i++ {
		fast = fast.Next
	}
	prev := dummyHead
	for fast != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
	}
	prev.Next = slow.Next
	return dummyHead.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	head := MakeList(data)
	newHead := removeNthFromEnd(head, 2)
	PrintList(newHead)

}
