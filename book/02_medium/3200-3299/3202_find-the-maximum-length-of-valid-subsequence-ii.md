# 3202 — Find The Maximum Length Of Valid Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumLength(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3202: Find the Maximum Length of Valid Subsequence II
// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/
// Difficulty: Medium
// Time: O(n * k) | Space: O(k)

import "fmt"

func maximumLength(nums []int, k int) int {
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, k)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, k)
	}

	ans := 0
	for _, v := range nums {
		cur := v % k
		for j := 0; j < k; j++ {
			need := (j - cur%k + k) % k
			dp[cur][j] = dp[need][j] + 1
			if dp[cur][j] > ans {
				ans = dp[cur][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 4, 2, 3, 1, 4}, 3)) // Expected: 4
	fmt.Println(maximumLength([]int{1, 2, 3, 4, 5}, 2))     // Expected: 3
}
```
