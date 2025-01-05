// 380. O(1) 时间插入、删除和获取随机元素

package algorithm_300

import "math/rand"

type RandomizedSet struct {
	NumMap map[int]int
	Array  []int
}

func ConstructorV2() RandomizedSet {
	return RandomizedSet{
		NumMap: make(map[int]int),
		Array:  make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.NumMap[val]; ok {
		return false
	}

	this.Array = append(this.Array, val)
	this.NumMap[val] = len(this.Array) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.NumMap[val]
	if !ok {
		return false
	}

	delete(this.NumMap, val)
	if idx != len(this.Array)-1 {
		this.Array[idx] = this.Array[len(this.Array)-1]
		this.NumMap[this.Array[idx]] = idx
	}
	this.Array = this.Array[:len(this.Array)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.Array))
	return this.Array[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
