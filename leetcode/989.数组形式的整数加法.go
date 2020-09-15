/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 */

// @lc code=start
func addToArrayForm(A []int, K int) []int {
	temp := len(A)
	res := []int{}

	for temp > 0 || K > 0 {
		temp--
		if temp >= 0 {
			K += A[temp]
		}
		res = append(res, K%10)
		K /= 10
	}
	reverse(res)
	return res
}
func reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

// @lc code=end

