# 1707 — Maximum Xor With An Element From Array

## Deskripsi

**Soal:** [1707. Maximum Xor With An Element From Array](https://leetcode.com/problems/maximum-xor-with-an-element-from-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Trie (pohon awalan)

**Fungsi Solusi:** `func NewBinaryTrie() *BinaryTrie`

## Solusi Go

```go
package main

// LeetCode #1707: Maximum XOR With an Element From Array
// https://leetcode.com/problems/maximum-xor-with-an-element-from-array/
// Difficulty: Hard
// Strategy: Offline Trie. Sort nums, sort queries by limit.
// Insert nums into trie as limit increases, then query max XOR.

import (
	"fmt"
	"sort"
)

// Binary Trie Node (bits 0..31 for ints up to 10^9)
type TrieNode struct {
	children [2]*TrieNode
}

type BinaryTrie struct {
	root *TrieNode
}

func NewBinaryTrie() *BinaryTrie {
	return &BinaryTrie{root: &TrieNode{}}
}

func (t *BinaryTrie) Insert(num int) {
	node := t.root
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
	}
}

func (t *BinaryTrie) QueryMaxXor(num int) int {
	// Returns max XOR value, -1 if trie is empty
	if t.root.children[0] == nil && t.root.children[1] == nil {
		return -1
	}
	node := t.root
	xor := 0
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		// Try to go opposite direction for max XOR
		desired := 1 - bit
		if node.children[desired] != nil {
			xor |= (1 << i)
			node = node.children[desired]
		} else {
			node = node.children[bit]
		}
	}
	return xor
}

func maximizeXor(nums []int, queries [][]int) []int {
	// Sort nums
	sort.Ints(nums)

	// Attach original indices to queries and sort by limit
  // Membuat slice untuk menyimpan hasil
	q := make([][3]int, len(queries)) // [x, limit, originalIdx]
	for i, query := range queries {
		q[i] = [3]int{query[0], query[1], i}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i][1] < q[j][1]
	})

	trie := NewBinaryTrie()
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	idx := 0

	for _, query := range q {
		x, limit, origIdx := query[0], query[1], query[2]
		// Insert all nums <= limit
		for idx < len(nums) && nums[idx] <= limit {
			trie.Insert(nums[idx])
			idx++
		}
		ans[origIdx] = trie.QueryMaxXor(x)
	}
	return ans
}

func main() {
	// Example 1: nums=[0,1,2,3,4], queries=[[3,1],[1,3],[5,6]] -> [3,3,7]
	nums1 := []int{0, 1, 2, 3, 4}
	queries1 := [][]int{{3, 1}, {1, 3}, {5, 6}}
	fmt.Printf("maximizeXor(%v, %v) = %v (expected [3 3 7])\n",
		nums1, queries1, maximizeXor(nums1, queries1))

	// Example 2: nums=[5,2,4,6,6,3], queries=[[12,4],[8,1],[6,3]] -> [15,-1,5]
	nums2 := []int{5, 2, 4, 6, 6, 3}
	queries2 := [][]int{{12, 4}, {8, 1}, {6, 3}}
	fmt.Printf("maximizeXor(%v, %v) = %v (expected [15 -1 5])\n",
		nums2, queries2, maximizeXor(nums2, queries2))

	// Edge case: empty result
	nums3 := []int{10, 20}
	queries3 := [][]int{{5, 5}}
	fmt.Printf("maximizeXor(%v, %v) = %v (expected [-1])\n",
		nums3, queries3, maximizeXor(nums3, queries3))
}
```
