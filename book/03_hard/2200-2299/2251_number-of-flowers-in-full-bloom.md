# 2251 — Number Of Flowers In Full Bloom

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func fullBloomFlowers(flowers [][]int, people []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2251: Number of Flowers in Full Bloom
// https://leetcode.com/problems/number-of-flowers-in-full-bloom/
// Difficulty: Hard
//
// You are given a 0-indexed 2D integer array flowers where flowers[i] = [start_i, end_i]
// means the i-th flower will be in full bloom from start_i to end_i (inclusive).
// You are also given a 0-indexed integer array people of size n.
// For each person, return the number of flowers in full bloom at the time people[i].

import (
	"fmt"
	"sort"
)

// fullBloomFlowers returns for each person the count of flowers in bloom at their arrival time.
func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	m := len(people)

  // Alokasi slice integer
	starts := make([]int, n)
  // Alokasi slice integer
	ends := make([]int, n)
	for i, f := range flowers {
		starts[i] = f[0]
		ends[i] = f[1]
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(starts)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(ends)

  // Alokasi slice integer
	result := make([]int, m)
	for i, p := range people {
		// number of flowers that started blooming <= p
		bloomed := sort.SearchInts(starts, p+1) // first index > p

		// number of flowers that ended blooming < p
		faded := sort.SearchInts(ends, p) // first index >= p

		result[i] = bloomed - faded
	}

	return result
}

func main() {
	// Example 1
	flowers1 := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	people1 := []int{2, 3, 7, 11}
	fmt.Println(fullBloomFlowers(flowers1, people1)) // Expected: [1,2,2,2]

	// Example 2
	flowers2 := [][]int{{1, 10}, {3, 3}}
	people2 := []int{3, 3, 2}
	fmt.Println(fullBloomFlowers(flowers2, people2)) // Expected: [2,2,1]
}
```
