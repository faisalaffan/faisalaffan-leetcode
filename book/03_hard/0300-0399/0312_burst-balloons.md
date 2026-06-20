# 0312 — Burst Balloons

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxCoins(nums []int) int
```

> **💡 Hint:** DP Interval (Divide and Conquer).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #312: Burst Balloons
// https://leetcode.com/problems/burst-balloons/
// Difficulty: Hard
//
// Approach: DP Interval (Divide and Conquer).
//   - Add sentinel balloons with value 1 at both ends (index 0 and n+1).
//   - Define dp[i][j] = max coins from bursting all balloons in (i, j) exclusively.
//   - For each k in (i, j), consider k as the LAST balloon to burst in this interval.
//     When k bursts, its neighbors are i and j (since all balloons in between
//     have already been burst).
//   - dp[i][j] = max over k: dp[i][k] + nums[i] * nums[k] * nums[j] + dp[k][j]
//   - Answer: dp[0][n+1] where n is the original length.

import (
	"fmt"
)

func main() {
	// Example 1: [3,1,5,8] -> 167
	// Explanation: nums = [3,1,5,8] -> [3,5,8] -> [3,8] -> [8] -> []
	// coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
	nums := []int{3, 1, 5, 8}
	fmt.Println("Burst Balloons:", maxCoins(nums)) // 167

	// Example 2: [1,5] -> 10
	fmt.Println("[1,5]:", maxCoins([]int{1, 5})) // 1*0*5? Wait let me recalc.
	// Actually with sentinel: [1,1,5,1]
	// Burst 1: 1*1*5 + burst 5: 1*5*1 = 5 + 5 = 10. Yes, 10.

	// Example 3: single balloon [5] -> 5
	fmt.Println("[5]:", maxCoins([]int{5})) // 5

	// Example 4: empty
	fmt.Println("[]:", maxCoins([]int{})) // 0
}

// maxCoins returns the maximum coins obtainable by bursting all balloons.
func maxCoins(nums []int) int {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	// Add sentinel balloons with value 1.
  // Alokasi slice integer
	arr := make([]int, n+2)
	arr[0] = 1
	arr[n+1] = 1
	for i := 0; i < n; i++ {
		arr[i+1] = nums[i]
	}

	// dp[i][j] = max coins from bursting all balloons strictly between i and j.
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+2)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// Fill dp by interval length.
	for length := 2; length <= n+1; length++ {
		for i := 0; i+length <= n+1; i++ {
			j := i + length
			// Try each k as the LAST balloon to burst in (i, j).
			for k := i + 1; k < j; k++ {
				// arr[k] is the last to burst, so its neighbors are arr[i] and arr[j].
				coins := dp[i][k] + arr[i]*arr[k]*arr[j] + dp[k][j]
				if coins > dp[i][j] {
					dp[i][j] = coins
				}
			}
		}
	}

	return dp[0][n+1]
}

// max returns the larger of two ints.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Stub compatibility.
func BurstBalloons() any {
	nums := []int{3, 1, 5, 8}
	return maxCoins(nums)
}
```
