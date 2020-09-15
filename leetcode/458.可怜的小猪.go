/*
 * @lc app=leetcode.cn id=458 lang=golang
 *
 * [458] 可怜的小猪
 */

// @lc code=start
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	states := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(states))))
}

// @lc code=end

