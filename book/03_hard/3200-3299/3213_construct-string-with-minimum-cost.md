# 3213 — Construct String With Minimum Cost

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func newTrieNode() *trieNode
```

> **💡 Hint:** Trie + DP. Build a trie from words (store min cost per node).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Trie

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3213: Construct String with Minimum Cost
// https://leetcode.com/problems/construct-string-with-minimum-cost/
// Difficulty: Hard
//
// Given a target string, an array of words, and an array of costs (same
// length), find the minimum total cost to construct the target by concatenat-
// ing words. Each word can be used any number of times. If impossible, return
// -1.
//
// Approach: Trie + DP. Build a trie from words (store min cost per node).
// DP[i] = min cost to build target[i:]. Walk trie from each position to find
// matches.

import (
	"fmt"
	"math"
)

type trieNode struct {
	child [26]*trieNode
	cost  int
}

func newTrieNode() *trieNode {
	return &trieNode{cost: math.MaxInt32}
}

func minimumCost(target string, words []string, costs []int) int {
	if len(target) == 0 {
		return 0
	}

	root := newTrieNode()
	for i, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.child[idx] == nil {
				node.child[idx] = newTrieNode()
			}
			node = node.child[idx]
		}
		if costs[i] < node.cost {
			node.cost = costs[i]
		}
	}

	n := len(target)
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		node := root
		for j := i; j < n; j++ {
			idx := target[j] - 'a'
			if node.child[idx] == nil {
				break
			}
			node = node.child[idx]
			if node.cost != math.MaxInt32 && dp[j+1] != math.MaxInt32 {
				candidate := node.cost + dp[j+1]
				if candidate < dp[i] {
					dp[i] = candidate
				}
			}
		}
	}

	if dp[0] == math.MaxInt32 {
		return -1
	}
	return dp[0]
}

func main() {
	// target="abcdef", words=["abdef","abc","d","def","ef"], costs=[100,1,1,10,5] => 7
	fmt.Println(minimumCost("abcdef", []string{"abdef", "abc", "d", "def", "ef"}, []int{100, 1, 1, 10, 5}))
	fmt.Println(minimumCost("xyz", []string{"ab", "cd"}, []int{5, 5})) // -1
	fmt.Println(minimumCost("hello", []string{"hello", "world"}, []int{3, 7})) // 3
	fmt.Println(minimumCost("aaa", []string{"a", "aa", "aaa"}, []int{5, 3, 1})) // 1
	fmt.Println(minimumCost("", []string{"a"}, []int{1})) // 0
}
```
