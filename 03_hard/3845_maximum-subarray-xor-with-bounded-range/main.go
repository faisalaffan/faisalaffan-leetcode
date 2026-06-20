package main

// LeetCode #3845: Maximum Subarray XOR with Bounded Range
// https://leetcode.com/problems/maximum-subarray-xor-with-bounded-range/
// Difficulty: Hard
//
// Find maximum XOR value of any subarray where subarray length does
// not exceed the given limit.
//
// Approach: Prefix XOR with sliding window + binary trie. Maintain
// a trie of prefix XORs within the window. For each i, query trie
// for max XOR with prefix[i], then insert prefix[i] into trie,
// removing prefix[i-limit] when window slides.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxXor([]int{1, 2, 3, 4}, 2))
	// Example 2
	fmt.Println(maxXor([]int{5, 1, 7, 3}, 3))
	// Edge: limit = 1
	fmt.Println(maxXor([]int{1, 1, 1}, 1))
	// Edge: all same
	fmt.Println(maxXor([]int{7, 7, 7, 7}, 4))
}

type trieNode struct {
	child [2]*trieNode
	cnt   int
}

type xorTrie struct {
	root *trieNode
}

func newXorTrie() *xorTrie {
	return &xorTrie{root: &trieNode{}}
}

func (t *xorTrie) insert(x int) {
	node := t.root
	for i := 31; i >= 0; i-- {
		b := (x >> i) & 1
		if node.child[b] == nil {
			node.child[b] = &trieNode{}
		}
		node = node.child[b]
		node.cnt++
	}
}

func (t *xorTrie) remove(x int) {
	node := t.root
	for i := 31; i >= 0; i-- {
		b := (x >> i) & 1
		node = node.child[b]
		node.cnt--
	}
}

func (t *xorTrie) maxXor(x int) int {
	if t.root.child[0] == nil && t.root.child[1] == nil {
		return 0
	}
	node := t.root
	res := 0
	for i := 31; i >= 0; i-- {
		b := (x >> i) & 1
		want := 1 - b
		if node.child[want] != nil && node.child[want].cnt > 0 {
			res |= (1 << i)
			node = node.child[want]
		} else {
			node = node.child[b]
		}
	}
	return res
}

func maxXor(nums []int, limit int) int {
	if len(nums) == 0 || limit <= 0 {
		return 0
	}

	trie := newXorTrie()
	prefix := 0
	best := 0

	for i := 0; i < len(nums); i++ {
		prefix ^= nums[i]
		if i < limit {
			trie.insert(prefix)
			if i == 0 {
				best = prefix
			} else {
				if v := trie.maxXor(prefix); v > best {
					best = v
				}
			}
			if prefix > best {
				best = prefix
			}
		} else {
			trie.remove(prefix ^ nums[i-limit])
			trie.insert(prefix)
			if v := trie.maxXor(prefix); v > best {
				best = v
			}
		}
	}

	return best
}
