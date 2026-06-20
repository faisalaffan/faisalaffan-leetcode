# 2741 — Special Permutations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SpecialPermutations(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** O(n^2 * 2^n)  
**Kompleksitas Ruang:** O(n * 2^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2741: Special Permutations
// https://leetcode.com/problems/special-permutations/
// Difficulty: Medium
// Time: O(n^2 * 2^n) | Space: O(n * 2^n)

import "fmt"

func SpecialPermutations(nums []int) int {
	n := len(nums)
	const mod = 1_000_000_007

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, 1<<n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}

	for mask := 1; mask < 1<<n; mask++ {
		for last := 0; last < n; last++ {
			if dp[mask][last] == 0 {
				continue
			}
			for next := 0; next < n; next++ {
				if mask&(1<<next) != 0 {
					continue
				}
				if nums[last]%nums[next] == 0 || nums[next]%nums[last] == 0 {
					newMask := mask | (1 << next)
					dp[newMask][next] = (dp[newMask][next] + dp[mask][last]) % mod
				}
			}
		}
	}

	var result int
	fullMask := (1 << n) - 1
	for i := 0; i < n; i++ {
		result = (result + dp[fullMask][i]) % mod
	}
	return result
}

func main() {
	fmt.Println(SpecialPermutations([]int{1, 2, 3}))
	fmt.Println(SpecialPermutations([]int{2, 3, 6}))
}
```
