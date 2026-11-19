/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val : preorder[0]}

	idx := 0

	for k, v := range inorder {
		if v == preorder[0] {
			idx = k
			break
		}
	}

	leftSize := idx

	root.Left = buildTree(preorder[1:1+leftSize], inorder[:idx])
	root.Right = buildTree(preorder[1+leftSize:], inorder[idx+1:])

	return root
}
