/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
    for len(lists) > 1 {
		next := make([]*ListNode, 0, (len(lists)+1)/2)
		for i := 0; i < len(lists); i = i+2 {
			if i+1 < len(lists) {
				next = append(next, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				next = append(next, lists[i])
			}
		}
		lists = next
	}
	return lists[0]
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	l := &ListNode{}
	cur := l
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return l.Next
}