# 0440 — K Th Smallest In Lexicographical Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func findKthNumber(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import "fmt"

// LeetCode #440: K-th Smallest in Lexicographical Order
// https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/
// Difficulty: Hard
//
// Walk the lexicographic prefix tree. Count how many numbers exist under a
// given prefix in [1,n], then skip or descend. O(log^2 n) time, O(1) space.
// n=13, k=2 => 10

func main() {
	// Example 1
	fmt.Println("n=13, k=2 =>", findKthNumber(13, 2)) // 10
	// Example 2
	fmt.Println("n=1, k=1 =>", findKthNumber(1, 1)) // 1
	// Example 3
	fmt.Println("n=100, k=10 =>", findKthNumber(100, 10)) // 17
	// Larger
	fmt.Println("n=1000, k=100 =>", findKthNumber(1000, 100))
	// Edge: n=10, k=3 => 11? Let's see: [1,10,11,12,2,3,4,5,6,7,8,9] k=3 => 11
	fmt.Println("n=12, k=5 =>", findKthNumber(12, 5))
}

func findKthNumber(n int, k int) int {
	cur := 1
	k-- // convert to 0-indexed
	for k > 0 {
		steps := countSteps(n, cur, cur+1)
		if steps <= k {
			cur++
			k -= steps
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

// countSteps returns how many numbers in [1,n] are between `prefix` and
// `nextPrefix` when arranged in lexicographical order.
func countSteps(n int, prefix int, nextPrefix int) int {
	var steps int
	for prefix <= n {
		if nextPrefix <= n {
			steps += nextPrefix - prefix
		} else {
			steps += n - prefix + 1
		}
		prefix *= 10
		nextPrefix *= 10
	}
	return steps
}
```
