/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	a := convertInt(num1)
	b := convertInt(num2)
	return convertString(a * b)

}
func convertInt(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
	}
	return num
}
func convertString(i int) string {
	str := ""
	arr := []int{}
	if i < 10 {
		return string(i - '0')
	}
	for i != 0 {
		if i == 0 {
			arr = append(arr, i)
			break
		}
		i %= 10
		arr = append(arr, i)
	}
	k := len(arr) - 1
	for k > 0 {
		str += string(arr[i] + '0')
		k--
	}
	return str
}

// @lc code=end

