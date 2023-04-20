package linklist

import (
	"testing"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//fast一定要等于head.next, 否则只有两个元素的时候会出问题
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	p := slow.Next
	slow.Next = nil
	//fmt.Println("head:")
	//PrintList(head)
	//fmt.Println("fast:")
	//PrintList(p)
	return mergeList(sortList(head), sortList(p))
}

func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead
	p, q := head1, head2
	for p != nil && q != nil {
		if p.Val < q.Val {
			prev.Next = p
			prev = p
			p = p.Next
		} else {
			prev.Next = q
			prev = q
			q = q.Next
		}
	}
	if p != nil {
		prev.Next = p
	} else if q != nil {
		prev.Next = q
	}
	return dummyHead.Next
}

func Test_sortList(t *testing.T) {
	head := MakeList([]int{4, 2, 1, 3})
	PrintList(head)
	head1 := sortList(head)
	PrintList(head1)
	h := mergeList(MakeList([]int{1, 2, 5}), MakeList([]int{2, 3, 4}))
	PrintList(h)
}
