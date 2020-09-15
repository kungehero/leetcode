/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	// 判断被除数与除数的正负性

	count, flag := 0, 1

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		flag = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}
	if divisor == 1 {
		if dividend >= math.MaxInt32 {
			if flag > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}
	for dividend >= divisor {
		count++
		dividend -= divisor
	}

	if count >= math.MaxInt32 {
		if flag > 0 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	return count * flag
}

// @lc code=end

