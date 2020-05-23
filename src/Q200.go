package src

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	result := 0
	for rk := range grid {
		for ck := range grid[rk] {
			if grid[rk][ck] == '1' {
				result++
				numIslandsDFS(grid, rk, ck)
			}
		}
	}
	return result
}

func numIslandsDFS(grid [][]byte, rk, ck int) {
	if rk < 0 || ck < 0 || rk >= len(grid) || ck >= len(grid[0]) || grid[rk][ck] == '0' {
		return
	}
	grid[rk][ck] = '0'
	numIslandsDFS(grid, rk-1, ck)
	numIslandsDFS(grid, rk+1, ck)
	numIslandsDFS(grid, rk, ck-1)
	numIslandsDFS(grid, rk, ck+1)
}
