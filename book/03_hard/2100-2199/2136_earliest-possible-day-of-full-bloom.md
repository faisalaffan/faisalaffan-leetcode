# 2136 — Earliest Possible Day Of Full Bloom

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func earliestFullBloom(plantTime []int, growTime []int) int
```

> **💡 Hint:** Sort by grow time descending (Greedy).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2136: Earliest Possible Day of Full Bloom
// https://leetcode.com/problems/earliest-possible-day-of-full-bloom/
// Difficulty: Hard
//
// Approach: Sort by grow time descending (Greedy).
// Plant seeds sequentially, always plant the one with the longest grow time first.
// This minimizes the overall completion time since longer-growing seeds start earlier.
// For each seed, bloom day = cumulative plant time + grow time. Answer = max bloom day.

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	plantTime1 := []int{1, 4, 3}
	growTime1 := []int{2, 3, 1}
	fmt.Printf("earliestFullBloom(%v, %v) = %d (expected 9)\n",
		plantTime1, growTime1, earliestFullBloom(plantTime1, growTime1))

	// Additional tests
	plantTime2 := []int{1, 2, 3}
	growTime2 := []int{1, 2, 3}
	fmt.Printf("earliestFullBloom(%v, %v) = %d (expected 7)\n",
		plantTime2, growTime2, earliestFullBloom(plantTime2, growTime2))

	plantTime3 := []int{1}
	growTime3 := []int{1}
	fmt.Printf("earliestFullBloom(%v, %v) = %d (expected 2)\n",
		plantTime3, growTime3, earliestFullBloom(plantTime3, growTime3))

	plantTime4 := []int{3, 2, 1}
	growTime4 := []int{1, 2, 3}
	fmt.Printf("earliestFullBloom(%v, %v) = %d\n",
		plantTime4, growTime4, earliestFullBloom(plantTime4, growTime4))
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
  // Alokasi slice integer
	seeds := make([][2]int, n)
	for i := 0; i < n; i++ {
		seeds[i] = [2]int{plantTime[i], growTime[i]}
	}

	// Sort by grow time descending
  // Custom sort dengan comparator
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i][1] > seeds[j][1]
	})

	day := 0
	ans := 0
	for _, s := range seeds {
		day += s[0] // plant this seed
		bloomDay := day + s[1]
		if bloomDay > ans {
			ans = bloomDay
		}
	}
	return ans
}
```
