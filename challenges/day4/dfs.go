package day4

var directions = [8][2]int{
	{0, 1},   // right
	{1, 0},   // down
	{1, 1},   // down-right
	{1, -1},  // down-left
	{0, -1},  // left
	{-1, 0},  // up
	{-1, -1}, // up-left
	{-1, 1},  // up-right
}

func dfs(grid [][]rune, node *TrieNode, x, y int, visited [][]bool, path string, result *[]string, direction [2]int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || visited[x][y] {
		return
	}

	char := grid[x][y]
	if _, exists := node.children[char]; !exists {
		return
	}

	visited[x][y] = true
	path += string(char)
	node = node.children[char]

	if node.isEnd {
		*result = append(*result, path)
	}

	if direction == [2]int{0, 0} {
		for _, dir := range directions {
			dfs(grid, node, x+dir[0], y+dir[1], visited, path, result, dir)
		}
	} else {
		dfs(grid, node, x+direction[0], y+direction[1], visited, path, result, direction)
	}

	visited[x][y] = false
}
