import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */
func addStrings(num1 string, num2 string) string {
	var res string
	count := 0
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 {
		a1, a2 := 0, 0
		if i >= 0 {
			a1, _ = strconv.Atoi(string((num1[i])))
		} else {
			a1 = 0
		}
		if j >= 0 {
			a2, _ = strconv.Atoi(string((num2[j])))

		} else {
			a2 = 0
		}
		fmt.Println(a1, a2)
		temp := a1 + a2 + count
		count = temp / 10
		res = strconv.Itoa(temp%10) + res
		i--
		j--
	}
	//res = revers(res)
	if count > 0 {
		return "1" + res
	}
	return res
}

/* func revers(s string) string {
	a := []byte(s)
	for i, j := 0, len(a)-1; j > i; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
} */
