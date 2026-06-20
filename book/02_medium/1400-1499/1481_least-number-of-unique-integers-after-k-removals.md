# 1481 — Least Number Of Unique Integers After K Removals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func findLeastNumOfUniqueInts(arr []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n) for sorting frequencies  |  **Ruang:** O(n) for frequency map

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1481: Least Number of Unique Integers after K Removals
// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(findLeastNumOfUniqueInts([]int{5, 5, 4}, 1)) // 1

	// Test case 2
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3)) // 2

	// Test case 3
	fmt.Println(findLeastNumOfUniqueInts([]int{1, 2, 3, 4, 5}, 5)) // 0
}

// Time: O(n log n) for sorting frequencies
// Space: O(n) for frequency map
func findLeastNumOfUniqueInts(arr []int, k int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

  // Alokasi slice
	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}

  // Sort O(n log n)
	sort.Ints(counts)

	remaining := k
	uniqueCount := len(counts)

	for _, c := range counts {
		if remaining >= c {
			remaining -= c
			uniqueCount--
		} else {
			break
		}
	}

	return uniqueCount
}
```
