# 1040 — Moving Stones Until Consecutive Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numMovesStonesII(stones []int) []int
```

> **💡 Hint:** Sort stones. Max moves: spread left or right. Min moves: sliding window.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1040: Moving Stones Until Consecutive II
// https://leetcode.com/problems/moving-stones-until-consecutive-ii/
// Difficulty: Medium
//
// Approach: Sort stones. Max moves: spread left or right. Min moves: sliding window.
// Time: O(n log n)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStonesII([]int{7, 4, 9}))    // [1,2]
	fmt.Println(numMovesStonesII([]int{6, 5, 4, 3, 10})) // [2,3]
}

func numMovesStonesII(stones []int) []int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(stones)
	n := len(stones)

	// Max moves: fill gaps from one end, leaving one stone at the other end
	maxMoves := max(stones[n-1]-stones[1]-(n-2), stones[n-2]-stones[0]-(n-2))

	// Min moves: sliding window of size n
	minMoves := n
	j := 0
	for i := 0; i < n; i++ {
		for j+1 < n && stones[j+1]-stones[i] < n {
			j++
		}
		already := j - i + 1
		moves := n - already
		// Special case: n-1 stones already consecutive, last one far away
		if moves == 1 && stones[j]-stones[i]+1 == n-1 && (j-i+1 == n-1) {
			moves = 2
		}
		if moves < minMoves {
			minMoves = moves
		}
	}

	return []int{minMoves, maxMoves}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
