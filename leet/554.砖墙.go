import "sort"

/*
 * @lc app=leetcode.cn id=554 lang=golang
 *
 * [554] 砖墙
 */
func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j < len(wall[i]); j++ {
			sum += wall[i][j]
			if _, ok := m[sum]; ok {
				m[sum]++
				continue
			}
			m[sum] = 1
		}
	}
	var arr []int
	for _, v := range m {
		arr = append(arr, v)
	}
	sort.Ints(arr)
	if len(arr) < 2 {
		return len(wall)
	}
	//fmt.Println(len(wall) - arr[len(arr)-2])
	return len(wall) - arr[len(arr)-2]
}
