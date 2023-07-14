package linklist

import "testing"

func partition(head *ListNode, x int) *ListNode {
	dummyHead1, dummyHead2 := &ListNode{}, &ListNode{}
	prev1, prev2 := dummyHead1, dummyHead2
	p := head
	for p != nil {
		q := p.Next
		if p.Val < x {
			prev1.Next = p
			prev1 = p
		} else {
			prev2.Next = p
			prev2 = p
		}
		p = q
	}
	prev2.Next = nil
	prev1.Next = dummyHead2.Next
	return dummyHead1.Next
}

func Test_Partition(t *testing.T) {
	head := MakeList([]int{1, 4, 3, 2, 5, 2})
	PrintList(head)
	h2 := partition(head, 3)
	PrintList(h2)

	head = MakeList([]int{2, 1})
	PrintList(head)
	h2 = partition(head, 2)
	PrintList(h2)
}
