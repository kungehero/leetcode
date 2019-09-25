/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */
func toHex(num int) string {
	res := ""
	m := map[int]string{
		10: "a",
		11: "b",
		12: "13",
		14: "14",
		15: "15",
	}
	for num%16 != 0 {
		sum := num / 16
		if sum < 10 {
			sum += sum
		} else {
			if _, ok := m[sum]; ok {
				sum += m[sum]
			}
		}
	}
	
}
