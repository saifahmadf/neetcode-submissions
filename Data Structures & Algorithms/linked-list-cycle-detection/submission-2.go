/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    x := head
	if head == nil || head.Next == nil {
		return false
	}
	y := head.Next

	for x != nil && y != nil {
		if x == y {
			return true
		} else if x == nil || y == nil || x.Next == nil || y.Next == nil || y.Next.Next == nil{
			return false
		}
		x = x.Next
		y = y.Next.Next
	}
	return false
}
