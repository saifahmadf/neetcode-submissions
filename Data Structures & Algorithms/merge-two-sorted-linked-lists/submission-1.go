/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current := &ListNode{}
	ll := current

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			ll.Next = &ListNode { Val: list2.Val}
			list2 = list2.Next
			
		} else {
			ll.Next = &ListNode { Val: list1.Val}
			list1 = list1.Next
		}
		ll = ll.Next
	}
	for list1 != nil {
		ll.Next = &ListNode { Val: list1.Val}
		list1 = list1.Next
		ll = ll.Next
		
	}
	for list2 != nil {
		ll.Next = &ListNode { Val: list2.Val}
		list2 = list2.Next		
		ll = ll.Next
	}
	return current.Next
}
