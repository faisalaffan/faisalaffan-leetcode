# 1824 — Minimum Sideway Jumps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minSideJumps(obstacles []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1824: Minimum Sideway Jumps
// https://leetcode.com/problems/minimum-sideway-jumps/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minSideJumps(obstacles []int) int {
	// dp[i] = min jumps to reach lane i (0-indexed: 0, 1, 2 for lanes 1, 2, 3)
	dp := []int{1, 0, 1} // start at lane 2 (index 1)

	for _, obs := range obstacles {
		if obs > 0 {
			lane := obs - 1
			dp[lane] = 1 << 30 // blocked
		}
		// Try jumping sideways from other lanes
		for i := 0; i < 3; i++ {
			if i != obs-1 {
				for j := 0; j < 3; j++ {
					if j != i && j != obs-1 {
						dp[i] = min(dp[i], dp[j]+1)
					}
				}
			}
		}
	}
	return min(dp[0], min(dp[1], dp[2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSideJumps([]int{0, 1, 2, 3, 0})) // Expected: 2
	fmt.Println(minSideJumps([]int{0, 1, 1, 3, 3, 0})) // Expected: 0
	fmt.Println(minSideJumps([]int{0, 2, 1, 0, 3, 0})) // Expected: 2
}
```
