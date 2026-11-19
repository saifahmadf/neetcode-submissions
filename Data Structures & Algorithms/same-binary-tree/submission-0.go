/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	queue1 := list.New()
	queue2 := list.New()
	queue1.PushBack(p)
	queue2.PushBack(q)

	for queue1.Len() > 0 && queue2.Len() > 0 {
		node1 := queue1.Front().Value.(*TreeNode)
		node2 := queue2.Front().Value.(*TreeNode)
		queue1.Remove(queue1.Front())
		queue2.Remove(queue2.Front())
		if (node1.Val != node2.Val) || queue1.Len() != queue2.Len() {
			return false
		}
		if node1.Left != nil && node2.Left != nil {
			queue1.PushBack(node1.Left)
			queue2.PushBack(node2.Left)
		} 
		if node1.Right != nil && node2.Right != nil{
			queue1.PushBack(node1.Right)
			queue2.PushBack(node2.Right)
		}
		if (node1.Left != nil && node2.Left == nil) || (node1.Left == nil && node2.Left != nil) || (node1.Right != nil && node2.Right == nil) || (node1.Right == nil && node2.Right != nil) {
			return false
		}
	}
	return true
}
