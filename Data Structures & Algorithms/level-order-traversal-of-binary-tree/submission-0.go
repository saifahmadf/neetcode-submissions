/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    var result [][]int
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		level := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			level = append(level, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, level)
		
	}
	return result
}
