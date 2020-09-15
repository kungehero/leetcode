/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	i:=0
	for i < len(nums) {
		if nums[i]<target{
			i++
		}else{
			break
		}
	} 
	return i
}
// @lc code=end

