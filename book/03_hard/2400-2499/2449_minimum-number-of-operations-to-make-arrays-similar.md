# 2449 — Minimum Number Of Operations To Make Arrays Similar

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func makeSimilar(nums []int, target []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2449: Minimum Number of Operations to Make Arrays Similar
// https://leetcode.com/problems/minimum-number-of-operations-to-make-arrays-similar/
// Difficulty: Hard
//
// Sort both arrays, separate by parity (odd/even). Each operation changes
// a value by +/-2 preserving parity. Match elements of same parity in sorted
// order. Count the total positive difference (sum of target[i] - nums[i] for
// those needing increase). Each operation fixes 2 units, so ans = totalPos / 2.
// Time O(N log N) | Space O(N)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(makeSimilar([]int{8, 12, 6}, []int{2, 14, 10}))
	// Example 2
	fmt.Println(makeSimilar([]int{1, 2, 5}, []int{4, 1, 3}))
	// Already similar
	fmt.Println(makeSimilar([]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}))
}

func makeSimilar(nums []int, target []int) int64 {
	numsOdd := collectOdd(nums)
	tgtOdd := collectOdd(target)
	numsEven := collectEven(nums)
	tgtEven := collectEven(target)

  // Urutkan secara ascending — O(n log n)
	sort.Ints(numsOdd)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(tgtOdd)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(numsEven)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(tgtEven)

	var posDiff int64
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(numsOdd); i++ {
		if tgtOdd[i] > numsOdd[i] {
			posDiff += int64(tgtOdd[i]-numsOdd[i]) / 2
		}
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(numsEven); i++ {
		if tgtEven[i] > numsEven[i] {
			posDiff += int64(tgtEven[i]-numsEven[i]) / 2
		}
	}
	return posDiff
}

func collectOdd(a []int) []int {
	var res []int
	for _, v := range a {
		if v%2 != 0 {
			res = append(res, v)
		}
	}
	return res
}

func collectEven(a []int) []int {
	var res []int
	for _, v := range a {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}
```
