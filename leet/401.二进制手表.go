import "strconv"

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */
func readBinaryWatch(num int) []string {
	var res []string

	if num == 0 {
		res = append(res, "0:00")
		return res
	}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if count1(i)+count1(j) == num {
				res = append(res, toString(i)+":"+sanExp(j < 10, "0"+toString(j), toString(j)))
			}
		}
	}
	return res
}

func sanExp(b bool, x, y string) string {
	if b {
		return x
	}
	return y
}

func toString(num int) string {
	return strconv.Itoa(num)
}

func count1(n int) int {
	res := 0
	for n != 0 {
		n = n & (n - 1)
		res++
	}
	return res
}
