package main

// LeetCode #2935: Maximum Strong Pair XOR II
// https://leetcode.com/problems/maximum-strong-pair-xor-ii/
// Difficulty: Hard
//
// Approach: Binary Trie + Sliding Window.
// A strong pair satisfies |x-y| <= min(x,y).
// For sorted array, if x <= y, then condition is y <= 2x.
// Sort the array. Use a sliding window where all elements satisfy
// the strong pair condition with the current right element.
// Maintain a binary trie with counts for insertion/deletion and
// query max XOR.

import (
	"fmt"
	"sort"
)

type trieNode struct {
	children [2]*trieNode
	cnt      int
}

type binaryTrie struct {
	root *trieNode
	sz   int
}

func newBinaryTrie() *binaryTrie {
	return &binaryTrie{root: &trieNode{}}
}

func (t *binaryTrie) insert(x int) {
	node := t.root
	for i := 20; i >= 0; i-- {
		bit := (x >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &trieNode{}
		}
		node = node.children[bit]
		node.cnt++
	}
	t.sz++
}

func (t *binaryTrie) remove(x int) {
	node := t.root
	for i := 20; i >= 0; i-- {
		bit := (x >> i) & 1
		node = node.children[bit]
		node.cnt--
	}
	t.sz--
}

func (t *binaryTrie) maxXor(x int) int {
	node := t.root
	res := 0
	for i := 20; i >= 0; i-- {
		bit := (x >> i) & 1
		want := 1 - bit
		if node.children[want] != nil && node.children[want].cnt > 0 {
			res |= (1 << i)
			node = node.children[want]
		} else if node.children[bit] != nil && node.children[bit].cnt > 0 {
			node = node.children[bit]
		} else {
			break
		}
	}
	return res
}

func maximumStrongPairXor(nums []int) int {
	sort.Ints(nums)
	trie := newBinaryTrie()
	left := 0
	ans := 0

	for _, val := range nums {
		// Maintain window: for all x in window, val <= 2*x (since x <= val in sorted order)
		for left < len(nums) && nums[left] < (val+1)/2 {
			trie.remove(nums[left])
			left++
		}
		// Query max XOR with current value
		if trie.sz > 0 {
			if xr := trie.maxXor(val); xr > ans {
				ans = xr
			}
		}
		trie.insert(val)
	}
	return ans
}

func main() {
	// Example: [1,2,3,4,5] -> 7 (strong pair 3 XOR 4 = 7)
	fmt.Println(maximumStrongPairXor([]int{1, 2, 3, 4, 5}))

	// Simple cases
	fmt.Println(maximumStrongPairXor([]int{10, 100}))
	fmt.Println(maximumStrongPairXor([]int{5, 6}))
}
