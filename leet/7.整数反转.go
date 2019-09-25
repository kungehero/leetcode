package leet

import (
	"math"
	"strconv"
)

/*Your runtime beats 12.89 % of golang submissions
 √ Your memory usage beats 68.95 % of golang submissions (2.2 MB)
* @lc app=leetcode.cn id=7 lang=golang
*
* [7] 整数反转
*/
func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = x * -1
	}
	var result int64 = 0
	str := strconv.Itoa(x)
	for i := len(str) - 1; i >= 0; i-- {
		r := int64(x % 10)
		for j := 0; j < i; j++ {
			r = r * 10
		}
		result += r
		x = x / 10
	}
	if result > math.MaxInt32 || result < -math.MaxInt32-1 {
		return 0
	}
	return int(result) * flag
}
