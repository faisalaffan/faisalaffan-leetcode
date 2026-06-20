# 2597 — The Number Of Beautiful Subsets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func beautifulSubsets(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2597: The Number of Beautiful Subsets
// https://leetcode.com/problems/the-number-of-beautiful-subsets/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func beautifulSubsets(nums []int, k int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Group by residue modulo k
  // Membuat map (HashMap) — pencarian O(1)
	grouped := make(map[int][]int)
	for v := range freq {
		grouped[v%k] = append(grouped[v%k], v)
	}

	var ans int = 1 // empty subset
	for _, vals := range grouped {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(vals)
		// DP within each group: dp0 (ways not taking current), dp1 (ways taking current)
		dp0, dp1 := 1, 0 // dp0 = ways for "not taking previous value", dp1 = ways for "taking previous value"
		for i, v := range vals {
			ways0 := dp0 + dp1 // skip current value
			ways1 := dp0 * (1 << uint(freq[v]-1))
			if i > 0 && v-vals[i-1] == k {
				ways1 = dp0 * ((1 << uint(freq[v])) - 1)
			} else {
				ways1 = (dp0 + dp1) * ((1 << uint(freq[v])) - 1)
			}
			dp0, dp1 = ways0, ways1
		}
		ans *= (dp0 + dp1)
	}

	return ans - 1 // exclude empty subset
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", beautifulSubsets([]int{2, 4, 6}, 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", beautifulSubsets([]int{1}, 1))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", beautifulSubsets([]int{1, 2, 3, 4}, 1))
	// Expected: 7
}
```
