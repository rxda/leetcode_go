package src

import "math/rand"

type RandomizedSet struct {
	Map  map[int]int
	List []int
}

/** Initialize your data structure here. */
func Constructorb() RandomizedSet {
	return RandomizedSet{
		make(map[int]int),
		make([]int, 0, 8),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Map[val]; ok {
		return false
	}
	this.Map[val] = len(this.List)
	this.List = append(this.List, val)
	return true

}

///** Removes a value from the set. Returns true if the set contained the specified element. */
//func (this *RandomizedSet) Remove(val int) bool {
//	if mapV, ok := this.Map[val]; !ok {
//		return false
//	} else {
//		if mapV < len(this.List)-1 {
//			lastOne := this.List[len(this.List)-1]
//			this.List[mapV] =
//		}
//	}
//
//}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.List[rand.Intn(len(this.List))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
