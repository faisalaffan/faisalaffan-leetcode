# 3785 — Minimum Swaps To Avoid Forbidden Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSwaps(nums []int, forbidden []int) int
```

> **💡 Hint:** Count forbidden values in prefix. Each swap can fix

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3785: Minimum Swaps to Avoid Forbidden Values
// https://leetcode.com/problems/minimum-swaps-to-avoid-forbidden-values/
// Difficulty: Hard
//
// Given array nums and list of forbidden values, find min swaps
// so that every prefix position has a non-forbidden value.
//
// Approach: Count forbidden values in prefix. Each swap can fix
// at most 2 positions. Answer = ceil(badPrefix / 2).

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumSwaps([]int{1, 2, 3, 4}, []int{1, 4}))
	// Example 2
	fmt.Println(minimumSwaps([]int{2, 1, 3}, []int{1}))
	// Edge: no forbidden
	fmt.Println(minimumSwaps([]int{1, 2, 3}, []int{}))
	// Edge: all forbidden
	fmt.Println(minimumSwaps([]int{1, 1, 1}, []int{1}))
}

func minimumSwaps(nums []int, forbidden []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	forbid := make(map[int]bool)
	for _, v := range forbidden {
		forbid[v] = true
	}

	// Count forbidden values not in correct position
	bad := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if forbid[nums[i]] {
			bad++
		}
	}

	return (bad + 1) / 2
}
```
