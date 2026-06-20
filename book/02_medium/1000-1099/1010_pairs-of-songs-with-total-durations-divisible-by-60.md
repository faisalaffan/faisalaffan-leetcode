# 1010 — Pairs Of Songs With Total Durations Divisible By 60

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func numPairsDivisibleBy60(time []int) int
```

> **💡 Hint:** Use modulo counting like Two Sum

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(60) = O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1010: Pairs of Songs With Total Durations Divisible by 60
// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
// Difficulty: Medium
//
// Approach: Use modulo counting like Two Sum
// Time: O(n)
// Space: O(60) = O(1)

import "fmt"

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40})) // 3
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))           // 3
}

func numPairsDivisibleBy60(time []int) int {
  // Alokasi slice integer
	count := make([]int, 60)
	result := 0

	for _, t := range time {
		mod := t % 60
		need := (60 - mod) % 60
		result += count[need]
		count[mod]++
	}

	return result
}
```
