/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 0, n
	for l <= r {
		mid := l + (r-l)/2
		if guess(mid) == 0 {
			return mid
		}

		if guess(mid) < 0 {
			r = mid - 1
		}
		if guess(mid) > 0 {
			l = mid + 1
		}
	}
	return -1
}

// @lc code=end

