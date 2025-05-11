package leetcode75

func findCircleNum(isConnected [][]int) int {
	n, res := len(isConnected[0]), 0
	if n == 0 {
		return res
	}
	seen := make([]bool, n)

	for i := 0; i < n; i++ {
		if !seen[i] {
			res++
			dfs_547(isConnected, seen, i)
		}
	}
	return res
}

func dfs_547(isConnected [][]int, seen []bool, currCity int) {
	seen[currCity] = true

	for j := 0; j < len(isConnected); j++ {
		if isConnected[currCity][j] == 1 && !seen[j] {
			dfs_547(isConnected, seen, j)
		}
	}
}
