/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	res := math.MinInt32
	for i := len(prices) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			res = max(res, prices[i]-prices[j])
		}
	}

	if res < 0 {
		return 0
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

