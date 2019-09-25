/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */
func countPrimes(n int) int {
	count := 0
	a := make([]bool, n)
	for i := 2; i < n; i++ {
		if !a[i] {
			count++
			for j := i + i; j < n; j += i {
				a[j] = true
			}
		}
	}
	return count
}
