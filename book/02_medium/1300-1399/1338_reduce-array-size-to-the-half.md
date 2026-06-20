# 1338 — Reduce Array Size To The Half

## Deskripsi

**Soal:** [1338. Reduce Array Size To The Half](https://leetcode.com/problems/reduce-array-size-to-the-half/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) for sorting frequencies  
**Kompleksitas Ruang:** O(n) for frequency map

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1338: Reduce Array Size to The Half
// https://leetcode.com/problems/reduce-array-size-to-the-half/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7})) // 2

	// Test case 2
	fmt.Println(minSetSize([]int{7, 7, 7, 7, 7, 7})) // 1

	// Test case 3
	fmt.Println(minSetSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 5
}

// Time: O(n log n) for sorting frequencies
// Space: O(n) for frequency map
func minSetSize(arr []int) int {
	n := len(arr)
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

  // Membuat slice untuk menyimpan hasil
	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	removed := 0
	half := n / 2
	for i, c := range counts {
		removed += c
		if removed >= half {
			return i + 1
		}
	}

	return len(counts)
}
```
