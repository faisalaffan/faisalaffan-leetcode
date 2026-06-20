# 3496 — Maximize Score After Pair Deletions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximizeScoreAfterPairDeletions(nums []int, cost []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3496: Maximize Score After Pair Deletions
// https://leetcode.com/problems/maximize-score-after-pair-deletions/
// Difficulty: Medium [Paid]
// Complexity: O(n^2) time, O(n^2) space

import "fmt"

func main() {
	// Test case 1
	nums := []int{1, 2, 3, 4}
	cost := []int{1, 2, 3, 4}
	fmt.Println("Test 1:", MaximizeScoreAfterPairDeletions(nums, cost))

	// Test case 2
	nums2 := []int{5, 1, 5, 1}
	cost2 := []int{10, 1, 10, 1}
	fmt.Println("Test 2:", MaximizeScoreAfterPairDeletions(nums2, cost2))

	// Test case 3
	nums3 := []int{1, 2}
	cost3 := []int{3, 4}
	fmt.Println("Test 3:", MaximizeScoreAfterPairDeletions(nums3, cost3))
}

func MaximizeScoreAfterPairDeletions(nums []int, cost []int) int {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	// dp[l][r] = max score for subarray nums[l..r]
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case: single element
	for i := 0; i < n; i++ {
		dp[i][i] = cost[i]
	}

	// Base case: pair
	for i := 0; i+1 < n; i++ {
		dp[i][i+1] = cost[i] + cost[i+1]
	}

	for length := 3; length <= n; length++ {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1
			maxScore := 0
			for k := l; k < r; k++ {
				score := dp[l][k] + dp[k+1][r]
				if score > maxScore {
					maxScore = score
				}
			}
			dp[l][r] = maxScore
		}
	}

	return dp[0][n-1]
}
```
