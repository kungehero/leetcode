/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	//sort.Ints(nums)
	pre, cur, rear := 0, 0, len(nums)-1

	for cur <= rear {
		if nums[cur] == 0 {
			nums[pre], nums[cur] = nums[cur], nums[pre]
			cur++
			pre++
		} else if nums[cur] == 1 {
			cur++
		} else {
			nums[cur], nums[rear] = nums[rear], nums[cur]
			rear--
		}
	}
}

// @lc code=end

