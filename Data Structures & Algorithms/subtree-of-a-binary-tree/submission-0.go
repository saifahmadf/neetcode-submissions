/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if (root != nil && subRoot == nil) || (root == nil && subRoot != nil) {
		return true
	}

	queue1 := list.New()
	queue1.PushBack(root)

	for queue1.Len() > 0 {
		node1 := queue1.Front().Value.(*TreeNode)
		queue1.Remove(queue1.Front())
		if isSameTree(node1, subRoot) {
			return true
		}
		if node1.Left != nil {
			queue1.PushBack(node1.Left)
		}
		if node1.Right != nil {
			queue1.PushBack(node1.Right)
		}
	}

    return false
}

func isSameTree(root, subRoot *TreeNode) bool {
	queue1 := list.New()
	queue2 := list.New()
	queue1.PushBack(root)
	queue2.PushBack(subRoot)

	for queue1.Len() > 0 {
		node1 := queue1.Front().Value.(*TreeNode)
		node2 := queue2.Front().Value.(*TreeNode)
		queue1.Remove(queue1.Front())
		queue2.Remove(queue2.Front())

		if node1.Val != node2.Val {
			return false
		}
		if (node1.Left != nil && node2.Left == nil) || (node1.Left == nil && node2.Left != nil) || (node1.Right != nil && node2.Right == nil) || (node1.Right == nil && node2.Right != nil) {
			return false
		}

		if node1.Left != nil {
			queue1.PushBack(node1.Left)
			queue2.PushBack(node2.Left)
		}
		if node1.Right != nil {
			queue1.PushBack(node1.Right)
			queue2.PushBack(node2.Right)
		}
	}
	return true
}
