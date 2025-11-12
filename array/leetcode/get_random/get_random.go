package getrandom

/*
https://leetcode.com/problems/insert-delete-getrandom-o1/?envType=study-plan-v2&envId=top-interview-150

	Struct should containe a set that can insert, remove el and give back a random el
	- set
	- map key = value to find in set; value = index of value in set
		- use map to find if the vale in the insert, remove methods are present

*/

import (
	"math/rand"
)

type RandomizedSet struct {
	set     []int
	dataMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set:     make([]int, 0),
		dataMap: make(map[int]int),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, exists := r.dataMap[val]; exists {
		return false
	}
	r.set = append(r.set, val)
	r.dataMap[val] = len(r.set) - 1
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	index, exists := r.dataMap[val]
	if !exists {
		return false
	}
	lastIndex := len(r.set) - 1
	r.set[index], r.set[lastIndex] = r.set[lastIndex], r.set[index]
	r.dataMap[lastIndex] = index
	r.set = r.set[:lastIndex]
	delete(r.dataMap, val)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	if len(r.set) == 0 {
		return -1
	}
	randomIndex := rand.Intn(len(r.set))
	return r.set[randomIndex]
}
