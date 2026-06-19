package main

// LeetCode #705: Design HashSet
// https://leetcode.com/problems/design-hashset/
// Difficulty: Easy

import "fmt"

const hashSetSize = 1000

type hashNode struct {
	key  int
	next *hashNode
}

// MyHashSet implements a hash set using separate chaining.
type MyHashSet struct {
	buckets []*hashNode
}

// Constructor creates a MyHashSet.
func Constructor() MyHashSet {
	return MyHashSet{buckets: make([]*hashNode, hashSetSize)}
}

// Add inserts a key into the set.
// Time: O(1) average. Space: O(n).
func (s *MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	idx := key % hashSetSize
	s.buckets[idx] = &hashNode{key: key, next: s.buckets[idx]}
}

// Remove deletes a key from the set.
func (s *MyHashSet) Remove(key int) {
	idx := key % hashSetSize
	curr := s.buckets[idx]
	var prev *hashNode
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				s.buckets[idx] = curr.next
			} else {
				prev.next = curr.next
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

// Contains checks if a key exists in the set.
func (s *MyHashSet) Contains(key int) bool {
	idx := key % hashSetSize
	curr := s.buckets[idx]
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.next
	}
	return false
}

func main() {
	hs := Constructor()
	hs.Add(1)
	hs.Add(2)
	fmt.Println(hs.Contains(1)) // true
	fmt.Println(hs.Contains(3)) // false
	hs.Add(2)
	fmt.Println(hs.Contains(2)) // true
	hs.Remove(2)
	fmt.Println(hs.Contains(2)) // false
}
