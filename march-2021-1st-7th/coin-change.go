coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 0; i <= amount; i++ {
		for _, c := range coins {
			if (i - c) >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-c]+1)))
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
