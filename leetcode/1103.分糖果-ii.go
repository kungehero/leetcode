/*
 * @lc app=leetcode.cn id=1103 lang=golang
 *
 * [1103] 分糖果 II
 */

// @lc code=start
func distributeCandies(candies int, num_people int) []int {
	var res []int
	for i := 1; i <= 4; i++ {
		res = append(res, i)
	}
	return res
}

// @lc code=end

