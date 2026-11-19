func countComponents(n int, edges [][]int) int {
	adj := make([][]int, n)
	visit := make([]bool, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	fmt.Println("adj: ", adj)
	
	var dfs func(int)
	dfs = func(node int) {
		for _, nei := range adj[node] {
			if !visit[nei] {
				visit[nei] = true
				dfs(nei)
			}
		}
	}

	result := 0
	for node := 0; node < n; node++{
		if !visit[node] {
			visit[node] = true
			dfs(node)
			result++
		}
	}
	return result

}
