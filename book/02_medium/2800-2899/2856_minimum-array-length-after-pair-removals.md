# 2856 — Minimum Array Length After Pair Removals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumArrayLengthAfterPairRemovals(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2856: Minimum Array Length After Pair Removals
// https://leetcode.com/problems/minimum-array-length-after-pair-removals/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumArrayLengthAfterPairRemovals(nums []int) int {
	n := len(nums)
	// Find max frequency
	maxFreq := 0
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
		if freq[v] > maxFreq {
			maxFreq = freq[v]
		}
	}

	remaining := n - maxFreq
	if maxFreq > remaining {
		return maxFreq - remaining
	}
	if n%2 == 0 {
		return 0
	}
	return 1
}

func main() {
	fmt.Println(MinimumArrayLengthAfterPairRemovals([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(MinimumArrayLengthAfterPairRemovals([]int{1, 1, 2, 3}))
}
```
