/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {

	num := 0
	for i := 1; i <= 2<<32; i++ {
		if num == n {
			return i - 1
		}
		if checkNthUglyNumber(i) {
			num++
		}
	}
	return 0
}

func checkNthUglyNumber(n int) bool {
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

// @lc code=end

