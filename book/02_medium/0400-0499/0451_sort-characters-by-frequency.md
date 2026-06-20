# 0451 — Sort Characters By Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func frequencySort(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

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
  // Alokasi slice
	freq := make([]int, 128)
	for _, ch := range s {
		freq[ch]++
	}

	// Bucket sort: index = frequency
  // Matriks 2D
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
