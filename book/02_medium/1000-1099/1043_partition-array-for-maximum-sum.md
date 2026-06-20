# 1043 — Partition Array For Maximum Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSumAfterPartitioning(arr []int, k int) int
```

> **💡 Hint:** DP. dp[i] = max sum for prefix ending at i.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** O(n * k)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1043: Partition Array for Maximum Sum
// https://leetcode.com/problems/partition-array-for-maximum-sum/
// Difficulty: Medium
//
// Approach: DP. dp[i] = max sum for prefix ending at i.
//           For each i, try all partition lengths up to k.
// Time: O(n * k)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3)) // 84
	fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4)) // 83
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
  // Alokasi slice integer
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maxVal := 0
		for j := 1; j <= k && i-j >= 0; j++ {
			if arr[i-j] > maxVal {
				maxVal = arr[i-j]
			}
			sum := dp[i-j] + maxVal*j
			if sum > dp[i] {
				dp[i] = sum
			}
		}
	}

	return dp[n]
}
```
