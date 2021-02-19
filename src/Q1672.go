package src

func maximumWealth(accounts [][]int) int {
	m := len(accounts)    // row
	n := len(accounts[0]) // col
	max := 0
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			accounts[i][0] += accounts[i][j]
		}
		if accounts[i][0]> max{
			max = accounts[i][0]
		}
	}
	return max
}
