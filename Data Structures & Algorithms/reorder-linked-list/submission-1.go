/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return 
	}

	slow, fast := head,head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil

	first := head
	second = revList(second)

	finalList := &ListNode{}
	cur := finalList 

	for second != nil {
		cur.Next = first
		cur = cur.Next
		first = first.Next
		cur.Next = second
		cur = cur.Next
		second = second.Next
	}
	
	for first != nil {
		cur.Next = first
		cur = cur.Next
		first = first.Next
	}
	head = finalList
}


func revList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}