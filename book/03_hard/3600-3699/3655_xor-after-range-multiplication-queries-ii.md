# 3655 — Xor After Range Multiplication Queries Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func xorAfterQueries(nums []int, queries [][]int) int
```

> **💡 Hint:** Use difference array to track multiplier per position,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3655: XOR After Range Multiplication Queries II
// https://leetcode.com/problems/xor-after-range-multiplication-queries-ii/
// Difficulty: Hard
//
// Given array nums and queries [l, r, k, v], multiply nums[i] by v for all
// positions i = l, l+k, l+2k, ... <= r. Return XOR of final array.
//
// Approach: Use difference array to track multiplier per position,
// apply queries efficiently with batch processing.

import "fmt"

func main() {
	// Example 1
	fmt.Println(xorAfterQueries([]int{1, 2, 3, 4}, [][]int{{0, 3, 1, 2}}))
	// Example 2
	fmt.Println(xorAfterQueries([]int{5, 3, 7}, [][]int{{0, 2, 1, 3}, {1, 1, 1, 2}}))
	// Edge: single element
	fmt.Println(xorAfterQueries([]int{10}, [][]int{{0, 0, 1, 5}}))
}

const MOD = 1000000007

func xorAfterQueries(nums []int, queries [][]int) int {
	n := len(nums)
	// Track multiplier per position
  // Alokasi slice integer
	mult := make([]int64, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range mult {
		mult[i] = 1
	}

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		for i := l; i <= r; i += k {
			mult[i] = (mult[i] * int64(v)) % MOD
		}
	}

	var result int64
	for i := 0; i < n; i++ {
		val := (int64(nums[i]) * mult[i]) % MOD
		result ^= val
	}
	return int(result)
}
```
