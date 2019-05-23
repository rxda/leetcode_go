package src

type MyHashSet map[int]bool


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return make(MyHashSet)
}


func (this *MyHashSet) Add(key int)  {
	(*this)[key] =true
}


func (this *MyHashSet) Remove(key int)  {
	if _,ok:=(*this)[key];ok {
		delete(*this,key)
	}
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if _,ok:=(*this)[key];ok {
		return true
	}else{
		return false
	}
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */