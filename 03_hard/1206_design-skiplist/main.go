package main

// LeetCode #1206: Design Skiplist
// https://leetcode.com/problems/design-skiplist/
// Difficulty: Hard

import (
	"fmt"
	"math/rand"
)

const maxLevel = 16
const prob = 0.5 // probability for level promotion

type Node struct {
	val  int
	next []*Node // next[i] = next node at level i
}

type Skiplist struct {
	head *Node
}

func Constructor() Skiplist {
	return Skiplist{
		head: &Node{val: -1, next: make([]*Node, maxLevel)},
	}
}

func randomLevel() int {
	level := 1
	for level < maxLevel && rand.Float64() < prob {
		level++
	}
	return level
}

func (sl *Skiplist) Search(target int) bool {
	cur := sl.head
	for i := maxLevel - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < target {
			cur = cur.next[i]
		}
	}
	cur = cur.next[0]
	return cur != nil && cur.val == target
}

func (sl *Skiplist) Add(num int) {
	update := make([]*Node, maxLevel)
	cur := sl.head
	for i := maxLevel - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}

	level := randomLevel()
	node := &Node{val: num, next: make([]*Node, level)}
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}
}

func (sl *Skiplist) Erase(num int) bool {
	update := make([]*Node, maxLevel)
	cur := sl.head
	for i := maxLevel - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}

	target := cur.next[0]
	if target == nil || target.val != num {
		return false
	}

	for i := 0; i < maxLevel; i++ {
		if update[i].next[i] == target {
			update[i].next[i] = target.next[i]
		}
	}
	return true
}

func main() {
	// Test case
	sl := Constructor()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	fmt.Println(sl.Search(0)) // false
	sl.Add(4)
	fmt.Println(sl.Search(1)) // true
	fmt.Println(sl.Erase(0))  // false
	fmt.Println(sl.Erase(1))  // true
	fmt.Println(sl.Search(1)) // false

	fmt.Println("---")

	// Test case 2: duplicates
	sl2 := Constructor()
	sl2.Add(1)
	sl2.Add(1)
	sl2.Add(1)
	fmt.Println(sl2.Search(1)) // true
	fmt.Println(sl2.Erase(1))  // true
	fmt.Println(sl2.Search(1)) // true (second copy still exists)
	fmt.Println(sl2.Erase(1))  // true
	fmt.Println(sl2.Erase(1))  // true
	fmt.Println(sl2.Search(1)) // false
	fmt.Println(sl2.Erase(1))  // false
}
