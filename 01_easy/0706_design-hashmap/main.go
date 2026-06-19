package main

// LeetCode #706: Design HashMap
// https://leetcode.com/problems/design-hashmap/
// Difficulty: Easy

import "fmt"

const hashMapSize = 1000

type kvNode struct {
	key   int
	value int
	next  *kvNode
}

// MyHashMap implements a hash map using separate chaining.
type MyHashMap struct {
	buckets []*kvNode
}

// Constructor creates a MyHashMap.
func Constructor() MyHashMap {
	return MyHashMap{buckets: make([]*kvNode, hashMapSize)}
}

// Put inserts a key-value pair into the map.
// Time: O(1) average. Space: O(n).
func (m *MyHashMap) Put(key int, value int) {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	for curr != nil {
		if curr.key == key {
			curr.value = value
			return
		}
		curr = curr.next
	}
	m.buckets[idx] = &kvNode{key: key, value: value, next: m.buckets[idx]}
}

// Get returns the value for a key, or -1 if not found.
func (m *MyHashMap) Get(key int) int {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	for curr != nil {
		if curr.key == key {
			return curr.value
		}
		curr = curr.next
	}
	return -1
}

// Remove deletes a key from the map.
func (m *MyHashMap) Remove(key int) {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	var prev *kvNode
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				m.buckets[idx] = curr.next
			} else {
				prev.next = curr.next
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

func main() {
	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	fmt.Println(hm.Get(1)) // 1
	fmt.Println(hm.Get(3)) // -1
	hm.Put(2, 1)
	fmt.Println(hm.Get(2)) // 1
	hm.Remove(2)
	fmt.Println(hm.Get(2)) // -1
}
