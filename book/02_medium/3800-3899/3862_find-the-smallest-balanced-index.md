# 3862 — Find The Smallest Balanced Index

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheSmallestBalancedIndex(nums []int) int
```

> **💡 Hint:** Compute prefix sums and suffix products, check equality at each index.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Prefix Sum

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3862: Find the Smallest Balanced Index
// https://leetcode.com/problems/find-the-smallest-balanced-index/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Compute prefix sums and suffix products, check equality at each index.

import "fmt"

func FindTheSmallestBalancedIndex(nums []int) int {
	n := len(nums)

  // Alokasi slice integer
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

  // Alokasi slice integer
	suffixProd := make([]int, n+1)
	suffixProd[n] = 1
	for i := n - 1; i >= 0; i-- {
		suffixProd[i] = suffixProd[i+1] * nums[i]
	}

	for i := 0; i < n; i++ {
		leftSum := prefixSum[i]
		rightProd := suffixProd[i+1]
		if leftSum == rightProd {
			return i
		}
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(FindTheSmallestBalancedIndex([]int{2, 1, 2})) // Expected: 1

	// Example 2
	fmt.Println(FindTheSmallestBalancedIndex([]int{2, 8, 2, 2, 5})) // Expected: 2

	// Example 3
	fmt.Println(FindTheSmallestBalancedIndex([]int{1})) // Expected: -1
}
```
