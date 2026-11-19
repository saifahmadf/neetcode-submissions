/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	// if root == nil {
	// 	return 0
	// }
	result := []int{root.Val}
	dfs(root, result)
	return result[0]
}

func dfs(root *TreeNode, result []int) int {
	if root == nil {
		return 0
	}
	leftMax := dfs(root.Left, result)
	rightMax := dfs(root.Right, result)
	leftMax = max(leftMax, 0)
	rightMax = max(rightMax, 0)
	result[0] = max(result[0], root.Val + leftMax + rightMax)
	r := root.Val + max(leftMax, rightMax)
	return r 
}


