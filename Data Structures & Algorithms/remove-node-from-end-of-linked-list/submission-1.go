/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    r := lenList(head) - n + 1
	if r == 1 {
		return head.Next
	}
	cur := head
	for count := 1 ; count < r-1; count++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	
	return head
}

func lenList(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
