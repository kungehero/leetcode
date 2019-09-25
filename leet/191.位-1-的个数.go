/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */
func hammingWeight(num uint32) int {
	val := 0
	for num > 0 {
		val += int(num & 1)
		num = num >> 1
	}
	return val
}
