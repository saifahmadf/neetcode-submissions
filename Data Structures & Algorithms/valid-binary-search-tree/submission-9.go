/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	type entry struct {
		node *TreeNode
		min int
		max int
	}

	queue := list.New()
	queue.PushBack(entry{root, math.MinInt, math.MaxInt})

	for queue.Len() > 0 {
		e := queue.Front().Value.(entry)
		queue.Remove(queue.Front())
		if e.node.Val <= e.min || e.node.Val >= e.max {
			return false
		}
		if e.node.Left != nil {
			queue.PushBack(entry{e.node.Left, e.min, e.node.Val})
		}
		if e.node.Right != nil {
			queue.PushBack(entry{e.node.Right, e.node.Val, e.max})
		}
	}
	return true
}
