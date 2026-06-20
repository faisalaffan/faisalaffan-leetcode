# 1803 — Count Pairs With Xor In A Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPairs(nums []int, low int, high int) int
```

> **💡 Hint:** Binary Trie.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1803: Count Pairs With XOR in a Range
// https://leetcode.com/problems/count-pairs-with-xor-in-a-range/
// Difficulty: Hard
//
// Approach: Binary Trie.
//   Use a binary trie to store numbers. For each number, query the count
//   of numbers already in the trie whose XOR with it is <= K.
//   Answer = count(high) - count(low-1).
//   During query, traverse bits MSB-first:
//     - If kBit == 1: count the XOR=0 branch (all are < k at this position),
//       then continue with XOR=1 branch.
//     - If kBit == 0: must keep XOR=0, continue with XOR=0 branch.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example (low=2, high=5):", countPairs([]int{1, 4, 2, 7}, 2, 5))
	// Expected: 4

	// Example 2: all pairs
	fmt.Println("Example (low=2, high=6):", countPairs([]int{1, 4, 2, 7}, 2, 6))
	// Expected: 6 (all six pairs)

	// Edge case: identical numbers
	fmt.Println("Edge (low=0, high=0):", countPairs([]int{1, 1}, 0, 0))
	// Expected: 1 (1^1=0)

	// Single element
	fmt.Println("Edge (single):", countPairs([]int{5}, 0, 100))
	// Expected: 0
}

type TrieNode struct {
	children [2]*TrieNode
	count    int
}

func countPairs(nums []int, low int, high int) int {
	if low == 0 {
		return countLessEqual(nums, high)
	}
	return countLessEqual(nums, high) - countLessEqual(nums, low-1)
}

func countLessEqual(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	root := &TrieNode{}
	result := 0
	for _, num := range nums {
		result += query(root, num, k)
		insert(root, num)
	}
	return result
}

func insert(root *TrieNode, num int) {
	node := root
	for i := 20; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
		node.count++
	}
}

func query(root *TrieNode, num int, k int) int {
	node := root
	result := 0
	for i := 20; i >= 0; i-- {
		if node == nil {
			break
		}
		numBit := (num >> i) & 1
		kBit := (k >> i) & 1
		if kBit == 1 {
			// XOR bit = 0 makes XOR < k at this position
			if node.children[numBit] != nil {
				result += node.children[numBit].count
			}
			// Continue with XOR bit = 1 branch
			node = node.children[1-numBit]
		} else {
			// Must have XOR bit = 0 to keep XOR <= k
			node = node.children[numBit]
		}
	}
	if node != nil {
		result += node.count // XOR exactly equals k
	}
	return result
}

// Stub kept for compatibility with the repo scaffold.
func CountPairsWithXorInARange() any {
	return countPairs([]int{1, 4, 2, 7}, 2, 5)
}
```
