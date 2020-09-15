/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
	Data []int
	mp   map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{Data: []int{}, mp: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mp[val]; ok {
		return false
	} else {
		this.Data = append(this.Data, val)
		this.mp[val] = len(this.Data) - 1
	}
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if k, ok := this.mp[val]; ok {
		this.mp[this.Data[len(this.Data)-1]] = k
		this.Data[k], this.Data[len(this.Data)-1] = this.Data[len(this.Data)-1], this.Data[k]
		this.Data = this.Data[:(len(this.Data) - 1)]
		delete(this.mp, val)
	} else {
		return false
	}
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.Data[rand.Intn(len(this.Data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

