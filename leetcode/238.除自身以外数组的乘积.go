/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {

	/* 	length := len(nums)

	   	// L 和 R 分别表示左右两侧的乘积列表
	   	L, R, answer := make([]int, length), make([]int, length), make([]int, length)

	   	// L[i] 为索引 i 左侧所有元素的乘积
	   	// 对于索引为 '0' 的元素，因为左侧没有元素，所以 L[0] = 1
	   	L[0] = 1
	   	for i := 1; i < length; i++ {
	   		L[i] = nums[i-1] * L[i-1]
	   	}

	   	// R[i] 为索引 i 右侧所有元素的乘积
	   	// 对于索引为 'length-1' 的元素，因为右侧没有元素，所以 R[length-1] = 1
	   	R[length-1] = 1
	   	for i := length - 2; i >= 0; i-- {
	   		R[i] = nums[i+1] * R[i+1]
	   	}

	   	// 对于索引 i，除 nums[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	   	for i := 0; i < length; i++ {
	   		answer[i] = L[i] * R[i]
	   	}
	   	return answer */

	var ans []int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res := 1
			for j := 1; j < len(nums); j++ {
				res *= nums[j]
			}
			ans = append(ans, res)
		}
		if i == len(nums)-1 {
			res := 1
			for j := i - 1; j >= 0; j-- {
				res *= nums[j]
			}
			ans = append(ans, res)

		}

		if i > 0 && i < len(nums)-1 {
			res := 1
			for j := 0; j < i; j++ {
				res *= nums[j]
			}
			for k := i + 1; k <= len(nums)-1; k++ {
				res *= nums[k]
			}
			ans = append(ans, res)
		}

	}
	return ans
}

// @lc code=end

