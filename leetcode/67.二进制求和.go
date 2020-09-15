/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
// @lc code=start
func addBinary(a string, b string) string {
	res := ""
	carry := 0
	length := max(len(a), len(b))
	for i := 0; i < length; i++ {
		if i < len(a) {
			carry += int(a[len(a)-i-1] - '0')
		}
		if i < len(b) {
			carry += int(b[len(b)-i-1] - '0')
		}
		res = strconv.Itoa(carry%2) + res
		carry /= 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

