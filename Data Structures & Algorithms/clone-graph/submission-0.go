/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    
	oldToNewGraph := make(map[*Node]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, found := oldToNewGraph[node]; found {
			return oldToNewGraph[node]
		}

		copy := &Node{Val: node.Val}
		oldToNewGraph[node] = copy

		for _, nei := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors, dfs(nei))
		}

		return copy
	}

	return dfs(node)
}
