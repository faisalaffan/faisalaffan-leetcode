# 3632 — Subarrays With Xor At Least K

## Deskripsi

**Soal:** [3632. Subarrays With Xor At Least K](https://leetcode.com/problems/subarrays-with-xor-at-least-k/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Trie (pohon awalan)

> **Ide Kunci:** Compute prefix XOR, use a binary trie to query how many prefixes

## Solusi Go

```go
package main

// LeetCode #3632: Subarrays with XOR at Least K
// https://leetcode.com/problems/subarrays-with-xor-at-least-k/
// Difficulty: Hard [Paid]
//
// Count the number of subarrays where XOR of elements is >= k.
//
// Approach: Compute prefix XOR, use a binary trie to query how many prefixes
// have XOR >= k with current prefix.

import "fmt"

func main() {
	// Example 1
	fmt.Println(subarraysWithXorAtLeastK([]int{1, 2, 3, 4}, 2))
	// Example 2
	fmt.Println(subarraysWithXorAtLeastK([]int{4, 2, 2, 6}, 6))
	// Edge: all zeros
	fmt.Println(subarraysWithXorAtLeastK([]int{0, 0, 0}, 1))
	// Edge: k = 0
	fmt.Println(subarraysWithXorAtLeastK([]int{1, 2, 3}, 0))
}

const MAX_BITS = 20

type TrieNode struct {
	children [2]*TrieNode
	count    int
}

func subarraysWithXorAtLeastK(nums []int, k int) int64 {
	root := &TrieNode{}
	var result int64
	prefix := 0

	// Insert 0 prefix
	insert(root, 0)

	for _, num := range nums {
		prefix ^= num
		// Count prefixes with XOR >= k
		result += int64(countXorGE(root, prefix, k, MAX_BITS))
		insert(root, prefix)
	}

	return result
}

func insert(root *TrieNode, val int) {
	node := root
	for i := MAX_BITS; i >= 0; i-- {
		bit := (val >> uint(i)) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
		node.count++
	}
}

func countXorGE(root *TrieNode, prefix, k int, bit int) int {
	if root == nil {
		return 0
	}
	if bit < 0 {
		return root.count
	}
	pBit := (prefix >> uint(bit)) & 1
	kBit := (k >> uint(bit)) & 1

	if kBit == 1 {
		// Need pBit ^ childBit >= 1 at this bit
		// childBit must be != pBit to make XOR bit = 1 >= kBit = 1
		return countXorGE(root.children[1-pBit], prefix, k, bit-1)
	} else {
		// kBit == 0
		// If childBit == 1-pBit, XOR bit = 1 > 0, all prefixes in this branch qualify
		cnt := 0
		if root.children[1-pBit] != nil {
			cnt += root.children[1-pBit].count
		}
		// If childBit == pBit, XOR bit = 0 == kBit, continue
		cnt += countXorGE(root.children[pBit], prefix, k, bit-1)
		return cnt
	}
}
```
