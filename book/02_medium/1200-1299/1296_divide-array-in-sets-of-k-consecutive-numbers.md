# 1296 — Divide Array In Sets Of K Consecutive Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isPossibleDivide(nums []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1296: Divide Array in Sets of K Consecutive Numbers
// https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
// Difficulty: Medium

// Sort array, greedily form groups of size k.
// Use frequency map to track available numbers.

// Time: O(n log n)
// Space: O(n)

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	for _, v := range nums {
		if freq[v] == 0 {
			continue
		}
		for i := 0; i < k; i++ {
			if freq[v+i] == 0 {
				return false
			}
			freq[v+i]--
		}
	}

	return true
}

func main() {
	fmt.Printf("%t (expected: true)\n", isPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
	fmt.Printf("%t (expected: false)\n", isPossibleDivide([]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3))
	fmt.Printf("%t (expected: true)\n", isPossibleDivide([]int{1, 2, 3, 4}, 2))
}
```
