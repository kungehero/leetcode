import "strconv"

/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */
func convertToBase7(num int) string {
	n := 0
	if num < 0 {
		n = num * -1
	} else {
		n = num
	}
	var res string
	// funny
	for n/7 != 0 {
		a := n % 7
		res = strconv.Itoa(a) + res
		n /= 7
	}
	res = strconv.Itoa(n) + res
	if num < 0 {
		res = "-" + res
	}
	return res
}
