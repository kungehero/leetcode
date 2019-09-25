/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			target := prices[j] - prices[i]
			if target > res {
				res = target
			}
		}
	}
	return res
}
