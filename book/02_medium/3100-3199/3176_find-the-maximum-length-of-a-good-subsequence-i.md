# 3176 — Find The Maximum Length Of A Good Subsequence I

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

**Kompleksitas Waktu:** O(n^2 * k)  
**Kompleksitas Ruang:** O(n * k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3176: Find the Maximum Length of a Good Subsequence I
// https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-i/
// Difficulty: Medium
// Time: O(n^2 * k) | Space: O(n * k)

import "fmt"

func maximumLength(nums []int, k int) int {
	n := len(nums)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			for p := 0; p < i; p++ {
				if nums[i] == nums[p] {
					if dp[p][j]+1 > dp[i][j] {
						dp[i][j] = dp[p][j] + 1
					}
				} else if j > 0 {
					if dp[p][j-1]+1 > dp[i][j] {
						dp[i][j] = dp[p][j-1] + 1
					}
				}
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 3}, 2)) // Expected: 4
	fmt.Println(maximumLength([]int{1, 2, 3, 4}, 0))     // Expected: 1
}
```
