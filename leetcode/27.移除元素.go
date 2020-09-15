/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */
package leetcode

// @lc code=start
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if k != i {
				swap(nums, k, i)
				k++

			} else {
				k++
			}
		}
	}
	return len(nums[:k])
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// @lc code=end
