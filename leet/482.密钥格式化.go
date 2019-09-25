import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */
func licenseKeyFormatting(s string, K int) string {
	count := 0
	var res string
	for i := len(s) - 1; i > -1; i-- {
		if s[i] != '-' {
			count++
			res += string(s[i])
			if count == K {
				res += "-"
				count = 0
			}
		}
	}
	if len(res) == 0 {
		return ""
	}
	l := Reverse(res)
	if l[0] == '-' {
		l = l[1:]
	}
	return strings.ToUpper(l)
}

func Reverse(s string) string {
	a := []byte(s)
	for i, j := 0, len(a)-1; j > i; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
