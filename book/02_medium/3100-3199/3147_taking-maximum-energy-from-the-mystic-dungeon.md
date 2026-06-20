# 3147 — Taking Maximum Energy From The Mystic Dungeon

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumEnergy(energy []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3147: Taking Maximum Energy From the Mystic Dungeon
// https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/
// Difficulty: Medium
// Time: O(n) | Space: O(k)

import "fmt"

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
  // Alokasi slice integer
	dp := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = energy[i]
	}

	for i := k; i < n; i++ {
		if dp[i-k] > 0 {
			dp[i] += dp[i-k]
		}
	}

	ans := dp[n-1]
	for i := n - k - 1; i >= 0; i -= k {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumEnergy([]int{5, 2, -10, -5, 1}, 3)) // Expected: 3
	fmt.Println(maximumEnergy([]int{-2, -3, -1}, 2))        // Expected: -1
}
```
