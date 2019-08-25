package src

func numRollsToTarget(d int, f int, target int) int {
	var m int = 10e9 + 7
	dp := make([][]int,30)
	dp[0] = make([]int, 1000)
	dp[0][0] =1
	for i:= 1;i<=d;i++{
		dp[i] = make([]int, 1000)
		for now :=1 ;now <= f;now ++{
			for prev:=0;prev+now <= target;prev++{
				dp[i][prev+now] += dp[i-1][prev] %m
			}
		}
	}
	return dp[d][target]
}
