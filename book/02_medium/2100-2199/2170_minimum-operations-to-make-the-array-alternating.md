# 2170 — Minimum Operations To Make The Array Alternating

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumOperations(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2170: Minimum Operations to Make the Array Alternating
// https://leetcode.com/problems/minimum-operations-to-make-the-array-alternating/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Frequency maps for even and odd positions
  // HashMap: O(1) lookup
	evenFreq := make(map[int]int)
  // HashMap: O(1) lookup
	oddFreq := make(map[int]int)

	for i, v := range nums {
		if i%2 == 0 {
			evenFreq[v]++
		} else {
			oddFreq[v]++
		}
	}

	// Get top 2 most frequent values for even positions
	evenTop1, evenTop2 := getTopTwo(evenFreq)
	oddTop1, oddTop2 := getTopTwo(oddFreq)

	evenCount := (n + 1) / 2 // number of even positions
	oddCount := n / 2        // number of odd positions

	if evenTop1.val != oddTop1.val {
		return (evenCount - evenTop1.count) + (oddCount - oddTop1.count)
	}

	// Try both combinations
	candidate1 := (evenCount - evenTop1.count) + (oddCount - oddTop2.count)
	candidate2 := (evenCount - evenTop2.count) + (oddCount - oddTop1.count)
	if candidate1 < candidate2 {
		return candidate1
	}
	return candidate2
}

type freqPair struct {
	val   int
	count int
}

func getTopTwo(freq map[int]int) (freqPair, freqPair) {
	top1 := freqPair{val: -1, count: 0}
	top2 := freqPair{val: -1, count: 0}

	for val, cnt := range freq {
		if cnt > top1.count {
			top2 = top1
			top1 = freqPair{val, cnt}
		} else if cnt > top2.count {
			top2 = freqPair{val, cnt}
		}
	}

	return top1, top2
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumOperations([]int{3, 1, 3, 2, 4, 3}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", minimumOperations([]int{1, 2, 2, 2, 2}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minimumOperations([]int{1, 1, 1, 1}))
	// Expected: 2
}
```
