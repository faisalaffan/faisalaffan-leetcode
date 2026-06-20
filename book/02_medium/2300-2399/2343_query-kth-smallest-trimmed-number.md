# 2343 — Query Kth Smallest Trimmed Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func smallestTrimmedNumbers(nums []string, queries [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2343: Query Kth Smallest Trimmed Number
// https://leetcode.com/problems/query-kth-smallest-trimmed-number/
// Difficulty: Medium
// Time: O(m * n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
  // Alokasi slice integer
	result := make([]int, len(queries))
	n := len(nums)

	type pair struct {
		val  string
		idx  int
	}

	for qi, q := range queries {
		k, trim := q[0], q[1]
		pairs := make([]pair, n)
		for i, s := range nums {
			pairs[i] = pair{s[len(s)-trim:], i}
		}
		sort.SliceStable(pairs, func(i, j int) bool {
			if pairs[i].val != pairs[j].val {
				return pairs[i].val < pairs[j].val
			}
			return pairs[i].idx < pairs[j].idx
		})
		result[qi] = pairs[k-1].idx
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(smallestTrimmedNumbers([]string{"102", "473", "251", "814"}, [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}}))
	// Expected: [2, 2, 1, 0]

	// Test case 2
	fmt.Println(smallestTrimmedNumbers([]string{"24", "37", "96", "04"}, [][]int{{2, 1}, {2, 2}}))
	// Expected: [3, 0]
}
```
