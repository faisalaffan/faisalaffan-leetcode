# 2935 — Maximum Strong Pair Xor Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func newBinaryTrie() *binaryTrie
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window, Trie

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Urutkan secara ascending — O(n log n)
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
