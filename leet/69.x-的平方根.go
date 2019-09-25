import (
	"log"
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
func mySqrt(x int) int {
	v := math.Sqrt(float64(x))
	r := strconv.FormatFloat(v, 'f', -1, 64)
	var s []rune
	for _, value := range r {
		if value != '.' {
			s = append(s, value)
		} else {
			break
		}
	}
	i, err := strconv.Atoi(string(s))
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func mySqrt1(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}
