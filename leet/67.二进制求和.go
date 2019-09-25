/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
import (
	"bytes"
	"log"
	"strconv"
)

func addBinary(a string, b string) string {
	al := len(a) - 1
	bl := len(b) - 1
	carry := 0
	var str bytes.Buffer
	for al >= 0 || bl >= 0 {
		sum := carry
		if al >= 0 {
			sum += getInt(al, a)
		}
		if bl >= 0 {
			sum += getInt(bl, b)
		}
		s := sum % 2
		carry = sum / 2
		str.WriteString(strconv.Itoa(s))
		al--
		bl--
	}
	if carry != 0 {
		str.WriteString("1")
	}
	return reverseString(str.String())
}

func getInt(i int, str string) int {
	v, err := strconv.Atoi(string(str[i]))
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
