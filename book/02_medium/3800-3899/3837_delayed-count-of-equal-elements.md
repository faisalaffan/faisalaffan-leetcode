# 3837 — Delayed Count Of Equal Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DelayedCountOfEqualElements(nums []int) int64
```

> **💡 Hint:** Track frequency of each value and count pairs where values are equal.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3837: Delayed Count of Equal Elements
// https://leetcode.com/problems/delayed-count-of-equal-elements/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(N)
// Approach: Track frequency of each value and count pairs where values are equal.

import "fmt"

func DelayedCountOfEqualElements(nums []int) int64 {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int64)
	var ans int64

	for _, v := range nums {
		ans += freq[v]
		freq[v]++
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(DelayedCountOfEqualElements([]int{1, 2, 1, 2, 1})) // Expected: 4

	// Example 2
	fmt.Println(DelayedCountOfEqualElements([]int{1, 1, 1, 1})) // Expected: 6

	// Example 3
	fmt.Println(DelayedCountOfEqualElements([]int{1, 2, 3})) // Expected: 0
}
```
