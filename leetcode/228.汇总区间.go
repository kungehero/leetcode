/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	var res []string
	l, r := 0, 0
	for r < len(nums) {
		if r+1 < len(nums) && nums[r+1]-nums[r] == 1 {
			r++
			continue
		}
		if l == r {
			res = append(res, strconv.Itoa(nums[r]))
		} else {
			res = append(res, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r]))
		}
		l = r + 1
		r++
	}
	return res
}

// @lc code=end

