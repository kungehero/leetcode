import "math"

/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 */
func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	cur := int(math.Sqrt(float64(num)))
	sum := 0
	for i := 1; i <= cur; i++ {
		if i == 1 {
			sum += 1
		} else {
			if num%i == 0 {
				sum += i
				sum += num / i
			}
		}
	}
	if sum == num {
		return true
	}
	return false
}
