# 3890 — Integers With Multiple Sum Of Two Cubes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func IntegersWithMultipleSumOfTwoCubes(n int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(C^2) precompute + O(N log N) sort  |  **Ruang:** O(C^2)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3890: Integers With Multiple Sum of Two Cubes
// https://leetcode.com/problems/integers-with-multiple-sum-of-two-cubes/
// Difficulty: Medium
// Time: O(C^2) precompute + O(N log N) sort | Space: O(C^2)
// Approach: Enumerate a,b up to 1000 (since 1000^3 = 1e9). Count frequency
// of each sum. Return sums with >= 2 representations, sorted.

import (
	"fmt"
	"sort"
)

func IntegersWithMultipleSumOfTwoCubes(n int) []int {
  // HashMap: O(1) lookup
	cubeCount := make(map[int]int)
	limit := 1000
	for a := 1; a <= limit; a++ {
		a3 := a * a * a
		if a3 > n {
			break
		}
		for b := a; b <= limit; b++ {
			b3 := b * b * b
			sum := a3 + b3
			if sum > n {
				break
			}
			cubeCount[sum]++
		}
	}

	ans := []int{}
	for sum, cnt := range cubeCount {
		if cnt >= 2 {
			ans = append(ans, sum)
		}
	}
  // Sort O(n log n)
	sort.Ints(ans)
	return ans
}

func main() {
	// Example 1
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(4104)) // Expected: [1729 4104]

	// Example 2
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(578)) // Expected: []

	// Extra
	fmt.Println(IntegersWithMultipleSumOfTwoCubes(1729)) // Expected: [1729]
}
```
