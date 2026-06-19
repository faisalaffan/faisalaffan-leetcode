package main

// LeetCode #381: Insert Delete GetRandom O(1) - Duplicates allowed
// https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/
// Difficulty: Hard
//
// RandomizedCollection supports duplicates with O(1) average time.
// Uses a slice to store values and a map from value to a set of indices.
// Remove swaps the target with the last element to maintain O(1).

import (
	"fmt"
	"math/rand"
)

func main() {
	rc := Constructor()
	fmt.Println(rc.Insert(1))  // true
	fmt.Println(rc.Insert(1))  // false (duplicate exists)
	fmt.Println(rc.Insert(2))  // true
	fmt.Println(rc.GetRandom()) // 1 or 2
	fmt.Println(rc.Remove(1))  // true (removes one occurrence of 1)
	fmt.Println(rc.GetRandom()) // 1 or 2
	fmt.Println(rc.Remove(1))  // true (removes the last 1)
	fmt.Println(rc.Remove(1))  // false (no more 1s)
}

type RandomizedCollection struct {
	vals []int
	idx  map[int]map[int]struct{} // val -> set of indices
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{idx: make(map[int]map[int]struct{})}
}

func (rc *RandomizedCollection) Insert(val int) bool {
	indices := rc.idx[val]
	if indices == nil {
		indices = make(map[int]struct{})
		rc.idx[val] = indices
	}
	rc.vals = append(rc.vals, val)
	indices[len(rc.vals)-1] = struct{}{}
	return len(indices) == 1
}

func (rc *RandomizedCollection) Remove(val int) bool {
	indices := rc.idx[val]
	if len(indices) == 0 {
		return false
	}

	// Pick any index of val
	var removeIdx int
	for k := range indices {
		removeIdx = k
		break
	}

	lastIdx := len(rc.vals) - 1
	lastVal := rc.vals[lastIdx]

	if removeIdx != lastIdx {
		// Swap with last element
		rc.vals[removeIdx] = lastVal
		// Update lastVal's indices: remove lastIdx, add removeIdx
		delete(rc.idx[lastVal], lastIdx)
		rc.idx[lastVal][removeIdx] = struct{}{}
	}

	// Remove val's index and truncate
	delete(indices, removeIdx)
	rc.vals = rc.vals[:lastIdx]

	return true
}

func (rc *RandomizedCollection) GetRandom() int {
	return rc.vals[rand.Intn(len(rc.vals))]
}
