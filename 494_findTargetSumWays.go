import "fmt"

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	bias := 1000
	dp := make(map[int]int, 2*bias)
	dp[bias+nums[0]] = 1
	dp[bias-nums[0]] = 1 + dp[bias]
	for i := 1; i < len(nums); i++ {
		tdp := make(map[int]int, 0)
		for k, _ := range dp {
			n0 := k - nums[i]
			n1 := k + nums[i]
			tdp[n0] += dp[k]
			tdp[n1] += dp[k]
		}
		//fmt.Println("dp:", dp)
		dp = tdp
	}
	//fmt.Println("dp:", dp)
	return dp[bias+S]
}
