# 2580 — Count Ways To Group Overlapping Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countWays(ranges [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2580: Count Ways to Group Overlapping Ranges
// https://leetcode.com/problems/count-ways-to-group-overlapping-ranges/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func countWays(ranges [][]int) int {
	const mod = 1_000_000_007

  // Custom sort dengan comparator
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	groups := 0
	end := -1
	for _, r := range ranges {
		if r[0] > end {
			groups++
		}
		if r[1] > end {
			end = r[1]
		}
	}

	// Each group can be in group 1 or group 2, so 2^groups ways
	ans := 1
	for i := 0; i < groups; i++ {
		ans = (ans * 2) % mod
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countWays([][]int{{6, 10}, {5, 15}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", countWays([][]int{{1, 3}, {10, 20}, {2, 5}, {4, 8}}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", countWays([][]int{{1, 2}, {3, 4}}))
	// Expected: 4
}
```
