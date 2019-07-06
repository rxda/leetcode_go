package src

import "math"

type AllOne struct {
	Map    map[string]int
	Max    int
	Min    int
	MaxKey string
	MinKey string
}

/** Initialize your data structure here. */
func Constructora() AllOne {
	return AllOne{
		make(map[string]int),
		math.MinInt64,
		math.MaxInt64,
		"",
		"",
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if val, ok := this.Map[key]; !ok {
		this.Map[key] = 1
		if 1 < this.Min {
			this.Min = 1
			this.MinKey = key
		}
		if 1 > this.Max {
			this.Max = 1
			this.MaxKey = key
		}
	} else {
		this.Map[key]++
		if val+1 > this.Max || this.MaxKey == key{
			this.MaxKey = key
			this.Max = val + 1
		}
		if val+1 < this.Min || this.MinKey == key{
			this.Min = val + 1
			this.MinKey = key
		}
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if val, ok := this.Map[key]; ok {
		this.Map[key]--
		if val+1 > this.Max || this.MaxKey == key{
			this.MaxKey = key
			this.Max = val + 1
		}
		if val+1 < this.Min || this.MinKey == key{
			this.Min = val + 1
			this.MinKey = key
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	return this.MaxKey
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	return this.MinKey
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
