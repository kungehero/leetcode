/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package leetcode

// @lc code=start
func moveZeroes(nums []int) {

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if k != i {
				Swap(nums, k, i)
				k++
			} else {
				k++
			}
		}
	}
}

func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// @lc code=end
