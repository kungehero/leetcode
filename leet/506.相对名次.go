import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */
func findRelativeRanks(nums []int) []string {
	var arr []int
	arr = append(arr, nums...)
	sort.Ints(nums)
	Reverse(nums)
	var res []string
	m := make(map[int]string)
	count := 0
	for _, v := range nums {
		count++
		if count > 3 {
			m[v] = strconv.Itoa(count)
		} else if count == 1 {
			m[v] = "Gold Medal"
		} else if count == 2 {
			m[v] = "Silver Medal"
		} else if count == 3 {
			m[v] = "Bronze Medal"
		}
	}
	for _, v := range arr {
		res = append(res, m[v])
	}
	return res
}

func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; j > i; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
