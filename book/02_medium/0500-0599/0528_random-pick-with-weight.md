# 0528 — Random Pick With Weight

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(w []int) Solution
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n) for init, O(log n) per pick  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #528: Random Pick with Weight
// https://leetcode.com/problems/random-pick-with-weight/
// Difficulty: Medium
// Time: O(n) for init, O(log n) per pick
// Space: O(n)

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	sol := Constructor([]int{1, 3})
	fmt.Println(sol.PickIndex())
	fmt.Println(sol.PickIndex())
	fmt.Println(sol.PickIndex())
}

type Solution struct {
	prefixSum []int
	totalSum  int
}

func Constructor(w []int) Solution {
  // Alokasi slice integer
	prefixSum := make([]int, len(w))
	sum := 0
	for i, weight := range w {
		sum += weight
		prefixSum[i] = sum
	}
	return Solution{prefixSum: prefixSum, totalSum: sum}
}

func (s *Solution) PickIndex() int {
	target := rand.Intn(s.totalSum) + 1
	return sort.SearchInts(s.prefixSum, target)
}
```
