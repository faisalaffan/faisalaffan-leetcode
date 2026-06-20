# 0451 — Sort Characters By Frequency

## Deskripsi

**Soal:** [0451. Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func frequencySort(s string) string`

## Solusi Go

```go
package main

// LeetCode #451: Sort Characters By Frequency
// https://leetcode.com/problems/sort-characters-by-frequency/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 128)
	for _, ch := range s {
		freq[ch]++
	}

	// Bucket sort: index = frequency
  // Membuat slice 2D untuk DP/tabel
	buckets := make([][]byte, len(s)+1)
	for ch := 0; ch < 128; ch++ {
		if freq[ch] > 0 {
			buckets[freq[ch]] = append(buckets[freq[ch]], byte(ch))
		}
	}

	var sb strings.Builder
	for count := len(buckets) - 1; count > 0; count-- {
		for _, ch := range buckets[count] {
			sb.WriteString(strings.Repeat(string(ch), count))
		}
	}
	return sb.String()
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", frequencySort("tree"))
	// Expected: "eert" or "eetr"

	// Test case 2
	fmt.Println("Test 2:", frequencySort("cccaaa"))
	// Expected: "aaaccc" or "cccaaa"

	// Test case 3
	fmt.Println("Test 3:", frequencySort("Aabb"))
	// Expected: "bbAa" or "bbaA"
}
```
