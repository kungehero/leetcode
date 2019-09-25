import "math"

/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */
func isUgly(num int) bool {
	if num == 1 {
		return true
	}
	if num == 0 {
		return false
	}

	if num >= math.MaxInt32 || num <= math.MinInt32 {
		return false
	}

	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}
	return num == 1
}
