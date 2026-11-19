/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	tokens := []string{}
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			tokens = append(tokens, "#")
			return
		}
		tokens = append(tokens, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(tokens, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	i := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if i >= len(tokens) || tokens[i] == "#" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(tokens[i])
		i++
		node := &TreeNode{Val: val}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}
