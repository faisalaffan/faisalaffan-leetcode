# 2935 — Maximum Strong Pair Xor Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** —

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func newBinaryTrie() *binaryTrie`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window, Sorting, Trie

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2935: Maximum Strong Pair XOR II
// https://leetcode.com/problems/maximum-strong-pair-xor-ii/
//
// A strong pair satisfies |x-y| <= min(x,y).
// For sorted array with x <= y, condition simplifies to y <= 2*x.
// Sort array, use sliding window with a binary trie to maintain candidates.
// For each right element, remove elements from left that violate y > 2*x,
// then query trie for max XOR with current element.

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
  // Sort O(n log n)
	sort.Ints(nums)
	trie := newBinaryTrie()
	left := 0
	ans := 0

	for _, val := range nums {
		// Maintain window where for all x, val <= 2*x (since x <= val in sorted order)
		for left < len(nums) && (val+1)/2 > nums[left] {
			trie.remove(nums[left])
			left++
		}
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
	// Example: [1,2,3,4,5] -> 7 (strong pair 3 XOR 4)
	fmt.Println(maximumStrongPairXor([]int{1, 2, 3, 4, 5}))

	// Edge cases
	fmt.Println(maximumStrongPairXor([]int{10, 100}))
	fmt.Println(maximumStrongPairXor([]int{5, 6}))
	fmt.Println(maximumStrongPairXor([]int{1, 1, 1}))
	fmt.Println(maximumStrongPairXor([]int{1, 2, 4, 8, 16}))
}
```
