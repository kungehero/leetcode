/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */
func isHappy(n int) bool {
	a, b := n, n
	for {
		a = sum(a)
		b = sum(sum(b))
		if a == 1 || b == 1 {
			return true
		}
		if a == b {
			return false
		}
	}
}

func sum(n int) int {
	var arr []int
	for n != 0 {
		arr = append(arr, n%10)
		n = n / 10
	}
	var sum int
	for _, v := range arr {
		sum += v * v
	}
	return sum
}
