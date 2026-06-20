# 2568 — Minimum Impossible Or

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minImpossibleOR(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2568: Minimum Impossible OR
// https://leetcode.com/problems/minimum-impossible-or/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minImpossibleOR(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool)
	for _, v := range nums {
		seen[v] = true
	}

	// Check powers of 2: 1, 2, 4, 8, ...
	// If any is missing, that's the answer (since it can't be formed by OR of smaller numbers)
	pow2 := 1
	for {
		if !seen[pow2] {
			return pow2
		}
		pow2 <<= 1
	}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minImpossibleOR([]int{2, 1}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minImpossibleOR([]int{5, 3, 2}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", minImpossibleOR([]int{1, 2, 4, 8}))
	// Expected: 16
}
```
