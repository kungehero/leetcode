import (
	"log"
	"regexp"
	"strings"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */
func isPalindrome(s string) bool {
	var b []string
	for _, v := range s {
		if !Math(string(v)) {
			b = append(b, strings.ToLower(string(v)))
		}
	}

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}
	}
	return true
}

func Math(v string) bool {
	flag, err := regexp.MatchString("[^A-Za-z0-9]", v)
	if err != nil {
		log.Fatal(err)
	}
	return flag
}
