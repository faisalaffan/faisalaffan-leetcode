# 1000 — Minimum Cost To Merge Stones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mergeStones(stones []int, k int) int
```

> **💡 Hint:** Interval DP.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1000: Minimum Cost to Merge Stones
// https://leetcode.com/problems/minimum-cost-to-merge-stones/
// Difficulty: Hard
//
// Approach: Interval DP.
//   dp[i][j] = min cost to merge stones[i:j+1] into (j-i) % (k-1) + 1 piles.
//   We can merge a subarray into 1 pile iff (len-1) % (k-1) == 0.
//   To merge dp[i][j] into 1 pile, we iterate mid where (mid-i) % (k-1) == 0,
//   and dp[i][j] = min(dp[i][mid] + dp[mid+1][j]) + sum(nums[i:j+1]).

import "fmt"

func main() {
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 2)) // 20
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 3)) // -1
	fmt.Println(mergeStones([]int{1}, 2))           // 0
}

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + stones[i]
	}

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := k; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			dp[i][j] = 1 << 60
			for m := i; m < j; m += k - 1 {
				cost := dp[i][m] + dp[m+1][j]
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
			if (j-i)%(k-1) == 0 {
				dp[i][j] += prefix[j+1] - prefix[i]
			}
		}
	}
	return dp[0][n-1]
}
```
