package linklist

import "testing"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p, q := list1, list2
	dummyHead := &ListNode{}
	prev := dummyHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			prev.Next = p
			p = p.Next
		} else {
			prev.Next = q
			q = q.Next
		}
		prev = prev.Next
	}
	if p != nil {
		prev.Next = p
	} else if q != nil {
		prev.Next = q
	}
	return dummyHead.Next
}

func Test_mergeTwoLists(t *testing.T) {
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}
	h1 := MakeList(l1)
	h2 := MakeList(l2)
	h3 := mergeTwoLists(h1, h2)
	PrintList(h3)

	l1 = []int{}
	l2 = []int{}
	h1 = MakeList(l1)
	h2 = MakeList(l2)
	h3 = mergeTwoLists(h1, h2)
	PrintList(h3)

	l1 = []int{}
	l2 = []int{0}
	h1 = MakeList(l1)
	h2 = MakeList(l2)
	h3 = mergeTwoLists(h1, h2)
	PrintList(h3)
}
